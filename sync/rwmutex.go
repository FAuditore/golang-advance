package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex
var wg *sync.WaitGroup
func main() {
	rwMutex = new(sync.RWMutex)
	wg = new(sync.WaitGroup)
	wg.Add(2)
	go read()
	go read()
	wg.Wait()
	fmt.Println("finfish")
}
func read(){
	rwMutex.RLock()
	fmt.Println("read")
	time.Sleep(1*time.Second)
	rwMutex.RUnlock()
	defer wg.Done()
}
