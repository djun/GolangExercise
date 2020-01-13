package main

import (
	"fmt"
	"math/rand"
	"time"
)


type IComparable interface {
	LessEqual(interface{}) bool
}

type ISortable interface {
	Sort()
}


type Integer int

func (a Integer) LessEqual(b interface{}) bool {
	k, _ := b.(Integer)
	return a <= k
}

type Float float64

func (a Float) LessEqual(b interface{}) bool {
	k, _ := b.(Float)
	return a <= k
}


type SortableArr struct {
	Arr []IComparable
	ISortable
}

func (a SortableArr) Sort() {
	// 采用快速排序算法（递归版本）

	if len(a.Arr) <= 1 {
		return
	}

	i, j, key := 0, len(a.Arr) - 1, a.Arr[0]
	for i < j {
		for i < j && key.LessEqual(a.Arr[j]) { j-- }
		a.Arr[i] = a.Arr[j]
		for i < j && !key.LessEqual(a.Arr[i]) { i++ }
		a.Arr[j] = a.Arr[i]
	}

	a.Arr[i] = key
	// fmt.Println(i, key, a.Arr)
	if i > 0 {
		go (SortableArr{ Arr: a.Arr[: i] }).Sort()
	}
	if i + 1 < len(a.Arr) {
		go (SortableArr{ Arr: a.Arr[i+1:] }).Sort()
	}
}


func main() {
	// 产生随机数组成的数组
	const size = 1000000
	rand.Seed(time.Now().UnixNano())
	arr := make([]IComparable, 0)
	for i:=0; i<size; i++ {
		arr = append(arr, Integer(rand.Intn(size)))
	}

	// 准备排序
	s := &SortableArr { Arr: arr }
	fmt.Println(s.Arr)

	// 进行排序，并计时
	start := time.Now()
	s.Sort()
	fmt.Printf("sort time: %s\n", time.Since(start))
	
	time.Sleep(1e9)
	// 看排序结果
	fmt.Println(s.Arr)
	
	// 验证排序结果是否准确
	for ni, _ := range arr {
		if ni > 0 {
			if !arr[ni-1].LessEqual(arr[ni]) {
				fmt.Println("failed!")
				return
			}
		}
	}
	fmt.Println("successful!")
}