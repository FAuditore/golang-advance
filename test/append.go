package test

import "strconv"

func Append() {
	var a []string
	for i := 0; i < 180; i++ {
		a = append(a, strconv.Itoa(i))
	}
}
