package discovery

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/internal/endpoint"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/resolver"
)

type discoveryResolver struct {
	w   registry.Watcher
	cc  resolver.ClientConn
	log *log.Helper

	ctx    context.Context
	cancel context.CancelFunc

	insecure bool
}

func (r *discoveryResolver) watch() {
	for {
		select {
		case <-r.ctx.Done():
			return
		default:
		}
		ins, err := r.w.Next()
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return
			}
			r.log.Errorf("[resolver] Failed to watch discovery endpoint: %v", err)
			time.Sleep(time.Second)
			continue
		}
		r.update(ins)
	}
}

func (r *discoveryResolver) update(ins []*registry.ServiceInstance) {
	var addrs []resolver.Address
	for _, in := range ins {
		endpoint, err := endpoint.ParseEndpoint(in.Endpoints, "grpc", !r.insecure)
		if err != nil {
			r.log.Errorf("[resolver] Failed to parse discovery endpoint: %v", err)
			continue
		}
		if endpoint == "" {
			continue
		}
		addr := resolver.Address{
			ServerName: in.Name,
			Attributes: parseAttributes(in.Metadata),
			Addr:       endpoint,
		}
		addrs = append(addrs, addr)
	}
	if len(addrs) == 0 {
		r.log.Warnf("[resolver] Zero endpoint found,refused to write, instances: %v", ins)
		return
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
	b, _ := json.Marshal(ins)
	r.log.Infof("[resolver] update instances: %s", b)
}

func (r *discoveryResolver) Close() {
	r.cancel()
	r.w.Stop()
}

func (r *discoveryResolver) ResolveNow(options resolver.ResolveNowOptions) {}

func parseAttributes(md map[string]string) *attributes.Attributes {
	pairs := make([]interface{}, 0, len(md))
	for k, v := range md {
		pairs = append(pairs, k, v)
	}
	return attributes.New(pairs...)
}
