package main

import (
	"bytes"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()

	err = channel.ExchangeDeclare("logs", "fanout", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	q, err := channel.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = channel.QueueBind(q.Name, "", "logs", false, nil)
	if err != nil {
		log.Fatal(err)
	}
	forever := make(chan bool)

	msgs, err := channel.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for d := range msgs {
			log.Printf("Receive message : %s\n", d.Body)
			dot_count := bytes.Count(d.Body, []byte("."))
			time.Sleep(time.Duration(dot_count) * time.Second)
			log.Printf("Done")
			//false->只返回一个ack true->所有之前的发送都被确认
		}
	}()

	log.Printf("[*] Waiting for the message ...")
	<-forever
}
