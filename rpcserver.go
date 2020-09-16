package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)
type Person struct {
	name string
	age int
}

func (p *Person)Get(req string,resp *string)error  {
	*resp = req+"abc"
	return nil
}
func main() {
	person:=new(Person)
	if err:=rpc.Register(person);err!=nil{
		log.Fatal(err)
	}

	rpc.HandleHTTP()

	listen,err:=net.Listen("tcp",":8081")
	if err!=nil{
		log.Fatal(err)
	}
	http.Serve(listen,nil)
}