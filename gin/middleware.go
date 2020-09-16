package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r:=gin.Default()

	//全局使用的中间件
	r.Use(gin.Logger())

	//传入两个中间件
	r.GET("/",func1,func2)
	r.Run(":8090")
}
func func1(c *gin.Context)  {
	c.Set("name",12345)
	//调用下一个中间件 等带下一个中间件执行完成之后执行下面代码
	c.Next()

	// 阻止下一个中间件执行
	//c.Abort()

	fmt.Println("func1中间件执行完了")
}
func func2(c *gin.Context)  {
	value,ok:=c.Get("name")
	if ok{
		fmt.Println("func2中间件 ",value)
	}
}