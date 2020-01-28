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
	"log"
	"net/http"
	"strings"
)

// HandlerFunc 是路由处理函数
type HandlerFunc func(c *Context)

// 路由分组
type RouterGroup struct {
	prefix     string        // 路由组前缀
	middleware []HandlerFunc // 中间件
	parent     *RouterGroup  // 父
	engine     *Engine       // engine
}

// Engine
// Engine本身也是一个路由组，相当于全局路由组
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup // all groups
}

// New 新建一个引擎
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	log.Printf("hello, I'm gaga.")
	return engine
}

// Default 包含默认的中间件实现
func Default() *Engine {
	g := New()
	// 添加日志和错误恢复中间件
	g.Use(Logger(), Recovery())
	return g
}

// Group 定义一个新的路由组
// 所有的group分享同一个engine实例
func (g *RouterGroup) Group(prefix string) (rg *RouterGroup) {
	engine := g.engine
	rg = &RouterGroup{
		prefix: g.prefix + prefix,
		parent: g,
		engine: engine,
	}
	return
}

// Use 添加中间件到路由组
func (g *RouterGroup) Use(middleware ...HandlerFunc) {
	g.middleware = append(g.middleware, middleware...)
}

// addRoute 绑定路由到handler
func (g *RouterGroup) addRoute(method string, path string, handler HandlerFunc) {
	g.engine.router.addRoute(method, g.prefix+path, handler)
}

// Get 实现了GET方法路由定义
func (g *RouterGroup) Get(path string, handler HandlerFunc) {
	g.engine.router.addRoute("GET", path, handler)
}

// Post 实现了POST方法路由定义
func (g *RouterGroup) Post(path string, handler HandlerFunc) {
	g.engine.router.addRoute("POST", path, handler)
}

// ServerHTTP 实现了官方Handler接口，用来接管所有request
// 所以我们可以定义自己的Get/Post方法，添加中间件等
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc
	// 遍历路由组
	for _, g := range e.groups {
		// 如果前缀相同，说明中间件作用域该路由组
		if strings.HasPrefix(r.URL.Path, g.prefix) {
			middlewares = append(middlewares, g.middleware...)
		}
	}
	c := newContext(w, r)
	c.handlers = middlewares
	e.router.Handle(c)
}

// Run 封装了http包的ListenAndServe
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
