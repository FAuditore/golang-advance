package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.NumCPU())

	//gosched
	go func() {
		for i:=0;i<5;i++{
			fmt.Println("Goroutine")
			runtime.Gosched()
		}
	}()
	for i:=0;i<10;i++{
		runtime.Gosched()
		fmt.Println("main ",i)
	}
}
