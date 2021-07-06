package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
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

	err = channel.ExchangeDeclare("logs_direct", "direct", true, true, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	q, err := channel.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [info] [warning] [error]", os.Args[0])
		os.Exit(0)
	}
	for _, s := range os.Args[1:] {
		err = channel.QueueBind(q.Name, s, "logs_direct", false, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	delivery, err := channel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range delivery {
			log.Printf("Receive message : %s\n", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
