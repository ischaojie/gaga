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
	r.addRoute("GET", "/hello/:name/profile", nil)
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
	n, params := r.getRoute("GET", "/static/css/base.css")
	n2, params2 := r.getRoute("GET", "/hello/gaga")
	n3, params3 := r.getRoute("GET", "/hello/gaga/profile")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}
	if n.path != "/static/*filepath" {
		t.Fatal("expected: /static/*filepath, but: ", n.path)
	}

	if n2.path != "/hello/:name" {
		t.Fatal("expected: /hello/:name, but: ", n2.path)
	}

	if n3.path != "/hello/:name/profile" {
		t.Fatal("expected: /hello/:name/profile, but: ", n3.path)
	}

	if params["filepath"] != "css/base.css" {
		t.Fatal("expected: css/base.css, but: ", params["filepath"])
	}

	if params2["name"] != "gaga" {
		t.Fatal("expected gaga, but: ", params2["name"])
	}
	if params3["name"] != "gaga" {
		t.Fatal("expected gaga, but: ", params2["name"])
	}

	fmt.Printf("matched path: %s, params['filepath']: %s\n", n.path, params["filepath"])
	fmt.Printf("matched path: %s, params['name']: %s\n", n2.path, params2["name"])
	fmt.Printf("matched path: %s, params['name']: %s\n", n3.path, params3["name"])

}

func TestGroup(t *testing.T) {
	g := New()
	v1 := g.Group("/v1")
	v2 := v1.Group("/v2")
	if v2.prefix != "/v1/v2" {
		t.Fatal("v2's prefix should be /v1/v2 but get: ", v2.prefix)
	}
}
