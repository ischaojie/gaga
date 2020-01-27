/*
go官方http包使用方法如下：
http.Handle("/foo", fooHandler)

http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

log.Fatal(http.ListenAndServe(":8080", nil))

ListenAndServe第二个参数是Handler接口，它的定义如下：
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

如果该参数为nil，则使用官方默认的实现。
ServerHTTP接管了所有的request请求，然后将不同的路由转发给不同的HandlerFunc。

默认的实现不支持http方法定义&动态路由，所以要自己实现ServerHTTP接管所有请求。

gin框架的实现：

g := gin.New()
g.get("/home", Home)
g.run(":8000")

func Home(w http.ResponseWriter, r *http.Request) {}


*/

package gaga

import (
	"net/http"
)

// HandlerFunc 是路由处理函数
type HandlerFunc func(c *Context)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

// Engine
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

// New 新建一个引擎
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (g *RouterGroup) Group(prefix string) (rg *RouterGroup) {
	engine := g.engine
	rg = &RouterGroup{
		prefix: g.prefix + prefix,
		parent: g,
		engine: engine,
	}
	return
}

// addRoute 绑定路由到handler
func (g *RouterGroup) addRoute(method string, path string, handler HandlerFunc) {
	g.engine.router.addRoute(method, g.prefix+path, handler)
}
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.router.Handle(c)
}

// Run 运行http服务
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (g *RouterGroup) Get(path string, handler HandlerFunc) {
	g.engine.router.addRoute("GET", path, handler)
}

func (g *RouterGroup) Post(path string, handler HandlerFunc) {
	g.engine.router.addRoute("POST", path, handler)
}
