package main

import (
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(rand.Intn(26) + 65)
	}
	return string(bytes)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	n := bodyFrom(os.Args)

	log.Printf(" [x] Requesting fib(%d)", n)
	res, err := fibonacciRPC(n)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(" [.] Got %d", res)
}

func fibonacciRPC(n int) (res int, err error) {
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

	q, err := channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		log.Fatal(err)
	}
	delivery, err := channel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	corrId := randomString(32)

	err = channel.Publish("", "rpc_queue", false, false, amqp.Publishing{
		ContentType:   "text/plain", //"application/json"
		CorrelationId: corrId,
		ReplyTo:       q.Name,
		Body:          []byte(strconv.Itoa(n)),
	})
	if err != nil {
		log.Fatal(err)
	}

	for d := range delivery {
		if corrId == d.CorrelationId {
			res, err = strconv.Atoi(string(d.Body))
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}
	return
}
func bodyFrom(args []string) int {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "30"
	} else {
		s = strings.Join(args[1:], " ")
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
