package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	t := time.Now()
	go func() {
		client := &http.Client{}
		request, _ := http.NewRequest("GET", "http://localhost:8999/hello", nil)
		response, _ := client.Do(request)
		result, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(result))
		wg.Done()
	}()
	go func() {
		client := &http.Client{}
		request, _ := http.NewRequest("GET", "http://localhost:8999/hello", nil)
		response, _ := client.Do(request)
		result, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(result))
		wg.Done()
	}()
	go func() {
		client := &http.Client{}
		request, _ := http.NewRequest("GET", "http://localhost:8999/hello", nil)
		response, _ := client.Do(request)
		result, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(result))
		wg.Done()
	}()
	wg.Wait()
	t2 := time.Since(t)
	fmt.Println(t2)
}
