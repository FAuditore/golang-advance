package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		声明通道
			var [name] chan [type]
			[name] = make(chan [type])

		读取数据
			[data] := <-[chanName]
		发送数据
			[chanName] <- [value]
		一次发送一次阻塞


		缓冲通道 通道满了阻塞
		make(chan type,capacity)

		双向通道
			chan <- data 发送数据 写出
			data <- chan 获取数据 读取
		单向通道
			chan <- T 只能写
			<- chan T 只能读
	*/
	var a chan int
	a = make(chan int)

	go func() {
		a <- 1
	}()
	data := <-a
	fmt.Println(data)

	ch1:=make(chan int,5)
	ch1 <- 100
	fmt.Println(len(ch1),cap(ch1)) // 1,5



	ch2 := make(chan int)
	go sendData(ch2)

	for v := range ch2 {
		fmt.Println(v)
	}


	var ch = make(chan int)
	//只写channel
	go func(ch chan<- int) {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}(ch)
forLoop:
	for {
		select {
		case i, ok := <-ch:
			if ok {
				fmt.Println("receive ", i)
			}else {
				fmt.Println("channel closed")
				break forLoop
			}
		case <-time.After(time.Second):
			fmt.Println("Time out")
			break forLoop
		}
	}
}

func sendData(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i * 2
	}
	close(c)
}
