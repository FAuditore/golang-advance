package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.OpenFile("trace.out", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	defer trace.Stop()

	fmt.Println("hi")
}
