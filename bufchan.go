package main

import "fmt"

func main() {
	/*
	缓冲通道 通道满了阻塞
	make(chan type,capacity)
	*/
	ch1:=make(chan int,5)
	ch1 <- 100
	fmt.Println(len(ch1),cap(ch1)) // 1,5

}
