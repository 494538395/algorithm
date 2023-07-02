package _0230624

import (
	"fmt"
	"sync"
	"testing"
)

// 使用两个 goroutine 打印 1-100 之间的数，一个打印奇数，一个打印偶数，打印要交叉打印
// 考点：channel，无缓冲 channel 控制顺序

// 解法1
func TestPrintHundred1(t *testing.T) {
	ch := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)

	// 打印奇数
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- struct{}{}
			if i%2 != 0 {
				fmt.Println("打印奇数  i-->", i)
			}
		}
	}()

	// 打印偶数
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			<-ch
			if i%2 == 0 {
				fmt.Println("打印偶数数  i-->", i)
			}
		}
	}()

	wg.Wait()

}

// 解法2
func TestPrintHundred2(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	oddNum, evenNum := make(chan int, 100), make(chan int, 100)
	oddNum <- 1

	// 打印奇数
	// 1 3 5
	// <- 2 4 6
	go func() {
		defer wg.Done()
		for num := range oddNum {

			fmt.Println("打印奇数  num-->", num)
			num++
			evenNum <- num
			if num == 100 {
				close(oddNum)
				break
			}
		}
	}()

	// 2 4
	// <-3 5
	// 打印偶数
	go func() {
		defer wg.Done()
		for num := range evenNum {
			if num == 100 {
				fmt.Println("打印偶数  num-->", num)
				close(evenNum)
				break
			}
			fmt.Println("打印偶数  num-->", num)
			num++
			oddNum <- num
		}
	}()

	wg.Wait()

}

//--------------------------------------------------------------------------------------------------------------------------

// 排序算法
// 冒泡排序
func TestBubble(t *testing.T) {
	arr := []int{8, 4, 1, 9, -10, 7, -2}

	func(arr []int) []int {
		if len(arr) < 2 {
			return arr
		}

		for i := 0; i < len(arr)-1; i++ {
			for j := i + 1; j < len(arr); j++ {
				if arr[i] > arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}
		return arr
	}(arr)

	fmt.Println(arr)
}

// 选择排序
// 每一次遍历把最小的放在最左边
func TestSelect(t *testing.T) {
	arr := []int{8, 4, 1, 9, -10, 7, -2}
	selectSort := func(arr []int) []int {
		if len(arr) < 2 {
			return arr
		}

		for i := 0; i < len(arr)-1; i++ {
			min := arr[i]
			minIdx := i
			for j := i + 1; j < len(arr); j++ {
				if min > arr[j] {
					min = arr[j]
					minIdx = j
				}
			}
			if minIdx != i {
				arr[minIdx], arr[i] = arr[i], arr[minIdx]
			}

		}
		return arr
	}
	selectSort(arr)
	fmt.Println(arr)
}

// 插入排序
// 从下标为 1 的元素开始遍历
// （ 下标为 1 的左边）认为之前的已经是一个有序序列，把 cur 当前插入之前的有序序列
func TestInsert(t *testing.T) {
	arr := []int{8, 4, 1, 9, -10, 7, -2}
	insert := func(arr []int) []int {
		if len(arr) < 2 {
			return arr
		}

		for i := 1; i < len(arr); i++ {
			cur := arr[i]
			j := i - 1

			for j >= 0 && arr[j] > cur {
				arr[j+1] = arr[j]
				j--
			}
			arr[j+1] = cur

		}
		return arr
	}
	insert(arr)
	fmt.Println(arr)
}

// 希尔排序
// 通过 gap 来实现插入排序
func TestShell(t *testing.T) {
	arr := []int{8, 4, 1, 9, -10, 7, -2}
	shell := func(arr []int) []int {
		if len(arr) < 2 {
			return arr
		}
		for gap := len(arr) >> 1; gap > 0; gap >>= 1 {
			for i := 0; i < len(arr); i++ {
				cur := arr[i]
				j := i
				for j >= gap && arr[j-gap] > cur {
					arr[j] = arr[j-gap]
					j -= gap
				}
				arr[j] = cur
			}
		}
		return arr
	}
	shell(arr)
	fmt.Println(arr)

}

// 快速排序
func quick(arr []int, left, right int) []int {
	if left >= right {
		return arr
	}
	pivot := arr[left]

	l, r := left, right
	for l < r {
		// 寻找小于 pivot 的数
		for l < r && arr[r] > pivot {
			r--
		}
		// 寻找大于 pivot 的数
		for l < r && arr[l] < pivot {
			l++
		}
		arr[l], arr[r] = arr[r], arr[l]
	}
	arr[l] = pivot

	quick(arr, left, l-1)
	quick(arr, l+1, right)
	return arr
}

func TestQuick(t *testing.T) {
	arr := []int{8, 4, 1, 9, -10, 7, -2}
	quick(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// 归并排序
// 分治
// 以中间开始分租，

func divide(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) >> 1

	left := divide(arr[0:mid])
	right := divide(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	l_length, r_length := len(left), len(right)
	l, r := 0, 0
	res := make([]int, 0, l_length+r_length)

	for l < l_length && r < r_length {
		if left[l] < right[r] {
			res = append(res, left[l])
			l++
		} else {
			res = append(res, right[r])
			r++
		}
	}
	if l < l_length {
		res = append(res, left[l:]...)
	}
	if r < r_length {
		res = append(res, right[r:]...,
		)
	}
	return res
}

func TestMerge(t *testing.T) {
	arr := []int{8, 4, 1, 9}
	ints := divide(arr)
	fmt.Println(ints)
}
