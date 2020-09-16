package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client,err:=rpc.DialHTTP("tcp","localhost:8081")
	if err!=nil{
		log.Fatal(err)
	}
	req :="llll"
	var resp *string

	//同步调用方法
	err=client.Call("Person.get",req,&resp)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(*resp)

	//异步调用方法
	syncCAll:=client.Go("Person.get",req,&resp,nil)
		replyDone:=<-syncCAll.Done
		fmt.Println(replyDone)
		fmt.Println(*resp)

}
