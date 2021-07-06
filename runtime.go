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
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine")
			runtime.Gosched()
		}
	}()
	for i := 0; i < 10; i++ {
		runtime.Gosched()
		fmt.Println("main ", i)
	}

	call()

}

func call() {
	/*
	   Caller报告当前go程调用栈所执行的函数的文件和行号信息。
	   实参skip为上溯的栈帧数，0表示Caller的调用者（Caller所在的调用栈）。
	   （由于历史原因，skip的意思在Caller和Callers中并不相同。）
	   函数的返回值为调用栈标识符、文件名、该调用在文件中的行号。
	   如果无法获得信息，ok会被设为false。
	*/

	fmt.Println(runtime.Caller(0))

	pc := make([]uintptr, 100)
	runtime.Callers(0, pc)
	for _, v := range pc {
		fun := runtime.FuncForPC(v)
		file, line := fun.FileLine(v)
		fmt.Println(file, line)
	}
}
