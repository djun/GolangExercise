package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

// 通用的比较接口
type IComparable interface {
	LessEqual(interface{}) bool
}

// 制作一个带比较接口的Int64
type Int64 int64

func (a Int64) LessEqual(b interface{}) bool {
	if bv, ok := b.(Int64); ok {
		return a <= bv
	}
	panic("Not supported type!")
}

func Sort(arr []interface{}) {
	wg.Add(1)

	defer wg.Done()

	if len(arr) <= 1 {
		return
	}

	i, j, ki := 0, len(arr)-1, arr[0]
	key, _ := ki.(IComparable)
	for i < j {
		for i < j {
			if aj, _ := arr[j].(IComparable); key.LessEqual(aj) {
				j--
			} else {
				break
			}
		}
		arr[i] = arr[j]
		for i < j {
			if ai, _ := arr[i].(IComparable); !(key.LessEqual(ai)) {
				i++
			} else {
				break
			}
		}
		arr[j] = arr[i]
	}

	arr[i] = key
	// fmt.Println(i, key, arr)
	if i > 0 {
		go Sort(arr[:i])
	}
	if i+1 < len(arr) {
		go Sort(arr[i+1:])
	}
}

func main() {
	// 产生随机数组成的数组
	size := 1000000
	rand.Seed(time.Now().UnixNano())
	arr := make([]interface{}, 0)
	for i := 0; i < size; i++ {
		arr = append(arr, Int64(rand.Intn(size)))
	}

	// 准备排序
	fmt.Printf("%v\n", arr)

	// 进行排序，并计时
	start := time.Now()
	Sort(arr)
	wg.Wait()
	fmt.Printf("sort time: %d ms\n", time.Since(start).Milliseconds())

	// time.Sleep(1e9)
	// 看排序结果
	fmt.Println(arr)

	// 验证排序结果是否准确
	for ni := range arr {
		if ni > 0 {
			last, _ := arr[ni-1].(int64)
			curr, _ := arr[ni].(int64)
			if !(last <= curr) {
				fmt.Println("failed!")
				return
			}
		}
	}
	fmt.Println("successful!")
}
