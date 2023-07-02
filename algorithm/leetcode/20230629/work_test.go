package _0230629

import (
	"container/list"
	"fmt"
	"math"
	"sort"
	"sync"
	"testing"
	"time"
)

/*
	2023.06.29
	来深圳整整一年了。
	回想起一年前这个时候的心态，充满了希望、斗志、和忐忑。此刻的我也有斗志，但也为接下来的不确定而担心。
	我不愿去到一个没法发胀的地方，我想要走的更远。
	所以，我现在能做的只有拼命寻找机会。
*/

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 12. 矩阵中的路径
var (
	direction = [][]int{
		{0, 1},  // 右
		{1, 0},  // 下
		{0, -1}, // 左
		{-1, 0}, // 上
	}
	rows, columns int
)

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	rows, columns = len(board), len(board[0])

	first_letter := word[0]

	visited := make([][]bool, rows)
	for i, _ := range visited {
		visited[i] = make([]bool, columns)
	}

	p := 0

	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			cur := board[x][y]
			if cur == first_letter {
				if visited[x][y] {
					continue
				}

				ans := findAndSearch(x, y, board, visited, p, word)
				if ans {
					return true
				}
			}

		}
	}
	return false
}
func findAndSearch(x, y int, board [][]byte, visited [][]bool, p int, word string) bool {
	if x < 0 || x >= rows || y < 0 || y >= columns || visited[x][y] || board[x][y] != word[p] {
		return false
	}
	visited[x][y] = true
	if p == len(word)-1 {
		return true
	}

	// 右边
	ans := findAndSearch(x+direction[0][0], y+direction[0][1], board, visited, p+1, word) ||
		// 下边
		findAndSearch(x+direction[1][0], y+direction[1][1], board, visited, p+1, word) ||
		// 左边
		findAndSearch(x+direction[2][0], y+direction[2][1], board, visited, p+1, word) ||
		// 上边
		findAndSearch(x+direction[3][0], y+direction[3][1], board, visited, p+1, word)

	if ans {
		return true
	}
	visited[x][y] = false
	return false

}
func TestExist(t *testing.T) {
	board := [][]byte{
		{'c', 'a', 'a'},
		{'a', 'a', 'a'},
		{'b', 'c', 'd'},
	}

	fmt.Println(exist(board, "aab"))
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 13. 机器人的运动范围
func movingCount(m int, n int, k int) int {
	if m == 0 || n == 0 {
		return 0
	}
	visited := make([][]bool, m)
	for i, _ := range visited {
		visited[i] = make([]bool, n)
	}
	count := 0
	dfs(0, 0, visited, &count, k, m, n)
	return count
}
func helper(x, y, k int) bool {
	// 获取十位和个位
	decade_x := x / 10
	units_x := x % 10

	decade_y := y / 10
	units_y := y % 10

	return decade_x+units_x+decade_y+units_y <= k

}
func dfs(x, y int, visited [][]bool, count *int, k int, rows, columns int) {
	// 合法性判断
	if x < 0 || x >= rows || y < 0 || y >= columns || visited[x][y] || !helper(x, y, k) {
		return
	}
	*count++
	visited[x][y] = true

	direction := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	dfs(x+direction[0][0], y+direction[0][1], visited, count, k, rows, columns)
	dfs(x+direction[1][0], y+direction[1][1], visited, count, k, rows, columns)
	dfs(x+direction[2][0], y+direction[2][1], visited, count, k, rows, columns)
	dfs(x+direction[3][0], y+direction[3][1], visited, count, k, rows, columns)

	return
}
func TestMovingCount(t *testing.T) {
	fmt.Println(movingCount(2, 2, 17))
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 34. 二叉树中和为某一值的路径
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == target {
			return [][]int{{root.Val}}
		}
	}
	ans := [][]int{}

	helper2(root, &ans, target, &[]int{})

	return ans
}
func helper2(node *TreeNode, ans *[][]int, target int, temp *[]int) {
	if node == nil {
		return
	}
	num := node.Val
	new_temp := append(make([]int, 0, len(*temp)+1), *temp...)
	new_temp = append(new_temp, num)

	target -= node.Val
	if target == 0 && node.Left == nil && node.Right == nil {
		*ans = append(*ans, new_temp)
	} else {
		helper2(node.Left, ans, target, &new_temp)
		helper2(node.Right, ans, target, &new_temp)
	}
}
func TestPathSum(t *testing.T) {
	node1 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   11,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 2}},
	}
	node2 := &TreeNode{
		Val:  8,
		Left: &TreeNode{Val: 13},
		Right: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 1}}}

	root := &TreeNode{Val: 5}
	root.Left = node1
	root.Right = node2

	//root := &TreeNode{Val: -2, Right: &TreeNode{Val: -3}}

	fmt.Println(pathSum(root, 22))

}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 54. 二叉搜索树的第k大节点
// 「 只要是平衡二叉树，就满足：前序遍历的结果是递增的 」
func kthLargest(root *TreeNode, k int) int {
	arr := []int{}
	preForeach(root, &arr)

	return arr[len(arr)-k]

}
func preForeach(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}
	preForeach(node.Left, arr)
	*arr = append(*arr, node.Val)
	preForeach(node.Right, arr)
}
func TestPreForeach(t *testing.T) {

	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}}
	arr := []int{}
	preForeach(root, &arr)
	fmt.Println(arr)

}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 55 - I. 二叉树的深度
// 层序遍历计算深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	depth := 0

	for queue.Len() > 0 {
		depth++
		temp := []*TreeNode{}
		size := queue.Len()
		for i := 0; i < size; i++ {
			front := queue.Front().Value.(*TreeNode)
			if front.Left != nil {
				temp = append(temp, front.Left)
			}
			if front.Right != nil {
				temp = append(temp, front.Right)
			}

			queue.Remove(queue.Front())
		}

		for _, node := range temp {
			queue.PushBack(node)
		}

	}
	return depth
}
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth2(root.Left), maxDepth2(root.Right)) + 1
}
func TestMaxDepth(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}

	fmt.Println(maxDepth(root))

}
func TestMaxDepth2(t *testing.T) {

	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2}}

	fmt.Println(maxDepth2(root))
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 55 - II. 平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return abs(depth(root.Left)-depth(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)

}
func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(depth(node.Left), depth(node.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 68 - II. 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if right == nil {
		return left
	}
	if left == nil {
		return right
	}

	return nil
}

func TestLowestCommonAncestor(t *testing.T) {

	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2}}
	fmt.Println(lowestCommonAncestor(root, &TreeNode{Val: 4}, &TreeNode{Val: 1}))

}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 10- I. 斐波那契数列
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	first, second := 0, 1

	for i := 2; i <= n; i++ {
		third := (first + second) % 1000000007
		first = second
		second = third
	}

	return int(second)

}
func TestFib(t *testing.T) {
	fmt.Println(fib(2))

}
func TestArr(t *testing.T) {
	a := make([]int, 0, 4)
	a[2] = 4
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 10- II. 青蛙跳台阶问题
// f(n)=f(n-1)+f(n-2)
// 这个式子怎么来的？
func numWays(n int) int {
	if n == 0 || n < 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}

	return int(dp[n])

}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 42. 连续子数组的最大和
func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	cur := nums[0]
	ans := cur

	for i := 1; i < len(nums); i++ {
		if cur < 0 {
			cur = nums[i]
		} else {
			cur += nums[i]
		}
		ans = max(ans, cur)
	}
	return ans
}

//----------------------------------------------------------------------------------------------------------------------------------
// 郑州超聚变面试题
func TestWork1(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	info := make(chan struct{})

	// oddNumber:
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i += 2 {
			info <- struct{}{}
			if i%2 != 0 {
				fmt.Println("oddNumber  -->", i)
			}
		}

	}()
	// evenNumber
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			<-info
			if i%2 == 0 {
				fmt.Println("evenNumber  -->", i)
			}
		}
	}()
	wg.Wait()
}
func TestWork2(t *testing.T) {
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	go func() {
		for i := 1; i <= 100; i += 2 {
			<-ch1
			fmt.Println("oddNumber -->", i)
			ch2 <- true
		}
	}()
	go func() {
		for i := 2; i <= 100; i += 2 {
			<-ch2
			fmt.Println("evenNumber -->", i)
			ch1 <- true
		}
	}()

	ch1 <- true
	time.Sleep(2)
}
func TestWork3(t *testing.T) {
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	go func() {
		for i := 1; i <= 100; i += 2 {
			fmt.Println("oddNumber -->", i)
			ch2 <- true
			<-ch1
		}
	}()
	go func() {
		for i := 2; i <= 100; i += 2 {
			<-ch2
			fmt.Println("evenNumber -->", i)
			ch1 <- true
		}
	}()

	ch1 <- true // 启动打印奇数的goroutine

	// 等待程序执行完成
	time.Sleep(1 * time.Second)
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 63. 股票的最大利润
// 方式一：贪心
func maxProfit1(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}

	m := math.MaxInt64
	ans := 0

	for _, price := range prices {
		m = min(m, price)
		ans = max(ans, price-m)
	}
	return ans
}

// 方式二：dp
func maxProfit2(prices []int) int {

	dp := make([]int, len(prices))
	m := math.MaxInt64

	dp[0] = 0

	for i := 1; i < len(prices); i++ {
		m = min(m, prices[i])
		dp[i] = max(dp[i-1], prices[i]-m)
	}

	sort.Ints(dp)
	return dp[len(dp)-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func TestMaxProfit(t *testing.T) {
	fmt.Println(maxProfit2([]int{7, 1, 5, 3, 6, 4}))
}

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer 47. 礼物的最大价值
func maxValue(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, columns := len(grid), len(grid[0])
	// dp[i][j]: 表示走到 grid[i][j] 能获取到的最大礼物值
	dp := make([][]int, rows)
	for i, _ := range dp {
		dp[i] = make([]int, columns)
	}
	dp[0][0] = grid[0][0]

	for j := 1; j < columns; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[rows-1][columns-1]
}
