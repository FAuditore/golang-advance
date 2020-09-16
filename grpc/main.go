package main

import (
	person "awesomeProject/grpc/protobuf"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	test := &person.Person{
		Id:     0,
		Name:   "abc",
		Scores: []int32{60, 100},
		Gender: person.Person_FEMALE,
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &person.Person{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println(newTest)
}
