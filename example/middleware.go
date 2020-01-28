package main

import (
	"github.com/shiniao/gaga"
	"log"
	"time"
)

// 日志中间件
// 记录时间
func Logger() gaga.HandlerFunc {
	return func(c *gaga.Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.R.RequestURI, time.Since(t))
	}
}
