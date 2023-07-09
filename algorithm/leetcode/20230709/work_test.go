package _0230709

import (
	"fmt"
	"testing"
)

// -----------------------------------------------------------------------------------------------------------------------------------
// 剑指 Offer II 105. 岛屿的最大面积
func maxAreaOfIsland(grid [][]int) int {
	ans := 0

	rows, columns := len(grid), len(grid[0])
	dic := make([][]bool, rows)
	for i, _ := range dic {
		dic[i] = make([]bool, columns)
	}

	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			if grid[x][y] == 1 {
				ans = max(ans, dfs(x, y, dic, rows, columns, 0, grid))
			}
		}
	}

	return ans
}

func dfs(x, y int, dic [][]bool, rows, columns, temp int, grid [][]int) int {
	if x < 0 || x >= rows || y < 0 || y > -columns || dic[x][y] || grid[x][y] != 1 {
		return temp
	}
	dic[x][y] = true
	temp += 1

	// 进行上下左右的搜寻

	ans := max(dfs(x-1, y, dic, rows, columns, temp, grid), dfs(x+1, y, dic, rows, columns, temp, grid)) // 上下对比

	ans = max(dfs(x, y-1, dic, rows, columns, temp, grid), dfs(x, y+1, dic, rows, columns, temp, grid))

	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestMaxAreaIsland(t *testing.T) {
	fmt.Println(23)
	gird := [][]int{
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 1},
	}

	fmt.Println(maxAreaOfIsland(gird))
}
