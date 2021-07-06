package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var a int32 = 0
	for i := 0; i < 1000; i++ {
		go func() {
			atomic.AddInt32(&a, 1)
		}()
	}
	time.Sleep(200 * time.Millisecond)
	fmt.Println(a)

	var v atomic.Value
	var m = 0
	go func() {
		for i := 0; ; i++ {
			m = i
			v.Store(m)
			time.Sleep(20 * time.Millisecond)
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				m := v.Load().(int)
				fmt.Println(m)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
