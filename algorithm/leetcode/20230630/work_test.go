package _0230630

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

/*
	六月的最后一天。离职半个月了。也准备挺久了，接下来，就是真正的战斗
*/

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 46. 把数字翻译成字符串
func translateNum(num int) int {
	// example : []int{ 1,2,2,5,8 }
	// dp[i] 表示以 i 结尾的字符串所能组成最多的翻译数
	// ascii 48==0  49==1

	itoa := strconv.Itoa(num)

	dp := make([]int, len(itoa))
	dp[0] = 1

	for i := 1; i < len(dp); i++ {
		dp[i] = dp[i-1]
		// 看看当前的数字是否能和前面进行组合
		cur := itoa[i]
		pre := itoa[i-1]
		num = int((pre-48)*10 + (cur - 48))
		if num > 9 && num < 26 {
			// 当前数字和前一位的数字可以进行组合
			if i >= 2 {
				dp[i] += dp[i-2]
			} else {
				dp[i] += 1
			}
		}

	}
	return dp[len(dp)-1]
}
func TestTranslateNum(t *testing.T) {
	fmt.Println(translateNum(12258))
}

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 48. 最长不含重复字符的子字符串
func lengthOfLongestSubstring(s string) int {
	if s == "" || len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	dic := make(map[uint8]int)
	dic[s[0]] = 0
	slow, fast := 0, 1
	ans := 1

	for slow < len(s) && fast < len(s) {
		// 如果找到，说明该调整窗口了
		if isExist(dic, s[fast]) {
			idx, _ := dic[s[fast]]
			if slow < idx+1 {
				slow = idx + 1
			}
		}

		dic[s[fast]] = fast
		ans = max(ans, fast-slow+1)
		fast++
	}
	return ans
}
func isExist(dic map[uint8]int, key uint8) bool {
	_, found := dic[key]
	return found
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abba"))
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 45. 把数组排成最小的数
func minNumber(nums []int) string {
	if nums == nil || len(nums) == 0 {
		return ""
	}

	sort.Slice(nums, func(i, j int) bool {
		x := fmt.Sprintf("%d%d", nums[i], nums[j])
		y := fmt.Sprintf("%d%d", nums[j], nums[i])
		return x < y
	})

	ans := ""

	for _, num := range nums {
		ans += fmt.Sprint(num)
	}

	return ans
}
func TestMinNumber(t *testing.T) {
	fmt.Println(minNumber([]int{3, 30, 34, 5, 9}))
	str := []string{"3", "30", "300", "3", "9", "40"}
	sort.Strings(str)
	fmt.Println(str)
}
func TestS(t *testing.T) {
	str := "102"

	strings.Split(str, "")

}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 61. 扑克牌中的顺子
func isStraight(nums []int) bool {
	sort.Ints(nums)
	zero_count := 0
	for _, num := range nums {
		if num != 0 {
			break
		}
		zero_count++
	}

	if zero_count == len(nums) {
		return true
	}

	for i := zero_count; i < len(nums); i++ {
		if i+1 < len(nums) {
			gap := nums[i+1] - nums[i]
			if gap == 0 {
				return false
			}
			if gap > 1 {
				gap -= 1
				zero_count -= gap
				if zero_count < 0 {
					return false
				}
			}
		}
	}
	return true
}

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 40. 最小的k个数
func getLeastNumbers(arr []int, k int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr[:k]
}
func quickSort(arr []int, left, right int) []int {
	if left >= right {
		return arr
	}

	pivot := arr[left]
	l, r := left, right

	for l < r {
		// 找到右指针第一个小于 pivot 的值
		for l < r && arr[r] >= pivot {
			r--
		}
		// 找到左指针第一个大于 pivot 的值
		for l < r && arr[l] <= pivot {
			l++
		}
		arr[l], arr[r] = arr[r], arr[l]
	}
	arr[l], arr[left] = arr[left], arr[l]
	quickSort(arr, left, l-1)
	quickSort(arr, l+1, right)
	return arr
}
func TestQuickSort2(t *testing.T) {
	arr := []int{1, 2, 1}
	quickSort(arr, 0, 2)
	fmt.Println(arr)
}
