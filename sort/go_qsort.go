package main

import (
	"fmt"
	"math/rand"
	"time"
)

type SortableArr struct {
	Arr []int64
}

func (a SortableArr) Sort() {
	// 采用快速排序算法（递归版本）

	if len(a.Arr) <= 1 {
		return
	}

	i, j, key := 0, len(a.Arr) - 1, a.Arr[0]

	for i < j {
		for i < j && key <= a.Arr[j] { j-- }
		a.Arr[i] = a.Arr[j]
		for i < j && key >= a.Arr[i] { i++ }
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
	rand.Seed(time.Now().UnixNano())
	arr := [1000000]int64{}
	for ni, _ := range arr {
		arr[ni] = int64(rand.Intn(len(arr)))
	}

	// 准备排序
	SortableArr := SortableArr { Arr: arr[:] }
	fmt.Println(SortableArr.Arr)

	// 进行排序，并计时
	start := time.Now()
	SortableArr.Sort()
	fmt.Printf("sort time: %s\n", time.Since(start))
	
	time.Sleep(1e9)
	// 看排序结果
	fmt.Println(SortableArr.Arr)
	
	// 验证排序结果是否准确
	for ni, _ := range arr {
		if ni > 0 {
			if arr[ni-1] > arr[ni] {
				fmt.Println("failed!")
				return
			}
		}
	}
	fmt.Println("successful!")
}