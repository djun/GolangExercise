package main

import (
	"fmt"
	"reflect"
)

type TestStruct struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Hello, 世界")

	fmt.Println()

	t := TestStruct{"test", 123}
	fmt.Println(reflect.TypeOf(t))

	fmt.Println()

	v := reflect.ValueOf(t)
	fmt.Println(reflect.TypeOf(v))
	fmt.Println(v)
	fmt.Println(v.Type())
	fmt.Println(v.Kind())

	fmt.Println()

	// fmt.Printf("%T\n", t)
	// fmt.Printf("%v\n", t)

	// fmt.Println()

	t2 := v.Interface()
	fmt.Println(reflect.TypeOf(t2))
	fmt.Println(t2)

	fmt.Println()

	t3 := t2.(TestStruct)
	fmt.Println(reflect.TypeOf(t3))
	fmt.Println(t3)

	fmt.Println()
}
