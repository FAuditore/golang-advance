package main

/*
#import <stdlib.h>
#import <stdio.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(int(C.random()))

	cs := C.CString("abc")
	defer C.free(unsafe.Pointer(cs))
	C.fputs(cs, (*C.FILE)(C.stdout))
}
