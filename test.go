package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	go func() {
		fmt.Println("start")
		for {
			v, ok := <-done
			fmt.Println(v, ok)
		}

	}()
	time.Sleep(1 * time.Second)
	close(done)
	time.Sleep(2 * time.Second)
}
