package main

import (
	"fmt"
	"time"
)

type T struct {
	Id   int
	Name string
}

func (t *T) test() {
	fmt.Println(t)
}


func main() {
	fmt.Println(HasFunc(&T{}))
}
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false

	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}


func HasFunc(a interface{}) bool {
	_, ok := a.(interface{ test(int) string })
	return ok
}
