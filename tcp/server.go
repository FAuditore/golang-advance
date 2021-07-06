package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8999")
	handleErr(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	handleErr(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		read := make([]byte, 1024)
		conn.Read(read)
		fmt.Println(string(read))
		fmt.Println(conn.RemoteAddr())
		conn.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
		conn.Close()
	}
}
func handleErr(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}
