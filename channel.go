package main

import "fmt"

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
	*/
	var a chan int
	a = make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("zi goroutine", i)
		}
		a <- 1
	}()
	data := <-a
	fmt.Println(data)

	ch2 := make(chan int)
	go senddata(ch2)
	for {
		v, ok := <-ch2
		if ok {
			fmt.Println(v, ok)
		} else {	//channel close之后，读取会返回0值和false
			break
		}
	}

	ch3 := make(chan int)
	go senddata2(ch3)

	for v := range ch3 {
		fmt.Println(v)
	}
}

func senddata(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func senddata2(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i * 2
	}
	close(c)
}
