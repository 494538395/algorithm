package sort

import (
	"fmt"
	"testing"
)

// 排序算法复习

// 冒泡
// 一上一下，低的沉下去，高的浮上来
func bubbleSortReview(arr []int) []int {
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
}

// 插入排序
// 每一轮将最小的放在最左边
func insertSortReview(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	for i := 0; i < len(arr)-1; i++ {
		min := arr[i]
		minIndex := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	return arr
}

// 选择排序
// 假定前一个已经是一个有序序列，就把当前的元素加入前一个有序序列
func selectSortReview(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		j := i - 1
		temp := arr[i]

		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j -= 1
		}
		arr[j+1] = temp

	}
	return arr
}

// 希尔排序
func shellSortReview(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for gap := len(arr) >> 1; gap > 0; gap >>= 1 {
		for i := gap; i < len(arr); i++ {
			cur := arr[i]
			j := i
			for j >= gap && arr[j-gap] > cur {
				arr[j] = arr[j-gap]
				j -= gap
			}
			if j != i {
				arr[j] = cur
			}
		}
	}
	return arr
}

// 快速排序
// 有一个基准点，小于基准点的放在左边，大于基准点的放在右边，让 l==r停止，将 基准点放在下标为 l或者 r的位置
// 此时，基准点左右两边都是一个有序序列
// 然后再对基准点左右序列执行同样的操作
func quickSortReview(arr []int, left, right int) []int {
	if left >= right {
		return arr
	}

	pivot := arr[left]  // 基准点
	l, r := left, right // 左右指针

	for l < r { // 当左右指针重合时，将基准点放在重合的位置
		// 寻找比 pivot 小的数
		for l < r && arr[r] > pivot {
			r--
		}
		// 寻找比 pivot 大的数
		for l < r && arr[l] <= pivot {
			l++
		}
		arr[l], arr[r] = arr[r], arr[l]
	}
	arr[l] = pivot
	quickSortReview(arr, left, l-1)
	quickSortReview(arr, l+1, right)
	return arr
}

func TestQuickSortReview(t *testing.T) {
	arr := []int{1, 2, 1}
	fmt.Println(quickSortReview(arr, 0, 2))
}

// 归并排序
// 分治
func mergeSortReview(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) >> 1
	left := mergeSortReview(arr[0:mid])
	right := mergeSortReview(arr[mid:])
	return mergeReview(left, right)
}

func mergeReview(left, right []int) []int {
	l_length, r_length := len(left), len(right)
	res := make([]int, 0, l_length+r_length)

	l, r := 0, 0
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
		res = append(res, right[r:]...)
	}
	return res
}

func TestIdentify(t *testing.T) {
	//arr := []int{9, 4, 8, 3, 1, -2, 6, -10}
	arr := []int{9, 1, 4, 3}
	//bubbleSortReview(arr)
	//insertSortReview(arr)
	//selectSortReview(arr)
	//shellSortReview(arr)
	//quickSortReview(arr, 0, len(arr)-1)
	mergeArr := mergeSortReview(arr)
	fmt.Println(mergeArr)
}
