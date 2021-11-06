package consul

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/registry"
	"github.com/hashicorp/consul/api"
)

// Client is consul client config
type Client struct {
	cli    *api.Client
	ctx    context.Context
	cancel context.CancelFunc
}

// NewClient creates consul client
func NewClient(cli *api.Client) *Client {
	c := &Client{cli: cli}
	c.ctx, c.cancel = context.WithCancel(context.Background())
	return c
}

// Service get services from consul
func (d *Client) Service(ctx context.Context, service string, index uint64, passingOnly bool) ([]*registry.ServiceInstance, uint64, error) {
	opts := &api.QueryOptions{
		WaitIndex: index,
		WaitTime:  time.Second * 55,
	}
	opts = opts.WithContext(ctx)
	entries, meta, err := d.cli.Health().Service(service, "", passingOnly, opts)
	if err != nil {
		return nil, 0, err
	}

	services := make([]*registry.ServiceInstance, 0)

	for _, entry := range entries {
		var version string
		for _, tag := range entry.Service.Tags {
			strs := strings.SplitN(tag, "=", 2)
			if len(strs) == 2 && strs[0] == "version" {
				version = strs[1]
			}
		}
		var endpoints []string
		for scheme, addr := range entry.Service.TaggedAddresses {
			if scheme == "lan_ipv4" || scheme == "wan_ipv4" || scheme == "lan_ipv6" || scheme == "wan_ipv6" {
				continue
			}
			endpoints = append(endpoints, addr.Address)
		}
		services = append(services, &registry.ServiceInstance{
			ID:        entry.Service.ID,
			Name:      entry.Service.Service,
			Metadata:  entry.Service.Meta,
			Version:   version,
			Endpoints: endpoints,
		})
	}
	return services, meta.LastIndex, nil
}

// Register register service instacen to consul
func (d *Client) Register(ctx context.Context, svc *registry.ServiceInstance, enableHealthCheck bool) error {
	var addr string
	var port uint64

	for _, endpoint := range svc.Endpoints {
		raw, err := url.Parse(endpoint)
		if err != nil {
			return err
		}
		serviceID := fmt.Sprintf("%v-%v", svc.ID, raw.Scheme)
		serviceName := fmt.Sprintf("%v-%v", svc.Name, raw.Scheme)
		addr = raw.Hostname()
		port, _ = strconv.ParseUint(raw.Port(), 10, 16)

		asr := &api.AgentServiceRegistration{
			ID:   serviceID,
			Name: serviceName,
			Meta: svc.Metadata,
			Tags: []string{
				fmt.Sprintf("version=%s", svc.Version),
				fmt.Sprintf("schema=%s", raw.Scheme),
			},
			Address: addr,
			Port:    int(port),
			Checks: []*api.AgentServiceCheck{
				{
					CheckID:                        "service:" + svc.ID,
					TTL:                            "30s",
					Status:                         "passing",
					DeregisterCriticalServiceAfter: "90s",
				},
			},
		}
		if enableHealthCheck {
			asr.Checks = append(asr.Checks, &api.AgentServiceCheck{
				TCP:                            fmt.Sprintf("%s:%d", addr, port),
				Interval:                       "20s",
				Status:                         "passing",
				DeregisterCriticalServiceAfter: "90s",
			})
		}

		err = d.cli.Agent().ServiceRegister(asr)
		if err != nil {
			return err
		}
	}

	go func() {
		ticker := time.NewTicker(time.Second * 20)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				_ = d.cli.Agent().UpdateTTL("service:"+svc.ID, "pass", "pass")
			case <-d.ctx.Done():
				return
			}
		}
	}()
	return nil
}

// Deregister deregister service by service ID
func (d *Client) Deregister(ctx context.Context, svc *registry.ServiceInstance) error {
	d.cancel()

	for _, endpoint := range svc.Endpoints {
		raw, err := url.Parse(endpoint)
		if err != nil {
			return err
		}
		id := fmt.Sprintf("%v-%v", svc.ID, raw.Scheme)
		err = d.cli.Agent().ServiceDeregister(id)
		if err != nil {
			return err
		}
	}

	return nil
}
