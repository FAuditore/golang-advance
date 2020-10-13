package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadCaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for c := range clients {
				c <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadCaster()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	ch := make(client)
	go write(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "you are " + who
	messages <- who + " is coming"
	entering <- ch
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		messages <- who + " : " + scanner.Text()
	}
	leaving <- ch
	messages <- who + " left"
}

func write(conn net.Conn, ch client) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
