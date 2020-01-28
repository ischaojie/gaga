# gaga
gaga is just another web framework base on go language, it's just for fun, for learn.

gaga是基于go语言的web框架，实现上参考了gin框架，请不要将其用于生产上。
它只是用来学习web框架的原理而造的另一个轮子。

> [web框架实现原理分析]()

## Feature
- 利用前缀树实现路由匹配
- 支持动态路由
- 支持路由分组
- 支持JSON，Html格式响应
- 支持中间件
- 错误恢复

## How to use

```text
go get -U github.com/shiniao/gaga
```

gaga的使用方法和gin很相似，**example** 文件夹有详细的例子。

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
        // c.JSON()
        // c.String()
    })
    
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