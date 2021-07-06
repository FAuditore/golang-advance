package main

import (
	"fmt"
)

func main() {
	/*
		类似switch
			随机执行一个满足的case(每个case都是通道操作)
			没有case，执行default，否则进入阻塞，等待case满足
	*/

	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- 100
	}()
	go func() {
		ch2 <- 200
	}()
	select {
	case num1 := <-ch1:
		fmt.Println("ch1: ", num1)
	case num2, ok := <-ch2:
		if ok {
			fmt.Println("ch2: ", num2)
		} else {
			fmt.Println(ch2)
		}
	default:
		fmt.Println("123")
	}
	fmt.Println("1111")

}
