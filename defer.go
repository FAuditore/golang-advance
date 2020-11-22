package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func main() {
	/*
		1.12
		defer注册: deferproc把要执行的函数信息保存起来
		defer执行: deferrun
		defer信息注册到一个链表,goroutine有这个链表的头指针 新注册的defer添加到链表头
		1.13
		增加局部变量,defer信息保存到当前栈帧局部变量区域 避免了堆分配
		1.14
		编译阶段插入代码,把defer函数执行直接展开在所属函数内,不需要注册defer链表 但是在panic发生后需要栈扫描恢复defer链表
	 */
	i := increaseA()
	fmt.Println(i)


	r := increaseB()
	fmt.Println(r)

	bigSlowOperation()


	foo(2, 0)
}

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func bigSlowOperation(){
	defer trace("bigSlowFunction")()
	time.Sleep(5*time.Second)
}

func trace(name string) func(){
	start:=time.Now()
	log.Printf("enter: %s ",name)
	return func() {
		log.Printf("exit: %s %s",name,time.Since(start))
	}
}

func foo(a, b int) (i int, err error) {
	defer fmt.Printf("first defer err %v\n", err)
	defer func(err error) { fmt.Printf("second defer err %v\n", err) }(err)
	defer func() { fmt.Printf("third defer err %v\n", err) }()
	if b == 0 {
		err = errors.New("divided by zero!")
		return
	}

	i = a / b
	return
}