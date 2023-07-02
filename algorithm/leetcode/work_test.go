package leetcode

import (
	"container/list"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	isNumber := func(s string) bool {
		isNum, isDot, ise_or_E := false, false, false
		s = strings.Trim(s, " ")

		for i, cur := range s {
			if cur >= '0' && cur <= '9' {
				// 当前是数字
				isNum = true
			} else if cur == '.' {
				// 当前是小数点
				if isDot || ise_or_E {
					// isDot： 之前是小数点，现在也是小数点，不合法  ..
					// ise_or_E： 之前是小数点，现在是 冥号，不合法 .e
					return false
				}
				isDot = true
			} else if cur == 'e' || cur == 'E' {
				// 当前是冥号
				if !isNum || ise_or_E {
					// !isNum 现在是冥号，如果前面不是数字，就不合法
					// ise_or_E 现在是冥号，前一个不能也是冥号，连着两个冥号，不合法  ee
					return false
				}
				ise_or_E = true
				isNum = false // 重置isNum，因为‘e’或'E'之后也必须接上整数，防止出现 123e或者123e+的非法情况
			} else if cur == '-' || cur == '+' {
				// 当前是加减号
				// 正负号只可能出现在第一个位置，或者出现在‘e’或'E'的后面一个位置
				if i != 0 && s[i-1] != 'e' && s[i-1] != 'E' {
					return false
				}

			} else {
				return false
			}
		}
		return isNum
	}

	fmt.Println("isNumber(+100)-->", isNumber("1.e"))

}

func Test2(t *testing.T) {
	str := "abcdef"

	for _, i := range str {
		fmt.Println("i-->", i)
	}
}

// 剑指 Offer 67. 把字符串转换成整数
func Test3(t *testing.T) {
	strToInt := func(str string) int {
		if len(str) == 0 {
			return 0
		}
		str = strings.Trim(str, " ")
		var flag int
		flag, firstNumIdx, lastNumIdx := 1, -1, -1
		invaildStr := false
		first := str[0]

		isNum := func(s uint8) bool {
			return s >= '0' && s <= '9'
		}
		isFlag := func(s uint8) bool {
			return s == '+' || s == '-'
		}

		if !isFlag(first) && !isNum(first) {
			// 第一位既不是 +、- 也不是数字，不合法，直接返回 0
			return 0
		}

		if isFlag(first) {
			if first == '+' {
				flag = 1
			} else {
				flag = -1
			}
			if len(str) == 1 {
				return 0
			}
			// 获取第二位
			second := str[1]
			if !isNum(second) {
				return 0
			}
			firstNumIdx = 1

		} else {
			// 第一位是数字
			firstNumIdx = 0
		}

		for i := 1; i < len(str); i++ {
			cur := str[i]
			if !isNum(cur) {
				lastNumIdx = i
				invaildStr = true
				break
			}
		}

		sub := ""
		if invaildStr {
			sub = str[firstNumIdx:lastNumIdx]
		} else {
			sub = str[firstNumIdx:]
		}
		if flag == -1 {
			sub = "-" + sub
		}

		parse, err := strconv.ParseInt(sub, 10, 32)
		if err != nil {

		}
		return int(parse)
	}
	toInt := strToInt("")

	fmt.Println("strToInt->", toInt)

}
func Test4(t *testing.T) {
	str := "-91283472332"
	parseInt, _ := strconv.ParseInt(str, 10, 32)
	fmt.Println(parseInt)
}
func strToInt(str string) int {
	str = strings.Trim(str, " ")
	if len(str) == 0 {
		return 0
	}

	if str[0] != '+' && str[0] != '-' && str[0] < '0' && str[0] > '9' {
		return 0
	}

	hasSign := false
	sign := 1
	num := 0

	if str[0] == '+' {
		hasSign = true
	}
	if str[0] == '-' {
		hasSign = true
		sign = -1
	}

	i := 0
	if hasSign {
		if len(str) == 1 || (str[1] < '0' && str[1] > '9') {
			return 0
		}
		i++
	}
	for ; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			num = num*10 + int(str[i]-'0')
			if sign*num > math.MaxInt32 {
				return math.MaxInt32
			} else if sign*num < math.MinInt32 {
				return math.MinInt32
			}
		} else {
			break
		}
	}

	return sign * num
}

// 从尾到头打印链表
// 考察点：递归、栈
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var f func(n *ListNode) []int
	f = func(n *ListNode) []int {
		if n == nil {
			return []int{}
		}
		return append(f(n.Next), n.Val)
	}
	return f(head)
}
func Test5(t *testing.T) {
	head := ListNode{
		Val: 1,
		//Next: &ListNode{Val: 2},
	}

	ints := reversePrint(&head)

	fmt.Println(ints)

}

// 反转链表
// 由于单向链表没有 prev ，所以需要我们手动存储 prev
func reverseList(head *ListNode) *ListNode {

	var prev *ListNode
	cur := head

	for {
		if head == nil {
			break
		}
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next

	}

	return head
}

// 98 验证二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helpVerify(root, math.MinInt64, math.MaxInt64)
}
func helpVerify(root *TreeNode, lower, higher int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= higher {
		return false
	}
	return helpVerify(root.Left, lower, root.Val) && helpVerify(root.Right, root.Val, higher)
}
func verify(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false
		}
		if root.Left.Left != nil {
			if root.Left.Left.Val >= root.Val {
				return false
			}
		}
		if root.Left.Right != nil {
			if root.Left.Right.Val >= root.Val {
				return false
			}
		}
	}
	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		}
		if root.Right.Left != nil {
			if root.Right.Left.Val <= root.Val {
				return false
			}
		}
		if root.Right.Right != nil {
			if root.Right.Right.Val <= root.Val {
				return false
			}
		}
	}
	return verify(root.Left) && verify(root.Right)
}
func TestIsValidBST(t *testing.T) {
	left := &TreeNode{
		Val: 4,
	}
	right := &TreeNode{
		Val:   6,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 7},
	}

	root := &TreeNode{Val: 5, Left: left, Right: right}
	//  	       5
	//   	   /       \
	//	      4         6
	// 	    	      /   \
	// 	    	     3      7

	res := isValidBST(root)
	fmt.Println(res)
}

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 35. 复杂链表的复制
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	dic := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		dic[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		dic[cur].Next = dic[cur.Next]
		dic[cur].Random = dic[cur.Random]

	}

	return dic[head]
}
func copyRandomList2(head *Node) *Node {
	if head == nil {
		return head
	}

	cur := head
	dum := &Node{}
	pre := dum

	for cur != nil {
		tempNext := &Node{Val: cur.Val}
		pre.Next = tempNext
		var tempRandom *Node
		if cur.Random != nil {
			tempRandom = &Node{Val: cur.Random.Val}
		}
		pre.Random = tempRandom

		pre = pre.Next
		cur = cur.Next
	}
	return dum.Next
}

type NodeNormal struct {
	Val  int
	Next *NodeNormal
}

// 复制普通链表
func copyNormalList(head *NodeNormal) *NodeNormal {
	dum := &NodeNormal{}
	cur := head

	pre := dum

	for cur != nil {
		temp := &NodeNormal{Val: cur.Val}
		pre.Next = temp
		cur = cur.Next
		pre = temp
	}
	return dum.Next
}
func TestCopyList(t *testing.T) {
	head := &NodeNormal{Val: 1, Next: &NodeNormal{Val: 2, Next: &NodeNormal{Val: 3}}}
	res := copyNormalList(head)
	fmt.Printf("head-->%p\n", head)
	fmt.Printf("res-->%p\n", res)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}

}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指Offer 18.删除链表的节点
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return head.Next
	}
	temp := &ListNode{}
	temp.Next = head
	cur := head
	pre := &ListNode{}

	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			cur.Next = nil
		}
		pre = cur
		cur = cur.Next

	}

	return temp.Next
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指Offer 22 链表中倒数第K个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k < 0 {
		return head
	}
	cur := head
	// 先遍历一遍得到链表的长度

	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}
	count := length - k

	// 遍历获取结果
	cur = head
	for cur != nil {
		if count == 0 {
			return cur
		}
		count--
		cur = cur.Next
	}
	return head
}

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指Offer 25。合并两个排序的链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummy := &ListNode{}
	cur := dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next

		} else {
			cur.Next = l2
			l2 = l2.Next

		}
		cur = cur.Next

	}

	if l1 == nil {
		cur.Next = l2
	}
	if l2 == nil {
		cur.Next = l1
	}
	return dummy.Next
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指Offer 52. 两个链表的第一个公共节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	if headA == headB {
		return headA
	}

	vis := make(map[*ListNode]bool)

	for temp := headA; temp != nil; temp = temp.Next {
		vis[temp] = true
	}

	for temp := headB; temp != nil; temp = temp.Next {
		if vis[temp] {
			return temp
		}
	}
	return nil
}
func TestGetIntersectionNode(t *testing.T) {
	intersect := &ListNode{Val: 3, Next: &ListNode{Val: 4}}
	headA := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: intersect}}
	headB := &ListNode{Val: 2, Next: intersect}

	res := getIntersectionNode(headA, headB)
	fmt.Println(res)
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指Offer 32 从上到下打印二叉树
// 层序遍历
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var ans []int
	queue := list.New()
	queue.PushBack(root) //把根节点入队列

	for queue.Len() != 0 {
		node := queue.Front().Value.(*TreeNode) // 拿出队列最左边的节点
		ans = append(ans, node.Val)             // 追加结果集
		queue.Remove(queue.Front())             // 删除队列最左边的节点
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return ans
}
func TestLevelOrder(t *testing.T) {
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	res := levelOrder(root)
	fmt.Println(res)

}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 21.调整数组顺序使奇数在偶数前面
func exchange(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	var oddNumbers []int
	var evenNumbers []int

	for _, num := range nums {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
			continue
		}
		oddNumbers = append(oddNumbers, num)
	}
	return append(oddNumbers, evenNumbers...)
}
func TestExchange(t *testing.T) {
	fmt.Println(exchange([]int{1, 2, 3, 4}))
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 57 和为 s 的两个数字
func twoSum(nums []int, target int) []int {
	dic := make(map[int]struct{})
	dic[nums[0]] = struct{}{}

	for i := 1; i < len(nums); i++ {
		cur := target - nums[i]
		if _, found := dic[cur]; found {
			return []int{cur, nums[i]}
		}
		dic[nums[i]] = struct{}{}

	}
	return []int{}
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 58-1 翻转单词顺序
func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	if s == "" || len(s) == 1 {
		return s
	}
	ans := ""

	fast, slow := len(s)-1, len(s)-1

	for fast >= 0 && slow >= 0 {
		cur := s[slow]
		fast = slow
		if cur != ' ' {
			fast--
			for fast >= 0 && s[fast] != ' ' {
				fast--
			}
			ans += s[fast+1 : slow+1]
			ans += " "
		}
		slow = fast - 1

	}

	return ans[:len(ans)-1]
}
func TestReverseWords(t *testing.T) {
	fmt.Println(reverseWords("l am a"))
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 05 替换空格
func replaceSpace(s string) string {
	ans := ""
	for i := 0; i < len(s); i++ {
		cur := s[i]
		if cur == ' ' {
			ans += "%20"
			continue
		}
		ans += string(s[i])
	}
	return ans
}
func TestReplaceSpace(t *testing.T) {
	fmt.Println(replaceSpace(" hello world!"))
}

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 58 - II. 左旋转字符串
func reverseLeftWords(s string, n int) string {
	if n == len(s) {
		return s
	}
	return s[n:] + s[0:n]
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 20. 表示数值的字符串
func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	if s == "" {
		return false
	}

	isNumberFunc := func(s uint8) bool {
		if s >= '0' && s <= '9' {
			return true
		}
		return false
	}
	isLabelFunc := func(s uint8) bool {
		if s == '+' || s == '-' {
			return true
		}
		return false
	}
	isEFunc := func(s uint8) bool {
		if s == 'e' || s == 'E' {
			return true
		}
		return false
	}
	isPointFunc := func(s uint8) bool {
		return s == '.'
	}

	isNum := false //表示上一是否为数字
	isPoint := false
	isE := false

	for i := 0; i < len(s); i++ {
		// 如果当前位是符号，符号只能出现在第一个位置和E后的第一个位置
		if isLabelFunc(s[i]) {
			if i != 0 && !isEFunc(s[i-1]) {
				return false
			}
		} else if isNumberFunc(s[i]) {
			isNum = true
		} else if isPointFunc(s[i]) {
			// 点前可以没有数字，但是点前不能是点，也不能是E    ..或E.或e. 是不满足的
			if isPoint || isE {
				return false
			}
			isPoint = true
		} else if isEFunc(s[i]) {
			// 当前位是 E ，前一位必须是数字，前一位一定不能是 E
			if !isNum || isE {
				return false
			}
			isE = true
			isNum = false
		} else {
			return false
		}

	}
	return isNum // 一定要以数字结尾
}
func TestIsNumber(t *testing.T) {
	fmt.Println(isNumber("11e"))
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 67. 把字符串转换成整数
func strToInt2(str string) int {
	str = strings.Trim(str, " ")
	isNegative, skipFirst := false, false
	isNumber := func(n uint8) bool {
		return n >= '0' && n <= '9'
	}

	if str[0] == '-' {
		isNegative = true
		skipFirst = true
	}
	if str[0] == '+' {
		skipFirst = true
	}
	start, end := 0, 0
	if skipFirst {
		start = 1
		end = 1
	}

	for i := start; i < len(str); i++ {
		cur := str[i]
		if isNumber(cur) {
			end++
		} else {
			break
		}
	}
	if start == end {
		return 0
	}

	numberStr := str[start:end]
	parseInt, err := strconv.ParseInt(numberStr, 10, 32)
	if err != nil {
		if isNegative {
			return math.MinInt32
		}
		return math.MaxInt32
	}
	if isNegative {
		return int(parseInt) * -1
	}

	return int(parseInt)
}
func TestStrToInt2(t *testing.T) {
	fmt.Println(strToInt2("words"))
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
//	剑指 Offer 35. 复杂链表的复制
func copyRandomList3(head *Node) *Node {
	dic := make(map[*Node]*Node)

	cur := head

	for cur != nil {
		temp := &Node{Val: cur.Val}
		dic[cur] = temp
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		dic[cur].Next = dic[cur.Next]
		dic[cur].Random = dic[cur.Random]
		cur = cur.Next
	}

	return dic[head]
}
func TestCopyRandomList3(t *testing.T) {
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}

	head := &Node{Val: 2, Next: node2, Random: nil}
	node2.Next = node3
	node2.Random = head
	node3.Random = node2

	fmt.Println(copyRandomList3(head))

}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 09. 用两个栈实现队列
type CQueue struct {
	inStack, outStack []int
}

func Constructor2() CQueue {
	return CQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}
func (this *CQueue) AppendTail(value int) {
	this.inStack = append(this.inStack, value)
}
func (this *CQueue) DeleteHead() int {
	if len(this.outStack) == 0 {
		if len(this.inStack) == 0 {
			return -1
		}

	}
	first := this.outStack[0]
	this.outStack = this.outStack[1:]

	return first
}
func (this *CQueue) inToOut() {
	if len(this.inStack) == 0 {
		return
	}
	this.outStack = append(this.outStack, this.inStack[0])
	this.inStack = this.inStack[1:]
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 30. 包含min函数的栈
type MinStack struct {
	stack []int
	min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min:   []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	top := this.min[len(this.min)-1]
	this.min = append(this.min, min(top, x))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]

}

func (this *MinStack) Min() int {
	return this.min[len(this.min)-1]
}

func TestMinStack(t *testing.T) {
	stack := Constructor()
	stack.Push(0)
	stack.Push(1)
	stack.Push(0)

	fmt.Println(stack.Min())
	stack.Pop()
	//fmt.Println(stack.Top())
	fmt.Println(stack.Min())

}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 59 - I. 滑动窗口的最大值
func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil || len(nums) < 2 || k == 1 {
		return nums
	}

	left, right := 0, k-1
	ans := make([]int, 0, len(nums))
	var temp []int

	for right < len(nums) {
		if right == len(nums) {
			temp = append(temp, nums[left:right]...)
		} else {
			temp = append(temp, nums[left:right+1]...)
		}

		sort.Ints(temp)
		ans = append(ans, temp[len(temp)-1])
		left++
		right++
		temp = []int{}
	}
	return ans
}
func maxSlidingWindow2(nums []int, k int) []int {
	if nums == nil || len(nums) < 2 || k == 1 {
		return nums
	}
	ans := make([]int, 0, len(nums))
	assist_stack := make([]int, 0, len(nums))

	temp := append([]int{}, nums[0:k]...)
	sort.Ints(temp)

	first_max := temp[len(temp)-1]
	temp = nil
	ans = append(ans, first_max)
	assist_stack = append(assist_stack, first_max)

	for i := k - 1; i < len(nums); i++ {
		cur := nums[i]
		top := ans[len(ans)-1]
		if cur > top {
			ans = append(ans, cur)
		} else {
			ans = append(ans, top)
		}
	}

	return ans
}
func maxSlidingWindow3(nums []int, k int) []int {
	if nums == nil || len(nums) < 2 {
		return nums
	}
	ans := make([]int, 0, len(nums)-k+1) // len(nums)-k+1 表示长度为 len(num) 的数组含有  len(nums)-k+1 个长度为 k 的窗口
	queue := make([]int, 0, len(nums))   // 双端队列，队列头部记录当前窗口的最大值

	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		// 因为 queue 最左侧维护的是窗口的最大值，会清除较小值
		for len(queue) > 0 && queue[len(queue)-1] < cur {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, cur)

		// 如果 1>=k 说明已经形成了最少一个区间且此时第一个区间已经处理完，现在处理的是第二个或者后面的区间，那就会判断当前队列维护的最大值是不是前一个区间的最大值
		if i >= k && queue[0] == nums[i-k] {
			queue = queue[1:]
		}

		// 当满足时，说明已经形成了区间，将队列最左侧的值放入结果集
		if i >= k-1 {
			ans = append(ans, queue[0])
		}

	}
	return ans
}
func TestMaxSlidingWindow(t *testing.T) {
	fmt.Println(maxSlidingWindow3([]int{9, 11}, 2))
}
func TestArr(t *testing.T) {
	arr := []int{1, 2, 3}
	fmt.Println(arr[2:3])
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 59 - II. 队列的最大值
// 辅助栈
type MaxQueue struct {
	dataQueue []int
	MaxQueue  []int //队列最左侧维护了当前 dataQueue里的最大值
}

func Constructor3() MaxQueue {
	return MaxQueue{
		dataQueue: []int{},
		MaxQueue:  []int{},
	}
}
func (this *MaxQueue) Max_value() int {
	if len(this.MaxQueue) == 0 {
		return -1
	}
	return this.MaxQueue[0]
}
func (this *MaxQueue) Push_back(value int) {
	this.dataQueue = append(this.dataQueue, value)
	// 如果新插入的值大于队列最右端的值，则清除队列中小于即将插入的值
	for len(this.MaxQueue) > 0 && this.MaxQueue[len(this.MaxQueue)-1] < value {
		this.MaxQueue = this.MaxQueue[:len(this.MaxQueue)-1]
	}
	this.MaxQueue = append(this.MaxQueue, value)
}
func (this *MaxQueue) Pop_front() int {
	if len(this.dataQueue) == 0 {
		return -1
	}
	front := this.dataQueue[0]
	this.dataQueue = this.dataQueue[1:]
	if front == this.MaxQueue[0] {
		this.MaxQueue = this.MaxQueue[1:]
	}
	return front
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 03. 数组中重复的数字
func findRepeatNumber(nums []int) int {
	sort.Ints(nums)
	if nums == nil || len(nums) < 2 {
		return -1
	}

	for left := 0; left < len(nums)-1; left++ {
		right := left + 1
		if nums[left] == nums[right] {
			return nums[left]
		}
	}
	return -1
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 53 - I. 在排序数组中查找数字 I
func search(nums []int, target int) int {
	dic := make(map[int]int, len(nums))

	for _, num := range nums {
		dic[num] = dic[num] + 1
	}
	count, found := dic[target]
	if found {
		return count
	}
	return -1
}

// 复习二分查找
func search2(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
func TestSearch2(t *testing.T) {
	fmt.Println(search2([]int{1, 2, 3, 4, 5}, 6))
}

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 53 - II. 0～n-1中缺失的数字
func missingNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}
	return len(nums)
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 04. 二维数组中的查找
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, rows := range matrix {
		i := search2(rows, target)
		if i >= 0 && i < len(rows) && rows[i] == target {
			return true
		}
	}
	return false
}
func TestFindNumberIn2DArray(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(findNumberIn2DArray(matrix, 30))

}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 11. 旋转数组的最小数字
// 排序方式
func minArray1(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	sort.Ints(numbers)
	return numbers[0]
}

// 二分方式
func minArray2(numbers []int) int {
	left, right := 0, len(numbers)-1

	for left < right {
		mid := left + (right-left)/2 // 防止溢出

		// numbers[mid] > numbers[right]，那最小值一定在 mid 右边
		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else if numbers[mid] < numbers[right] {
			right = mid
		}

	}

	return numbers[left]
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 50. 第一个只出现一次的字符
func firstUniqChar(s string) byte {
	dic := make(map[uint8]int, len(s))

	for i, _ := range s {
		cur := s[i]

		dic[cur] = dic[cur] + 1
	}

	for i, _ := range s {
		cur := s[i]
		if count, found := dic[cur]; found {
			if count == 1 {
				return cur
			}
		}
	}
	return ' '
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 29. 顺时针打印矩阵
func spiralOrder(matrix [][]int) []int {
	ans := []int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return ans
	}

	rows, columns := len(matrix), len(matrix[0])
	visited := make([][]bool, rows)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, columns)
	}

	total := rows * columns
	x, y := 0, 0

	var (
		directions = [][]int{
			{0, 1},  // 右
			{1, 0},  // 下
			{0, -1}, // 左
			{-1, 0}, // 上
		}
		directionIdx = 0
	)

	for i := 0; i < total; i++ {
		cur := matrix[x][y]
		visited[x][y] = true
		ans = append(ans, cur)

		next_x, next_y := x+directions[directionIdx][0], y+directions[directionIdx][1]
		if next_x < 0 || next_x >= rows || next_y < 0 || next_y >= columns || visited[next_x][next_y] {
			directionIdx = (directionIdx + 1) % 4 // 为什么是4？ 因为只有上下左右四个方式
		}
		x += directions[directionIdx][0]
		y += directions[directionIdx][1]
	}

	return ans
}
func TestDoubleArr(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(spiralOrder(matrix))
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 31. 栈的压入、弹出序列
func validateStackSequences(pushed []int, popped []int) bool {
	st := []int{}
	j := 0

	for _, num := range pushed {
		st = append(st, num)
		for len(st) > 0 && st[len(st)-1] == popped[j] {
			st = st[:len(st)-1] //模拟出栈
			j++
		}

	}
	return len(st) == 0 // 如果能全部出栈，说明正确
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 32 - II. 从上到下打印二叉树 II
func levelOrder2(root *TreeNode) [][]int {
	ans := [][]int{}
	queue := list.New()

	if root == nil {
		return ans
	}
	queue.PushBack(root)
	var temp []int

	for queue.Len() != 0 {
		size := queue.Len()

		for i := 0; i < size; i++ {
			front := queue.Front().Value.(*TreeNode)
			temp = append(temp, front.Val)
			queue.Remove(queue.Front())
			if front.Left != nil {
				queue.PushBack(front.Left)
			}
			if front.Right != nil {
				queue.PushBack(front.Right)
			}
		}
		ans = append(ans, temp)
		temp = []int{}
	}
	return ans
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 32 - III. 从上到下打印二叉树 III
func levelOrder3(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	queue := list.New()
	count := 1

	queue.PushBack(root)

	for queue.Len() > 0 {
		size := queue.Len()
		temp := make([]int, size)

		for i := 0; i < size; i++ {
			front := queue.Front().Value.(*TreeNode)
			queue.Remove(queue.Front())

			if count%2 == 0 {
				temp[size-1-i] = front.Val
			} else {
				temp[i] = front.Val
			}
			if front.Left != nil {
				queue.PushBack(front.Left)
			}
			if front.Right != nil {
				queue.PushBack(front.Right)
			}
		}
		count++

		ans = append(ans, temp)
	}
	return ans
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
//剑指 Offer 26. 树的子结构
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if B == nil {
		return false
	}
	res := false
	if A.Val == B.Val {
		res = helper(A, B)
	}
	if !res {
		res = isSubStructure(A.Left, B)
	}
	if !res {
		res = isSubStructure(A.Right, B)

	}
	return res
}
func helper(a, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return helper(a.Left, b.Left) && helper(a.Right, b.Right)
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 27. 二叉树的镜像
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	res := reverse(root)
	return res
}
func reverse(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	node.Left, node.Right = node.Right, node.Left
	reverse(node.Left)
	reverse(node.Right)
	return node
}
func TestMirrorTree(t *testing.T) {
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 7}
	node3 := &TreeNode{Val: 1}
	node4 := &TreeNode{Val: 3}
	node1.Left = node3
	node1.Right = node4

	root := &TreeNode{Val: 4}
	root.Left = node1
	root.Right = node2
	treeNode := mirrorTree(root)
	fmt.Println(treeNode.Left.Val)
	fmt.Println(treeNode.Right.Val)

	fmt.Println(midForeach(treeNode))
}
func midForeach(root *TreeNode) []int {
	ans := []int{}
	inorder(root, &ans)
	return ans
}

func inorder(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	*ans = append(*ans, root.Val)
	inorder(root.Left, ans)
	inorder(root.Right, ans)

}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 28. 对称的二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil || root.Right == nil {
		return false
	}

	return isSymmetricNode(root.Left, root.Right)
}

func isSymmetricNode(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}

	return isSymmetricNode(a.Left, b.Right) && isSymmetricNode(a.Right, b.Left)
}

func TestIsSymmetric(t *testing.T) {
	node1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}
	node2 := &TreeNode{Val: 2, Right: &TreeNode{Val: 1}}

	root := &TreeNode{Val: 1}
	root.Left = node1
	root.Right = node2

	fmt.Println(isSymmetric(root))
}

//----------------------------------------------------------------------------------------------------------------------------------
