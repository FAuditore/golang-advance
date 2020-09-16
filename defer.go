package main

import "fmt"

func main() {
	i := increaseA()
	fmt.Println(i)


	r := increaseB()
	fmt.Println(r)

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
