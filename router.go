package gaga

import (
	"fmt"
	"net/http"
	"strings"
)

type node struct {
	path     string           // 路由路径
	part     string           // 路由中由'/'分隔的部分， 比如路由/hello/:name，那么part就是hello和:name
	children map[string]*node // 子节点
	isWild   bool             // 是否精确匹配，true代表当前节点是通配符，模糊匹配
	isPath   bool             // 是否是路由结尾
}

type router struct {
	root  map[string]*node       // 路由树根节点，每个http方法都是一个路由树
	route map[string]HandlerFunc // 路由处理handler
}

func newRouter() *router {
	return &router{root: make(map[string]*node), route: make(map[string]HandlerFunc)}
}

func (n *node) String() string {
	return fmt.Sprintf("node{path=%s, part=%s, isWild=%t, isPath=%t}", n.path, n.part, n.isWild, n.isPath)
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
// g.Get() 会调用addRoute方法将path添加到路由树上面
func (r *router) addRoute(method, path string, handler HandlerFunc) {
	parts := parsePath(path)
	if _, ok := r.root[method]; !ok {
		r.root[method] = &node{children: make(map[string]*node)}
	}
	root := r.root[method]
	key := method + "-" + path
	// 将parts插入到路由树
	for _, part := range parts {
		if root.children[part] == nil {
			root.children[part] = &node{
				part:     part,
				children: make(map[string]*node),
				isWild:   part[0] == ':' || part[0] == '*'}
		}
		root = root.children[part]
	}
	root.isPath = true
	root.path = path
	// 绑定路由和handler
	r.route[key] = handler
}

// getRoute 获取路由树节点以及路由变量
// method用来判断属于哪一个方法路由树，path用来获取路由树节点和参数
func (r *router) getRoute(method, path string) (node *node, params map[string]string) {
	params = map[string]string{}
	searchParts := parsePath(path)

	// get method trie
	var ok bool
	if node, ok = r.root[method]; !ok {
		return nil, nil
	}

	// 在路由树上查找该路径
	for i, part := range searchParts {
		var temp string
		// 查找child是否等于part
		for _, child := range node.children {
			if child.part == part || child.isWild {
				// 添加参数
				if child.part[0] == '*' {
					params[child.part[1:]] = strings.Join(searchParts[i:], "/")
				}
				if child.part[0] == ':' {
					params[child.part[1:]] = part
				}
				temp = child.part
			}

		}
		// 遇到通配符*，直接返回
		if temp[0] == '*' {
			return node.children[temp], params
		}
		node = node.children[temp]

	}

	return

}

// handle 用来绑定路由和handlerFunc
func (r *router) Handle(c *Context) {

	// 获取路由树节点和动态路由中的参数
	node, params := r.getRoute(c.Method, c.Path)
	if node != nil {
		c.Params = params
		key := c.Method + "-" + node.path
		// 将路由对应的handler添加到c.handler
		// c.handler会根据顺序调用中间件handler和路由handler
		c.handlers = append(c.handlers, r.route[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			// 	error
			c.String(http.StatusNotFound, "404 NOT FOUND %s \n", c.Path)
		})

	}
	c.Next()
}
