package gaga

func main() {
	g := New()
	g.Get("/", Home)
	g.Run(":8080")
}

func Home(c *Context) {
	c.JSON(200, H{
		"name": "gaga",
	})
}
