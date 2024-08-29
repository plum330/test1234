package base

import (
	"os"
	"path/filepath"
	"testing"
)

func TestModuleVersion(t *testing.T) {
	v, err := ModuleVersion("golang.org/x/mod")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}

func TestModulePath(t *testing.T) {
	if err := os.Mkdir("/tmp/test_mod", os.ModePerm); err != nil {
		t.Fatal(err)
	}

	f, err := os.Create("/tmp/test_mod/go.mod")
	if err != nil {
		t.Fatal(err)
	}

	mod := `module github.com/go-kratos/kratos/v2

go 1.19`
	_, err = f.WriteString(mod)
	if err != nil {
		t.Fatal(err)
	}

	p, err := ModulePath("/tmp/test_mod/go.mod")
	if err != nil {
		t.Fatal(err)
	}
	if p != "github.com/go-kratos/kratos/v2" {
		t.Fatalf("want: %s, got: %s", "github.com/go-kratos/kratos/v2", p)
	}

	t.Cleanup(func() { os.RemoveAll("/tmp/test_mod") })
}

func TestModuleName(t *testing.T) {
	goModPath := filepath.Join("/tmp/test_mod", "go.mod")
	moduleName := "helloworld/example"

	if err := os.Mkdir("/tmp/test_mod", os.ModePerm); err != nil {
		t.Fatal(err)
	}

	f, err := os.Create(goModPath)
	if err != nil {
		t.Fatal(err)
	}

	mod := `module github.com/go-kratos/kratos/v2

go 1.19`
	_, err = f.WriteString(mod)
	if err != nil {
		t.Fatal(err)
	}

	err = ModuleName(goModPath, moduleName, "github.com/go-kratos/kratos/v2")
	if err != nil {
		t.Fatal(err)
	}

	p, err := ModulePath(goModPath)
	if err != nil {
		t.Fatal(err)
	}
	if p != moduleName {
		t.Fatalf("want: %s, got: %s", "github.com/go-kratos/kratos/v2", p)
	}

	t.Cleanup(func() { os.RemoveAll("/tmp/test_mod") })
}
