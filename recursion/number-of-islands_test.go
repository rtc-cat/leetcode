package recursion

import "testing"

func TestNumIslands(t *testing.T) {
	testCases := []struct {
		input [][]byte
	}{
		{[][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}},
		{[][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		}},
	}
	for _, tc := range testCases {
		t.Log(numIslands(tc.input))
	}
}

func numIslands(grid [][]byte) int {
	var visited [][]bool
	for _, g := range grid {
		visited = append(visited, make([]bool, len(g)))
	}
	var result int
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if visited[x][y] {
				continue
			}
			if grid[x][y] == '1' {
				expand(grid, &visited, x, y)
				result++
			}
		}
	}
	return result
}

func expand(grid [][]byte, visited *[][]bool, x, y int) {
	if x < 0 || x >= len(grid) {
		return
	}
	if y < 0 || y >= len(grid[x]) {
		return
	}
	if (*visited)[x][y] {
		return
	}
	if grid[x][y] == '0' {
		return
	}
	(*visited)[x][y] = true
	expand(grid, visited, x-1, y)
	expand(grid, visited, x+1, y)
	expand(grid, visited, x, y-1)
	expand(grid, visited, x, y+1)
}
