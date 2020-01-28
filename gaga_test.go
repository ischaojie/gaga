package gaga

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/a/b", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/static/*filepath", nil)
	return r
}

func TestParsePath(t *testing.T) {
	ok := reflect.DeepEqual(parsePath("/hello/:name"), []string{"hello", ":name"})
	ok = ok && reflect.DeepEqual(parsePath("/hello/*name/*"), []string{"hello", "*name"})
	if !ok {
		t.Fatal("test parsePath failed.")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, params := r.getRoute("GET", "/hello/shiniao")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}
	fmt.Println(n.part)

	if n.path != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}
	if params["name"] != "shiniao" {
		t.Fatal("name should equal shiniao")
	}
	fmt.Printf("matched path: %s, params['name']: %s\n", n.path, params["name"])
}

func TestGroup(t *testing.T) {
	g := New()
	v1 := g.Group("/v1")
	v2 := v1.Group("/v2")
	if v2.prefix != "/v1/v2" {
		t.Fatal("v2's prefix should be /v1/v2 but get: ", v2.prefix)
	}
}
