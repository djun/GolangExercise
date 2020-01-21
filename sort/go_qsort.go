package main

import (
	"fmt"
	"sync"
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

func Sort(x SortableArr, wg *sync.WaitGroup) {
	wg.Add(1)

	if len(x.Arr) <= 1 {
		return
	}

	i, j, key := 0, len(x.Arr) - 1, x.Arr[0]
	for i < j {
		for i < j && key.LessEqual(x.Arr[j]) { j-- }
		x.Arr[i] = x.Arr[j]
		for i < j && !key.LessEqual(x.Arr[i]) { i++ }
		x.Arr[j] = x.Arr[i]
	}

	x.Arr[i] = key
	// fmt.Println(i, key, x.Arr)
	if i > 0 {
		go Sort(SortableArr{ Arr: x.Arr[: i] }, wg)
	}
	if i + 1 < len(x.Arr) {
		go Sort(SortableArr{ Arr: x.Arr[i+1:] }, wg)
	}

	wg.Done()
}

func (a SortableArr) RunSort() {
	// 采用快速排序算法（递归版本）

	var waitGroup sync.WaitGroup
	go Sort(a, &waitGroup)
	waitGroup.Wait()
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
	s.RunSort()
	fmt.Printf("sort time: %d ms\n", time.Since(start).Milliseconds())
	
	// time.Sleep(1e9)
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