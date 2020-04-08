package main

import "fmt"

func main() {
	m := make(map[interface{}]interface{})
	m["test"] = "hello world!"
	m[1] = 1000
	for k, v := range m {
		fmt.Printf("%v = %v\n", k, v)
	}
	for k := range m {
		fmt.Printf("%v\n", k)
	}
}
