package _0230702

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

/*
	2023.07.01
	星期日，昨晚失眠到 3 点，床上有很多小虫子，该找个时间去晒洗晒洗被子了。
	今天再做一些算法题，明天开始就要结合八股文来复习。
	深圳，我会留下来的。
*/

//--------------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 43. 1～n 整数中 1 出现的次数
func countDigitOne(n int) int {
	// 这道题的本质是：求 1 在 个位上出现的次数 + 1 在十位上出现的次数 + 1 在千位上出现的次数 ......
	// 假设有数字 num：30101592
	// 现在要求在这个数字中，1 在千位上出现的次数
	// 3 0 1 0 1 5 9 2
	//         x
	// 即 x 所在位千位 1 出现的次数，
	// base 设置位 1000 （因为要求在千位上出现 1 的次数，如果是求十位，base 就是 10）
	// x 所在位的值为 cur，cur 的求算数字公式是：num/base/%10--> 30101592 / 1000 % 10 ==1
	// x 左边的数字称为 a，a 的求算公式是：num/base/10--> 30101592 / 10000 / 10 == 3010
	// x 右边的数字称为 b，b 的求算公式是: num % base --> 30101592 % 1000 = 592
	// 因为 cur == 1，所以 x 位上 1 出的次数会有两种情况
	// 情况1： a 的范围是 0000-3009 ,那 b 可以取到 000-999
	// 情况2： a 的值是 3010 ，则 b 可以取到 000-592
	// 所以当 cur ==1 时候，求 1 出现的公式是 a * base + 1 * (b+1)
	// 以此类推：
	// cur >1 时，公式是 (a+1)*base
	// cur <1 时，公式是 a*base

	var base int64
	base = 1
	var ans int64
	ans = 0

	for base <= int64(n) {
		a := int64(n) / base / 10
		cur := int64(n) / base % 10
		b := int64(n) % base

		if cur > 1 {
			ans += (a + 1) * base
		} else if cur == 1 {
			ans += (a * base) + (1*b + 1)
		} else {
			//cur <1
			ans += a * base
		}
		base *= 10

	}

	return int(ans)
}
func TestConDigitOne(t *testing.T) {
	fmt.Println(countDigitOne(12))
}

//--------------------------------------------------------------------------------------------------------------------------------------

//--------------------------------------------------------------------------------------------------------------------------------------
// 1768. 交替合并字符串
// 有点类似于归并排序中合并的思想
func mergeAlternately(word1 string, word2 string) string {
	if word1 == "" {
		return word2
	}
	if word2 == "" {
		return word1
	}
	ans := ""

	l1, l2 := 0, 0

	for l1 < len(word1) && l2 < len(word2) {
		ans += string(word1[l1])
		ans += string(word2[l2])
		l1++
		l2++
	}

	if l1 < len(word1) {
		ans += word1[l1:]
	}
	if l2 < len(word2) {
		ans += word2[l2:]
	}
	return ans
}

//--------------------------------------------------------------------------------------------------------------------------------------

//--------------------------------------------------------------------------------------------------------------------------------------
// 1071. 字符串的最大公因子
// 终止条件：如果 str1 和 str2 长度相等，如果 str1.equals(str2)==true ，返回 str1，否则返回 ""
// 不断递归，进入终止条件
func gcdOfStrings(str1 string, str2 string) string {

	l1_len, l2_len := len(str1), len(str2)
	if l1_len < l2_len {
		return gcdOfStrings(str2, str1)
	}

	if l1_len == l2_len {
		if str1 == str2 {
			return str1
		}
		return ""
	}

	for i := 0; i < l2_len; i++ {
		if str1[i] != str2[i] {
			return ""
		}
	}
	return gcdOfStrings(str1[l2_len:], str2)
}
func TestGcdOfStrings(t *testing.T) {
	fmt.Println(gcdOfStrings("AB", "ABAB"))
}

//--------------------------------------------------------------------------------------------------------------------------------------

//--------------------------------------------------------------------------------------------------------------------------------------
// 1431. 拥有最多糖果的孩子
func kidsWithCandies(candies []int, extraCandies int) []bool {
	most := 0
	for _, candy := range candies {
		most = max(most, candy)
	}

	ans := make([]bool, len(candies))
	for i, candy := range candies {
		if candy+extraCandies >= most {
			ans[i] = true
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//--------------------------------------------------------------------------------------------------------------------------------------

//--------------------------------------------------------------------------------------------------------------------------------------
// 605. 种花问题
func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	if flowerbed == nil || len(flowerbed) == 0 {
		return false
	}

	for i := 0; i < len(flowerbed); {
		if flowerbed[i] == 1 {
			// 此处位置已经有花，则跳到下一个可以种花的地方（i+2）
			i += 2
		} else if i == len(flowerbed)-1 || flowerbed[i+1] == 0 {
			count++
			i += 2
		} else {
			i += 3
		}
	}

	return count >= n
}
func TestCanPlaceFlowers(t *testing.T) {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
}

//--------------------------------------------------------------------------------------------------------------------------------------

//--------------------------------------------------------------------------------------------------------------------------------------
// 345. 反转字符串中的元音字母
// 有点类似快速排序
func reverseVowels(s string) string {
	if len(s) < 2 {
		return s
	}

	s1 := []byte(s)

	l, r := 0, len(s)-1

	for l < r {
		// 从右至左，找第一个元音字母
		for l < r && !isLetter(s1[r]) {
			r--
		}
		// 从左至右，找第一个元音字母

		for l < r && !isLetter(s1[l]) {
			l++
		}
		s1[l], s1[r] = s1[r], s1[l]
		l++
		r--
	}

	return string(s1)
}
func isLetter(letter uint8) bool {
	return letter == 'A' || letter == 'a' || letter == 'E' || letter == 'e' || letter == 'I' || letter == 'i' || letter == 'O' || letter == 'o' ||
		letter == 'U' || letter == 'u'
}
func TestReverseVowels(t *testing.T) {
	fmt.Println(reverseVowels("leetcode"))
}

//--------------------------------------------------------------------------------------------------------------------------------------

//--------------------------------------------------------------------------------------------------------------------------------------
// 151. 反转字符串中的单词
func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	if len(s) == 1 {
		return s
	}

	split := strings.Split(s, " ")
	l, r := 0, len(split)-1
	for l < r {
		split[l], split[r] = split[r], split[l]
		l++
		r--
	}
	ans := ""
	for _, t := range split {
		if t == "" {
			continue
		}
		ans += t + " "
	}
	return ans[:len(ans)-1]
}
func TestReverseWords(t *testing.T) {
	fmt.Println(reverseWords("a good   example"))
}

//--------------------------------------------------------------------------------------------------------------------------------------

//--------------------------------------------------------------------------------------------------------------------------------------
// 238. 除自身以外数组的乘积
func productExceptSelf(nums []int) []int {
	// 维护两个（left、right）分别代表前 i 的乘积的乘积和数组
	// 举例：假设有数组 []int{1, 2, 3}
	//               left:{1, 1, 2}
	//                idx  0, 1, 2
	// 下标为 1 表示除开下标为 1 的元素外，下标 1 左边所以元素的乘积和

	if nums == nil || len(nums) < 2 {
		return []int{0}
	}

	left, right := make([]int, len(nums)), make([]int, len(nums))
	left[0] = 1
	right[len(nums)-1] = 1

	for i := 1; i < len(nums); i++ {
		left[i] = nums[i-1] * left[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	ans := make([]int, len(nums))

	for i, _ := range ans {
		ans[i] = left[i] * right[i]
	}
	return ans
}

//--------------------------------------------------------------------------------------------------------------------------------------

//--------------------------------------------------------------------------------------------------------------------------------------
// 334. 递增的三元子序列
// 这道题和上一题：「238. 除自身以外数组的乘积」有点类似
// 假设有数组：nums = []int{1, 5, 0, 4, 1, 3}
// 会设置两个辅助数组 left_min 和 right_max ，分别表示 在此位置(包括当前位置)左边的最小，和在此(包括当前位置)右边的最大值
// 比如根据 nums 建立 left_min 和 right_max
// 		  left_min: []int{1, 1, 0, 0, 0, 0}
//       right_max: []int{5, 5, 4, 4, 3, 3}
// 如果 nums[i] 大于了在此左边的最小值且小于了在此右边的最大值，就满足要求，返回 true
func increasingTriplet(nums []int) bool {
	if nums == nil || len(nums) < 3 {
		return false
	}

	left_min, right_max := make([]int, len(nums)), make([]int, len(nums))

	left_min[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		left_min[i] = min(nums[i], left_min[i-1])
	}

	right_max[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		right_max[i] = max(nums[i], right_max[i+1])
	}

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > left_min[i-1] && nums[i] < right_max[i+1] {
			return true
		}
	}

	return false
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func TestIncreasingTriplet(t *testing.T) {
	fmt.Println(increasingTriplet([]int{20, 100, 10, 12, 5, 13}))
}

//--------------------------------------------------------------------------------------------------------------------------------------

//--------------------------------------------------------------------------------------------------------------------------------------
// 443. 压缩字符串
func compress(chars []byte) int {
	slow, fast, res := 0, 0, 0

	for fast < len(chars) {
		fast++

		if fast != len(chars) && chars[slow] == chars[fast] {
			continue
		}

		chars[res] = chars[slow]
		res++
		num := fast - slow

		if num > 1 {
			itoa := strconv.Itoa(num)
			for i := 0; i < len(itoa); i++ {
				chars[res] = itoa[i]
				res++
			}
		}
		slow = fast
	}
	return res
}
func TestCompress(t *testing.T) {
	fmt.Println(compress([]byte{'a', 'b', 'b', 'b', 'b'}))
}

//--------------------------------------------------------------------------------------------------------------------------------------

//--------------------------------------------------------------------------------------------------------------------------------------
// 283. 移动零
func moveZeroes(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}

	slow, fast := 0, 0

	for fast < len(nums) {
		fast++
		if fast == len(nums) {
			break
		}
		if nums[slow] != 0 {
			slow = fast
			continue
		}
		if nums[fast] == 0 {
			continue
		}
		nums[slow], nums[fast] = nums[fast], nums[slow]
		slow++
		fast = slow
	}
}
func TestMoveZeroes(t *testing.T) {
	arr := []int{0, 1, 1, 0}
	moveZeroes(arr)
	fmt.Println(arr)
}

//--------------------------------------------------------------------------------------------------------------------------------------

//--------------------------------------------------------------------------------------------------------------------------------------
// 392. 判断子序列
func isSubsequence1(s string, t string) bool {
	s_arr := make([]uint8, 26)
	t_arr := make([]uint8, 26)

	for _, temp := range s {
		s_arr[temp-'a'] += 1
	}
	for _, temp := range t {
		t_arr[temp-'a'] += 1
	}

	for i, _ := range s_arr {
		if t_arr[i] < s_arr[i] {
			return false
		}
	}
	return true
}
func isSubsequence2(s string, t string) bool {
	p1, p2 := 0, 0

	for p1 < len(s) && p2 < len(t) {
		if s[p1] == t[p2] {
			p1++
			p2++
			continue
		}

		p2++
	}
	return p1 == len(s)
}
func TestIsSubsequence(t *testing.T) {
	fmt.Println(isSubsequence2("axc", "ahbgdc"))
}

//--------------------------------------------------------------------------------------------------------------------------------------

//--------------------------------------------------------------------------------------------------------------------------------------
// 11. 盛最多水的容器
// 使用双指针
// 不断收缩，来维护最大值
func maxArea(height []int) int {
	ans := -1
	l, r := 0, len(height)-1

	for l < r {
		ans = max(ans, (r-l)*min(height[l], height[r]))

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

//--------------------------------------------------------------------------------------------------------------------------------------

//--------------------------------------------------------------------------------------------------------------------------------------
// 1679. K 和数对的最大数目
func maxOperations(nums []int, k int) int {
	if nums == nil || len(nums) < 2 {
		return 0
	}
	ans := 0

	dic := make(map[int]int) // 记录数字的个数
	for _, num := range nums {
		dic[num] += 1
	}

	for _, num := range nums {
		gap := k - num
		if gap == num {
			if count, _ := dic[gap]; count >= 2 {
				ans++
				dic[gap] -= 2
			}
		} else {
			if count, found := dic[gap]; found && count > 0 {
				if count2, found2 := dic[num]; found2 && count2 > 0 {
					ans++
					dic[gap] -= 1
					dic[num] -= 1
				}
			}
		}

	}
	return ans
}
func TestMaxOperations(t *testing.T) {
	fmt.Println(maxOperations([]int{4, 4, 1, 3, 1, 3, 2, 2, 5, 5, 1, 5, 2, 1, 2, 3, 5, 4}, 2))
}

//--------------------------------------------------------------------------------------------------------------------------------------

//--------------------------------------------------------------------------------------------------------------------------------------
// 643. 子数组最大平均数 I
func findMaxAverage(nums []int, k int) float64 {
	if nums == nil || len(nums) < 2 {
		return float64(nums[0])
	}
	max_value := math.MinInt64
	sum := 0

	for _, num := range nums[:k] {
		sum += num
	}
	max_value = sum

	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		max_value = max(max_value, sum)
	}

	return float64(max_value) / float64(k)
}
func TestFindMaxAverage(t *testing.T) {
	fmt.Println(findMaxAverage([]int{0, 4, 0, 3, 2}, 1))
}

//--------------------------------------------------------------------------------------------------------------------------------------

//--------------------------------------------------------------------------------------------------------------------------------------
// 1456. 定长子串中元音的最大数目
func maxVowels(s string, k int) int {
	ans := 0
	sum := 0
	for _, temp := range s[:k] {
		if isMatch(uint8(temp)) {
			sum++
		}
	}
	ans = sum

	for i := k; i < len(s); i++ {
		if isMatch(s[i-k]) {
			sum--
		}
		if isMatch(s[i]) {
			sum++
		}
		ans = max(ans, sum)
	}
	return ans
}
func isMatch(n uint8) bool {
	return n == 'A' || n == 'a' || n == 'E' || n == 'e' || n == 'I' || n == 'i' || n == 'O' || n == 'o' || n == 'U' || n == 'u'
}

//--------------------------------------------------------------------------------------------------------------------------------------

//--------------------------------------------------------------------------------------------------------------------------------------
// 1004. 最大连续1的个数 III
// 滑动窗口的区间最多可以包含 多个 1 和 k 个 0
func longestOnes(nums []int, k int) int {
	max_value := 0
	zeros := 0

	left, right := 0, 0

	for right < len(nums) {
		if nums[right] == 0 {
			zeros++
		}

		if zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}

		right++

		value := right - left
		max_value = max(max_value, value)
	}
	return max_value
}
