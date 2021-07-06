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

	queue, err := channel.QueueDeclare("task_queue", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	delivery, err := channel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	forever := make(chan bool)

	err = channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	go func() {
		for d := range delivery {
			log.Printf("Receive message : %s\n", d.Body)
			dot_count := bytes.Count(d.Body, []byte("."))
			time.Sleep(time.Duration(dot_count) * time.Second)
			log.Printf("Done")
			//false ->只返回一个ack
			//true ->所有之前的发送都被确认
			d.Ack(false)
		}
	}()
	log.Printf("[*] Waiting for the message ...")
	<-forever
}
