package main

import "fmt"

func main() {
	/*
	双向通道
		chan <- data 发送数据 写出
		data <- chan 获取数据 读取
	单向通道
		chan <- T 只能写
		<- chan T 只能读
	 */
	ch1 := make(chan string)
	done :=make(chan bool)
	go sendData(ch1,done)

	//双向通道
	data:=<-ch1
	fmt.Println(data)
	ch1<-"bcd"
	<-done
	fmt.Println("over")

	//ch2:=make(chan <-int)//只能写 不能读
	//ch3 := make(<-chan int)//只能读 不能写

}

func sendData(c chan string,b chan bool){
	c<-"abc" //发送
	data:=<-c //接收
	fmt.Println(data)
	b<-true
}