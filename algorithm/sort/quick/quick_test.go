package quick

import (
	"fmt"
	"testing"
)

//func quickSort(arr []int, left, right int) []int {
//	if left >= right {
//		return arr
//	}
//
//	pivot := arr[left]
//	l, r := left, right
//
//	for l < r {
//		// 找到右指针第一个小于 pivot 的值
//		for l < r && arr[r] > pivot {
//			r--
//		}
//		// 找到左指针第一个大于 pivot 的值
//		for l < r && arr[l] < pivot {
//			l++
//		}
//		arr[l], arr[r] = arr[r], arr[l]
//	}
//	arr[l] = pivot
//	quickSort(arr, left, l-1)
//	quickSort(arr, l+1, right)
//	return arr
//}

func quickSort(arr []int, left, right int) []int {
	if left >= right {
		return arr
	}

	pivot := arr[left]
	l, r := left, right

	for l < r {
		// 找到右指针第一个小于等于 pivot 的值
		for l < r && arr[r] >= pivot {
			r--
		}
		// 找到左指针第一个大于等于 pivot 的值
		for l < r && arr[l] <= pivot {
			l++
		}
		arr[l], arr[r] = arr[r], arr[l]
	}

	arr[left], arr[l] = arr[l], arr[left] // 将基准点放在正确的位置上

	quickSort(arr, left, l-1)
	quickSort(arr, l+1, right)

	return arr
}

func TestQuick(t *testing.T) {
	arr := []int{1, 2, 1}
	fmt.Println(quickSort(arr, 0, len(arr)-1))
}
