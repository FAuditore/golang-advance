package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a [4]int
	var b string
	var c struct{}
	fmt.Println(
		unsafe.Sizeof(a),
		unsafe.Sizeof(b),
		unsafe.Sizeof(c))

	fmt.Println(
		unsafe.Sizeof(struct {
			bool
			float64
			int16
		}{}),
		unsafe.Sizeof(struct {
			float64
			int16
			bool
		}{}),
		unsafe.Sizeof(struct {
			bool
			int16
			float64
		}{}))

	var f float64 = 1.0
	fmt.Printf("%p\n", &f)
	fmt.Println(unsafe.Pointer(&f))
	fmt.Printf("%#016x\n", *(*uint64)(unsafe.Pointer(&f)))
	var x struct {
		a bool
		b int16
		c []int
	}

	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b)
	tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	pb = (*int16)(unsafe.Pointer(tmp))
	*pb = 42
}
