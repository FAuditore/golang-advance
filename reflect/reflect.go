package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type person struct {
	name string
	Age  int
}

func main() {
	/*
		value struct{
			typ *rtype
			ptr unsafe.Pointer
			flag
		}
		Type interface{
			Align()
			NumMethod()
			Name()
			...
		}
	*/
	var a person
	fmt.Println(reflect.ValueOf(a).Kind())
	fmt.Println(reflect.ValueOf(a).NumField())
	fmt.Println(reflect.ValueOf(a).Type().Field(0).Name)
	var b int
	var c *int
	c = &b
	fmt.Println(reflect.ValueOf(c).Kind())
	fmt.Printf("%v\n", reflect.ValueOf(c).Elem().Type())
	fmt.Printf("%v\n", reflect.ValueOf(c).Elem())

	fmt.Println(reflect.ValueOf(a).Type().NumMethod())
	fmt.Println(reflect.ValueOf(a).Type().Method(0).Name)
	fmt.Println(reflect.ValueOf(formatAtom).Call([]reflect.Value{reflect.ValueOf(reflect.ValueOf(a))}))



	fmt.Println(reflect.ValueOf(c).CanSet())
	fmt.Println(reflect.ValueOf(c).Elem().CanSet())
	fmt.Println(reflect.Indirect(reflect.ValueOf(c)).CanSet())
	reflect.ValueOf(c).Elem().Set(reflect.ValueOf(777))
	fmt.Println(*c)

	fmt.Println(reflect.ValueOf([]string{"abc"}).Index(0).CanAddr())
}

func (a person) Abcdefg() {
	fmt.Println("AAAAAAA")
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)

	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}
