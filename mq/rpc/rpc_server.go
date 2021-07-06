package main

import (
	"github.com/streadway/amqp"
	"log"
	"strconv"
)

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()

	q, err := channel.QueueDeclare("rpc_queue", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = channel.Qos(1, 0, false)
	if err != nil {
		log.Fatal(err)
	}

	delivery, err := channel.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range delivery {
			n, err := strconv.Atoi(string(d.Body))
			if err != nil {
				log.Fatal(err)
			}
			log.Printf(" [.] fib(%d)", n)
			response := fib(n)
			err = channel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
				ContentType:   "text/plain",
				CorrelationId: d.CorrelationId,
				Body:          []byte(strconv.Itoa(response)),
			})
			if err != nil {
				log.Fatal(err)
			}
			d.Ack(false)
		}
	}()
	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}
