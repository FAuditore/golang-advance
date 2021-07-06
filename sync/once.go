package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	once.Do(func() {
		fmt.Println(1)
	})
	once.Do(func() {
		fmt.Println(1)
	})

}
