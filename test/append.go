package test

import "strconv"

func App() {
	var a []string
	a = make([]string, 0, 102)
	for i := 0; i < 100; i++ {
		a = append(a, strconv.Itoa(i))
	}
}
