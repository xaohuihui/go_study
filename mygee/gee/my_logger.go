package gee

import (
	"log"
	"time"
)

// author: songyanhui
// datetime: 2022/1/4 15:05:13
// software: GoLand

func Logger() HandlerFunc {
	return func(c *Context) {
		// start timer
		t := time.Now()
		// process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
