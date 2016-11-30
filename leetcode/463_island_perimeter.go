package leetcode

// 463. Island Perimeter
func islandPerimeter(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	rowSize := len(grid)
	colSize := len(grid[0])

	var res int
	for i, rows := range grid {
		for j, cell := range rows {
			if cell == 0 {
				continue
			}
			if i-1 < 0 || grid[i-1][j] == 0 {
				res++
			}
			if i+1 >= rowSize || grid[i+1][j] == 0 {
				res++
			}
			if j-1 < 0 || grid[i][j-1] == 0 {
				res++
			}
			if j+1 >= colSize || grid[i][j+1] == 0 {
				res++
			}
		}
	}
	return res
}
