package main

import "fmt"

type Slice []int

func (A Slice) Append(value int) {
	A1 := append(A, value)
	fmt.Println(A, A1)
	fmt.Printf("%p\n%p\n", A, A1)

	A = A1
}

func main() {
	mSlice := make(Slice, 10, 20)
	fmt.Println(mSlice)
	mSlice.Append(5)
	fmt.Println(mSlice)
}
