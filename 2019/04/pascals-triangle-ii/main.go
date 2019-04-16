package main

import "fmt"

// 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。

// 在杨辉三角中，每个数是它左上方和右上方的数的和。

// 示例:

// 输入: 3
// 输出: [1,3,3,1]
// 进阶：

// 你可以优化你的算法到 O(k) 空间复杂度吗？
func main() {
	fmt.Println(getRow(3))
}

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	previous := getRow(rowIndex - 1)
	index := rowIndex + 1
	result := make([]int, 0, index)
	for i := 0; i < index; i++ {
		if i == 0 || i == index-1 {
			result = append(result, 1)
			continue
		}
		result = append(result, previous[i-1]+previous[i])
	}
	return result
}
