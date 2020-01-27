package gaga

import (
	"net/http"
	"strings"
)

type router struct {
	root  map[string]*node       // 路由树根节点，每个http方法都是一个路由树
	route map[string]HandlerFunc // 路由处理handler
}

func newRouter() *router {
	return &router{root: make(map[string]*node), route: make(map[string]HandlerFunc)}
}

// parsePath 分隔路由为part字典
// 比如路由/home/:name将被分隔为["home", ":name"]
func parsePath(path string) (parts []string) {
	// 将path以"/"分隔为parts
	par := strings.Split(path, "/")
	for _, p := range par {
		if p != "" {
			parts = append(parts, p)
			// 如果part是以通配符*开头的
			if p[0] == '*' {
				break
			}
		}
	}
	return
}

// addRoute 绑定路由到handler
func (r *router) addRoute(method string, path string, handler HandlerFunc) {
	parts := parsePath(path)

	// 查找是否存在当前http方法的路由树
	// 不存在的话新建一个
	if _, ok := r.root[method]; !ok {
		r.root[method] = &node{}

	}
	key := method + "-" + path
	// 将parts插入到路由树
	r.root[method].insert(path, parts, 0)
	// 绑定路由和handler
	r.route[key] = handler
}

// getRoute 获取路由树节点以及路由变量
func (r *router) getRoute(method, path string) (node *node, params map[string]string) {
	params = map[string]string{}
	searchParts := parsePath(path)

	root, ok := r.root[method]
	// 找不到该method对应的路由树
	if !ok {
		return nil, nil
	}

	// 在路由树上查找该路径
	node = root.search(searchParts, 0)
	if node != nil {
		// 处理part中的通配符
		parts := parsePath(node.path)
		// 添加动态路由中的参数
		for i, part := range parts {
			// 如果part是路由参数
			if part[0] == ':' {
				params[part[1:]] = searchParts[i]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[i:], "/")
				break
			}
		}
		return
	}
	return nil, nil
}

// handle 用来绑定路由和handlerFunc
func (r *router) Handle(c *Context) {

	// 获取路由树节点和动态路由中的参数
	node, params := r.getRoute(c.Method, c.Path)
	if node != nil {
		c.Params = params
		key := c.Method + "-" + node.path
		// 将路由绑定到对应的处理函数
		r.route[key](c)
	} else {
		// 	error
		c.String(http.StatusNotFound, "404 NOT FOUND %s \n", c.Path)
	}
}
