package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	engine := gin.Default()
	engine.GET("/hello", func(c *gin.Context) {
		t := time.Now()
		time.Sleep(2 * time.Second)
		c.Writer.WriteString("hello")
		t2 := time.Since(t)
		fmt.Println("duration: ", t2)
	})
	engine.Run(":8999")
}
