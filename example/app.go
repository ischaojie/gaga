package main

import (
	"fmt"
	"github.com/shiniao/gaga"
	"net/http"
)

func main() {
	g := gaga.New()
	g.Get("/", Home)
	g.Get("/hello/:name", Hello)

	_ = g.Run(":8000")
}

func Home(c *gaga.Context) {
	c.Html(http.StatusOK, "<h1>I'm gaga.</h1>")
}
func Hello(c *gaga.Context) {
	html := fmt.Sprintf("<h1>hello, %s </h1>", c.Params["name"])
	c.Html(http.StatusOK, html)
}
