package main

import (
	"fmt"
	"reflect"
)

const (
	A int = iota
	B
	C
)

type TS struct {
	A, B, C int
}

func (t TS) Print(test string) {
	fmt.Printf("%s %v\n", test, t)
}

func NewDefaultTS() TS {
	return TS{
		A, B, C,
	}
}

func main() {
	t := NewDefaultTS()
	fmt.Printf("%v\n", t)

	vt := reflect.ValueOf(t)
	mPrint := vt.MethodByName("Print")
	args := []reflect.Value{
		reflect.ValueOf("I wanna to test! "),
	}
	mPrint.Call(args)
}
