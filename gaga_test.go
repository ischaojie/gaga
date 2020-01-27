package gaga

import (
	"reflect"
	"testing"
)

func newTestRouter() (r *router) {
	r = newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/a/b", nil)
	r.addRoute("GET", "/static/*filepath", nil)
	return
}

func TestParsePath(t *testing.T) {
	ok := reflect.DeepEqual(parsePath("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePath("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePath failed.")
	}
}
