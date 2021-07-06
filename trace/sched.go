package main

import (
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	done := false

	go func() {
		done = true
	}()

	for !done {
		println("not done !")
		runtime.Gosched()
	}

	println("done !")
}
