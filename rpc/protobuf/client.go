package main

import (
	"awesomeProject/rpc/protobuf/person"
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.DialHTTP("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	request := person.Request{
		Id:        1,
		Timestamp: time.Now().Unix(),
	}
	var resp *person.Person

	err = client.Call("PersonService.GetPersonInfo", request, &resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*resp)


	call := client.Go("PersonService.GetPersonInfo", request, &resp,nil)
	replyDone := <-call.Done
	fmt.Println(replyDone)
	fmt.Println(*resp)

}
