package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float32 = 1.5
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.ValueOf(x))
	var A interface{}
	A=1.5
	fmt.Println(reflect.TypeOf(A))
	fmt.Println(reflect.ValueOf(A))
	A="stringA"
	fmt.Println(reflect.TypeOf(A))
	fmt.Println(reflect.ValueOf(A))

	fmt.Println()

	value:=reflect.ValueOf(A)
	convertValue:=value.Interface().(string)
	fmt.Println(convertValue)

	var B float64 = 1.5
	pointer := reflect.ValueOf(&B)
	convertValue2 := pointer.Interface().(*float64)
	fmt.Println(convertValue2)


	p1 := person{"Abc",8}
	getMessage(p1)


	var num float64 = 1.23
	fmt.Println("num: ",num)

	nump:=reflect.ValueOf(&num)
	numv:=nump.Elem()
	numv.SetFloat(1)
	fmt.Println(num)



	s1:=person{"bbb",100}
	vs1:=reflect.ValueOf(&s1)
	fmt.Println(vs1)
	newValue := vs1.Elem()
	fmt.Println(newValue)
	f1:=newValue.FieldByName("Name")
	f1.SetString("ccc")
}

type person struct {
	Name string
	age int
}

func (p person) Say()  {
	fmt.Println("say ")
}

func getMessage(input interface{}){
	t:=reflect.TypeOf(input)
	fmt.Println(t.Kind())
	fmt.Println(t.Name())
	v :=reflect.ValueOf(input)
	fmt.Println(v.Kind())
	fmt.Println(v)

	/*
	获取字段
	1.获取Type对象 reflect.Type
		NumField
		Field(index)
	2.通过Field()获取每一个Field字段
	3.Interface方法获取数值
	 */
	for i:=0;i<t.NumField();i++{
		field:=t.Field(i)
		value:=v.Field(i)
		fmt.Println(field)
		fmt.Println(value)
	}

	//获取方法
	for i:=0;i<t.NumMethod();i++{
		method := t.Method(i)
		fmt.Println(method.Name)//方法名小写获取不到
	}
}