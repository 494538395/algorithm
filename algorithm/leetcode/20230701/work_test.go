package _0230701

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

/*
	2023.7.1
	七月了。
	一年前的这个时候，应该是在办理 ONES 的入职相关材料吧。
	路需要靠自己走出来，自己和别人的差距很是很大的。
*/

//------------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 41. 数据流中的中位数

// 堆
// 大顶堆存放较小的一部分数字，从左至右，是递减的
// 小顶堆存放较大堆一部分数字，从左至右，是递增的
// 小顶堆最小的元素都比大顶堆最大堆元素大
// 举例。假设有数字集合 []int{1, 2, 3, 4, 5, 6}
// big_heap   大顶堆：[]int{3, 2, 1}
// small_heap 小顶堆：[]int{4, 5, 6}
// 大小顶堆的长度一样，则中位数是 ( big_heap[0] + small_heap[0] ) / 2
// --------
// 如果大小堆长度不一样，如：
// big_heap   大顶堆：[]int{3, 2, 1}
// small_heap 小顶堆：[]int{4, 5, 6}
type MedianFinder struct {
	count      int
	small_heap smallHeap
	big_heap   bigHeap
}

// 定义小顶堆
// 小顶堆是递增的，所以无需重新定义 Less() 排序规则（默认是升序）
type smallHeap struct {
	sort.IntSlice
}

func (s *smallHeap) Push(v interface{}) {
	s.IntSlice = append(s.IntSlice, v.(int))
}
func (s *smallHeap) Pop() interface{} {
	a := s.IntSlice
	last := a[len(a)-1]
	s.IntSlice = a[:len(a)-1]
	return last
}
func (s *smallHeap) Peek() int {
	return s.IntSlice[0]
}

// 定义大顶堆
// 因为大顶堆是递减的，所以需要重新定义 Less()
type bigHeap struct {
	sort.IntSlice
}

func (b *bigHeap) Push(v interface{}) {
	b.IntSlice = append(b.IntSlice, v.(int))
}
func (b *bigHeap) Pop() interface{} {
	a := b.IntSlice
	last := a[len(a)-1]
	b.IntSlice = a[:len(a)-1]
	return last
}
func (b *bigHeap) Less(i, j int) bool {
	return b.IntSlice[i] > b.IntSlice[j] // 递减 / 降序
}
func (b *bigHeap) Peek() int {
	return b.IntSlice[0]
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	samllHeap, bigHeap := &smallHeap{}, &bigHeap{}
	heap.Init(samllHeap)
	heap.Init(bigHeap)

	return MedianFinder{
		small_heap: *samllHeap,
		big_heap:   *bigHeap,
	}

}
func (this *MedianFinder) AddNum(num int) {
	if len(this.big_heap.IntSlice) != len(this.small_heap.IntSlice) {
		// 数量不一样，就往小顶堆插入，最后弹入大顶堆
		// 数量不一样的情况？
		// 只会存在大顶堆的数量少于小顶堆
		heap.Push(&this.small_heap, num)
		heap.Push(&this.big_heap, heap.Pop(&this.small_heap))
	} else {
		// 往大顶堆中插入
		heap.Push(&this.big_heap, num)
		// 弹出大顶堆最大值到小顶堆
		heap.Push(&this.small_heap, heap.Pop(&this.big_heap))
	}
}
func (this *MedianFinder) FindMedian() float64 {
	// 大小堆数量一样时，返回两个堆顶（最左边）的和再除以二
	if len(this.big_heap.IntSlice) == len(this.small_heap.IntSlice) {
		return float64(this.small_heap.Peek()+this.big_heap.Peek()) * 0.5
	}
	// 数量不一样（小顶堆多一个）
	// 直接返回小顶堆的堆顶元素（最左边）
	return float64(this.small_heap.Peek())
}
func quick(arr []int, left, right int) []int {
	if left >= right {
		return arr
	}

	pivot := arr[left]

	l, r := left, right

	for l < r {
		// 找到 r 指针的值小于 pivot 的位置
		for l < r && arr[right] >= pivot {
			r--
		}
		// 找到 l 指针的值大于 pivot 的位置
		for l < r && arr[left] <= pivot {
			l++
		}
		arr[l], arr[r] = arr[r], arr[l]
	}

	arr[left], arr[l] = arr[l], arr[left]

	quick(arr, left, l-1)
	quick(arr, l+1, right)
	return arr
}
func TestMedianFinder(t *testing.T) {
	finder := Constructor()
	finder.AddNum(-1)
	fmt.Println(finder.FindMedian())

	finder.AddNum(2)
	fmt.Println(finder.FindMedian())
	finder.AddNum(3)
	finder.AddNum(4)
	finder.AddNum(5)

	fmt.Println(finder.FindMedian())

}

//------------------------------------------------------------------------------------------------------------------------------------

//------------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 15. 二进制中1的个数
func hammingWeight(num uint32) int {
	ans := 0
	var d uint32
	d = 00000000000000000000000000000001

	for num > 0 {
		res := num & d
		if res != 0 {
			ans++
		}
		num = num >> 1
	}

	return ans
}
func TestNum(t *testing.T) {
	ans := 0

	var d uint32
	d = 00000000000000000000000000000001

	var a uint32
	a = 00000000000000000000000000001111

	for a > 0 {
		res := a & d
		if res != 0 {
			ans++
		}
		a = a >> 1
	}

	fmt.Println(ans)

}

//------------------------------------------------------------------------------------------------------------------------------------

//------------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 65. 不用加减乘除做加法
func add(a int, b int) int {
	for b != 0 {
		// 通过 与来获取进位
		carry := uint(a&b) << 1
		a ^= b
		b = int(carry)
	}
	return a
}

//------------------------------------------------------------------------------------------------------------------------------------

//------------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 56 - I. 数组中数字出现的次数
// 这题的问题是：数组中有 「两个」数字只出现了一次，出自之外的其他数字都出现了两次。
// 让找出这两个只出现一次的数字
// 使用异或
func singleNumbers(nums []int) []int {
	// 假定数字中有若干数字，其中只有 num1 和 num2 只出现了一次，其余均出现两次
	// 求整个数字的异或和
	// "整个数字的异或和本质上等与那两个只出现一次的异或和"，也就是  num1 ^ num2
	temp := 0
	for _, num := range nums {
		temp ^= num
	}

	// 根据 num1 ^ num2 异或结果从右至左寻找 num1 和 num2 第一个不一样的「特征」，根据这个「特征」拆分数组
	mask := 1
	for (temp & mask) == 0 {
		mask <<= 1
	}
	// mask 就是 num1 和 num2 的不同点
	// 比如 num1 是 0011 , num2 是 0001，那两者从右至左第一处不一样的「特征」就是 0010

	x, y := 0, 0

	for _, num := range nums {
		if num&mask == 0 {
			x ^= num
		} else {
			y ^= num
		}
	}
	return []int{x, y}

}

//------------------------------------------------------------------------------------------------------------------------------------

//------------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 56 - II. 数组中数字出现的次数 II
// 这道题和上个题的区别是：现在数组中只有「一个」数字出现了一次，其余数字均出现了三次
// 找出这个只出现「一次」的数字
func singleNumber(nums []int) int {
	ans := 0

	for i := 0; i < 64; i++ {
		bit := 0
		for _, num := range nums {
			bit += (num >> i) & 1
		}
		ans += (bit % 3) << i
	}

	return ans
}

//------------------------------------------------------------------------------------------------------------------------------------

//------------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 39. 数组中出现次数超过一半的数字
func majorityElement(nums []int) int {
	nums = quickSort(nums, 0, len(nums)-1)
	return nums[len(nums)>>1]
}
func quickSort(arr []int, left, right int) []int {
	if left >= right {
		return arr
	}
	pivot := arr[left]
	l, r := left, right
	for l < r {
		// 找右指针第一个小于 pivot 的数字
		for l < r && arr[r] >= pivot {
			r--
		}
		// 找左指针第一个大于 pivot 的数字
		for l < r && arr[l] <= pivot {
			l++
		}
		arr[l], arr[r] = arr[r], arr[l]
	}
	arr[left], arr[l] = arr[l], arr[left]
	quickSort(arr, left, l-1)
	quickSort(arr, l+1, right)
	return arr
}

//------------------------------------------------------------------------------------------------------------------------------------

//------------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 66. 构建乘积数组
// 方式一：暴力计算，超时
func constructArr1(a []int) []int {
	ans := make([]int, len(a))

	for i, _ := range ans {
		ans[i] = compute(a, i)
	}
	return ans
}
func compute(arr []int, idx int) int {
	ans := 1

	for i := 0; i < len(arr); i++ {
		if i == idx {
			continue
		}
		ans *= arr[i]
	}
	return ans
}
func constructArr(a []int) []int {
	// 维护 L 和 R 两个数字。分别表示对于的前缀乘积
	// 举例：有数组：[]int{1, ,2, 3}
	// L[0]=1 L[1]=1 L[2]=2
	// R[2]=1 R[1]=3 R[0]=6
	// 则 ans[0]=R[0]* L[0]  ans[1]=L[1]*R[1]

	l, r := make([]int, len(a)), make([]int, len(a))
	l[0] = 1
	r[len(a)-1] = 1
	// 初始化左右前缀乘积数组
	for i := 1; i < len(a); i++ {
		l[i] = a[i-1] * l[i-1]
	}

	for i := len(a) - 2; i >= 0; i-- {
		r[i] = r[i+1] * a[i+1]
	}

	ans := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		ans[i] = l[i] * r[i]
	}
	return ans
}
func TestConstructArr(t *testing.T) {
	fmt.Println(constructArr([]int{1, 2, 3}))
}

//------------------------------------------------------------------------------------------------------------------------------------

//------------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 14- I. 剪绳子
// 多3
func cuttingRope(n int) int {
	// 尽量多3
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	if n == 3 {
		// 3 只能切成 1、2
		return 2
	}

	res := n / 3
	mod := n % 3
	ans := 0

	if mod == 0 {
		// mod 为 0，表明 n 是 3 的整数倍，则全切成3
		ans = helpPow(3, res)
	} else if mod == 1 {
		// mod 为1 ，说明可以切成 res 个 3 和一个 1
		// 但是 1 乘积是没有任何意义的（乘 1 结果不会变大），所以会改成：切成 res-1 个 3 和 一个 4
		// 举例：7 。切成：2 个 3 和一个 1，结果是 6。而改成切成一个 3 和一个 4 ，结果是 12
		ans = helpPow(3, res-1) * 4
	} else {
		// mod 为 2
		// 就切成 res 个 3 和 一个 2
		ans = helpPow(3, res) * 2
	}
	return ans
}
func helpPow(a, mod int) int {
	sum := 1
	for i := mod; i > 0; i-- {
		sum *= a
	}
	return sum
}
func TestCutting(t *testing.T) {
	fmt.Println(cuttingRope(2))
}

//------------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 14- II. 剪绳子 II
// 和上面一题一样的，多切3
func cuttingRope2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 2 || n == 1 {
		return 1
	}
	if n == 3 {
		return 2
	}

	count := n / 3 // 3 的个数
	mod := n % 3
	ans := 0

	if mod == 0 {
		// 表明 n 是 3 的倍数，那全部切成3。如：9 可以切成 ：3、3、3
		ans = quickPow(count)
	} else if mod == 1 {
		// 表明会切成出 count 个 3 和 一个 1，但是乘 1 结果不会增加。所以去改成切成 count-1 个 3 和一个 4
		// 如 7：本来是 3、3、1 ,但是 3*3*1=9，不如切成：3、4，3*4=12
		ans = 4 * quickPow(count-1) % 1000000007
	} else {
		// 表明会切出 count 个 3 和一个 2，乘 2 是有意义的，所以保留
		ans = 2 * quickPow(count) % 1000000007
	}
	if ans == 1000000008 {
		return 1
	}
	return ans
}

// 对 3 求快速冥
func quickPow(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 3
	}
	if n%2 == 0 {
		half := quickPow(n / 2)
		return half * half % 1000000007
	} else {
		half := quickPow(n / 2)
		return 3 * half * half % 1000000007
	}
}
func TestCutting2(t *testing.T) {
	fmt.Println(cuttingRope2(127))
}

//------------------------------------------------------------------------------------------------------------------------------------

//------------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 57 - II. 和为s的连续正数序列
func findContinuousSequence(target int) [][]int {
	if target == 0 || target == 1 || target == 2 {
		return [][]int{}
	}
	ans := make([][]int, 0, target)
	temp := 0

	for slow := 1; slow <= target-1; slow++ {
		temp += slow

		for fast := slow + 1; fast <= target; fast++ {
			temp += fast
			if temp == target {
				t := []int{}
				for k := slow; k <= fast; k++ {
					t = append(t, k)
				}
				ans = append(ans, t)

				temp = 0
				break

			} else if temp > target {
				temp = 0
				break
			}
		}

	}
	return ans

}
func TestFindContinuousSequence(t *testing.T) {
	fmt.Println(findContinuousSequence(4))
}

//------------------------------------------------------------------------------------------------------------------------------------

//------------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 62. 圆圈中最后剩下的数字
// 约瑟夫环
func lastRemaining(n int, m int) int {
	if n == 0 {
		return 0
	}

	cur := -1
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	for len(arr) > 1 {
		cur += m
		if cur >= len(arr) {
			cur %= len(arr)
		}

		if cur == len(arr)-1 {
			arr = arr[0:cur]
		} else {
			arr = append(arr[0:cur], arr[cur+1:]...)
		}
		cur--
	}

	return arr[0]
}

func lastRemaining2(n int, m int) int {
	ans := 0

	for i := 1; i <= n; i++ {
		ans = (ans + m) % i
	}
	return ans
}

func TestLastRemaining(t *testing.T) {

	fmt.Println(lastRemaining(5, 3))
}
