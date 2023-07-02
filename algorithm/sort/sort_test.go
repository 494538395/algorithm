package sort

import (
	"fmt"
	"testing"
)

// 排序算法

// 冒泡排序
func bubble(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

func selectSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	for i := 0; i < len(arr)-1; i++ {
		min := arr[i]
		minIndex := -1
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				minIndex = j
			}
		}
		if minIndex != -1 {
			temp := arr[i]
			arr[i] = min
			arr[minIndex] = temp
		}

	}
	return arr
}

func insertSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		cur := arr[i]
		j := i - 1

		for j >= 0 && cur < arr[j] {
			arr[j+1] = arr[j]
			j = j - 1
		}

		arr[j+1] = cur

	}
	return arr
}

func shellSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	length := len(arr)

	for gap := length / 2; gap > 0; gap /= 2 {

		for i := gap; i < length; i++ {
			j := i
			temp := arr[i]
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
	}
	return arr
}

func quickSort(arr []int, left, right int) []int {
	if left >= right {
		return arr
	}
	pivot := arr[left]

	l, r := left, right

	for l < r {
		// 寻找比 pivot 小的数字
		for l < r && arr[r] > pivot {
			r--
		}
		// 寻找比 pivot 大的数字
		for l < r && arr[l] < pivot {
			l++
		}
		arr[l], arr[r] = arr[r], arr[l]
	}
	arr[l] = pivot

	quickSort(arr, left, l-1)
	quickSort(arr, l+1, right)

	return arr
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) >> 1

	left := mergeSort(arr[0:mid])
	right := mergeSort(arr[mid:])

	arr = merge(left, right)

	return arr
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
		res = append(res, right[r:]...)
	}
	return res
}

func Test1(t *testing.T) {
	//arr := []int{1, 2, -3, 9, -4, 6, -8}
	arr := []int{4, 2, 3, 6}
	//arr := []int{19, 97, 9, 17, 1, 8}

	//fmt.Println("bubble(arr)-->", bubble(arr))
	//
	//fmt.Println("select(arr)-->", selectSort(arr))

	//fmt.Println("insert(arr)-->", insertSort(arr))

	//fmt.Println("select(arr)-->", shellSort(arr))

	//quickSort(arr, 0, len(arr)-1)
	//fmt.Println("select(arr)-->", arr)

	fmt.Println("mergeSort(arr)-->", mergeSort(arr))

}
