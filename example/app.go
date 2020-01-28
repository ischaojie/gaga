package main

import (
	"fmt"
	"github.com/shiniao/gaga"
	"net/http"
)

func main() {
	g := gaga.New()

	// use middleware
	g.Use(gaga.Logger())
	g.Use(gaga.Recovery())

	// router group
	v1 := g.Group("/v1")
	{
		v1.Get("/", Home)
		v1.Get("/hello/:name", Hello)
	}

	// 错误恢复
	// 访问会报错，但不会导致服务崩溃，后台会打印错误信息
	// 前端会显示500错误
	g.Get("/panic", func(c *gaga.Context) {
		c.String(http.StatusOK, []string{"shiniao"}[88])
	})

	_ = g.Run("8000")
}

func Home(c *gaga.Context) {
	c.Html(http.StatusOK, "<h1>I'm gaga.</h1>")
}
func Hello(c *gaga.Context) {
	html := fmt.Sprintf("<h1>hello, %s </h1>", c.Params["name"])
	c.Html(http.StatusOK, html)
}
