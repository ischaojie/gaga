package gaga

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	W http.ResponseWriter
	R *http.Request

	Method     string
	Path       string
	Params     map[string]string
	StatusCode int

	handlers []HandlerFunc
	index    int // 记录中间件位置
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W:      w,
		R:      r,
		Method: r.Method,
		Path:   r.URL.Path,
		index:  -1,
	}
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

// SetHeader 设置http头信息
func (c *Context) SetHeader(key, value string) {
	c.W.Header().Set(key, value)
}

// Status 设置http状态
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.W.WriteHeader(code)
}

// Query 获取http请求中的query
func (c *Context) Query(key string) (value string) {
	value = c.R.URL.Query().Get(key)
	return
}

// PostFrom 获取表单信息
func (c *Context) PostForm(key string) (value string) {
	value = c.R.FormValue(key)
	return
}

func (c *Context) Param(key string) (value string) {
	value, _ = c.Params[key]
	return
}

// JSON 以json格式返回消息
func (c *Context) JSON(statusCode int, obj interface{}) {
	c.StatusCode = statusCode
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.W)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.W, err.Error(), 500)
	}
}

func (c *Context) String(statusCode int, format string, value ...interface{}) {
	c.StatusCode = statusCode
	c.SetHeader("Content-Type", "text/plain")
	_, _ = fmt.Fprintf(c.W, format, value...)
}

func (c *Context) Html(statusCode int, html string) {
	c.StatusCode = statusCode
	c.SetHeader("Content-Type", "text/html")
	_, _ = c.W.Write([]byte(html))
}

func (c *Context) Fail(statusCode int, err string) {
	c.index = len(c.handlers)
	c.JSON(statusCode, H{"message": err})
}
