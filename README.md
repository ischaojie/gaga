# gaga
gaga is just another web framework base on go language, the implementation refers to the gin framework.

it's just for fun, for learn.

Please don't use it in production, It's just another wheel built to learn the principles of the web framework.


> [前缀树算法实现路由匹配原理解析](https://blog.shiniao.fun/2020/01/28/前缀树算法实现路由匹配原理解析/)

## 特点
- The prefix tree algorithm is used to realize route matching
- Support dynamic routing
- Support routing group
- Support JSON & Html & string formatted response
- Support middleware
- Support error recovery

## 如何使用

```text
go get -U github.com/shiniao/gaga
```

The **example** package have more examples。

a simple example：
```go
package main
import (
    "github.com/shiniao/gaga"
    "net/http"
)

func main() {
	g := gaga.New()
    // 或者使用gaga.Default()
    // Default默认添加了内置的logger和recover中间件 
    
g.Get("/", func(c *gaga.Context) {
        c.Html(http.StatusOK, "<h1>hello, gaga !</h2>")
        // 还可以使用 c.JSON() 和 c.String()
    })
    
    g.Run(":6000")
}
 
```

Of course，gaga also support routing's group：
```go
package main
import (
    "github.com/shiniao/gaga"
    "net/http"
)

func main() {
	g := gaga.Default()
    v1 := g.Group("/v1")    
    {
        v1.Get("/", func(c *gaga.Context) {
            c.Html(http.StatusOK, "<h1>hello, gaga !</h2>")
        })
        v1.Get("/profile", func(c *gaga.Context){})
    }

    g.Run(":6000")
}

```

## Wheels

> The best way to learn is build wheels ！

- [nan](https://github.com/shiniao/nan): 一个语言解释器实现
- [mid](https://github.com/shiniao/mid)：markdown编译器实现
- [gaga](https://github.com/shiniao/gaga): web框架实现

## Contact

有任何问题欢迎到微博找我[@潮戒](https://weibo.com/zhuzhezhe)