package main

import (
	"advance/rpc/protobuf/person"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type PersonService struct {
}

var (
	person1 = person.Person{
		Id:     1,
		Name:   "aaaa",
		Scores: []int32{2, 3, 4},
		Gender: 1,
	}
	person2 = person.Person{
		Id:     2,
		Name:   "bbbb",
		Scores: []int32{95, 98, 100},
		Gender: 0,
	}
)

func (ps *PersonService) GetPersonInfo(request person.Request, response *person.Person) error {
	fmt.Println("request:", request)
	current := time.Now().Unix()
	if current < request.Timestamp {
		response = nil
		return errors.New("查询异常")
	}
	if request.Id == 1 {
		response = &person1
	} else {
		*response = person2
	}
	return nil
}
func main() {
	err := rpc.Register(&PersonService{})
	if err != nil {
		log.Fatal(err)
	}
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	http.Serve(listener, nil)

}
