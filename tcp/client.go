package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	tcpAddr,err:=net.ResolveTCPAddr("tcp","localhost:8999")
	if err!=nil {
		log.Fatal(err.Error())
	}
	conn,err:=net.DialTCP("tcp",nil,tcpAddr)
	if err!=nil {
		log.Fatal(err.Error())
	}
	conn.Write([]byte("helloServer"))
	result,err:=ioutil.ReadAll(conn)
	if err!=nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(result))


}



