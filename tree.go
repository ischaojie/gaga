package gaga
/*
一个路由前缀树看起来像这样：
				      /
			/hello	/static       /home
         /:name    /*filepath
       /profile




import (
	"fmt"
	"strings"
)
/*


func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.path, n.part, n.isWild)
}

// matchChild 返回当前part对应的树节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren 返回当前part之前的所有节点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert 用来插入路由节点，将parts挂靠到路由树上面
// 具体过程可查找前缀树的插入算法
func (n *node) insert(path string, parts []string, tail int) {

	// 如果到达路由树末尾返回
	// 比如路由 /home/:name， len(parts)==2, tail==2
	// 说明这条路由已经插入完成
	if len(parts) == tail {
		n.path = path
		return
	}

	// 依次插入part到node上，利用递归

	part := parts[tail]
	child := n.matchChild(part)
	// 如果当前节点没有孩子，插入
	if child == nil {
		// 如果part第一个字符是 : 或者 * ，isWild成立
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	// 递归往下
	child.insert(path, parts, tail+1)
}


// search 用来查询路由树上是否存在该路由
func (n *node) search(parts []string, tail int) *node {

	// 当找到路由末尾或者遇到通配符*，返回当前node
	if len(parts) == tail || strings.HasPrefix(n.part, "*") {
		if n.path == "" {
			return nil
		}
		return n
	}

	part := parts[tail]
	children := n.matchChildren(part)
	// 递归查找路由树上是否存在该路径
	for _, child := range children {
		res := child.search(parts, tail+1)
		if res != nil {
			return res
		}
	}

	return nil
}
*/