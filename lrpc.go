package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
	反射方法调用
	1.接口变量->反射对象：Value()
	2.获取对应的方法调用 MethodByName()
	3.Call()
	 */

	s1:=student{"abc",90}
	value:=reflect.ValueOf(s1)
	fmt.Println(value)
	methodValue1:=value.MethodByName("PrintInfo")
	fmt.Println(methodValue1.Kind(),methodValue1.Type())
	
}

type student struct {
	name string
	age int
}

func (s student)say(msg string)  {
	fmt.Println(msg)
}

func (s student)PrintInfo()  {
	fmt.Println(s)
}
