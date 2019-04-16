package main

import (
	"fmt"
)

// 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

// 在杨辉三角中，每个数是它左上方和右上方的数的和。

// 示例:

// 输入: 5
// 输出:
// [
//      [1],
//     [1,1],
//    [1,2,1],
//   [1,3,3,1],
//  [1,4,6,4,1]
// ]

func main() {
	r := generate(5)
	for _, a := range r {
		fmt.Println(a)
	}
}

func generate(numRows int) [][]int {
	var result [][]int
	for i := 1; i <= numRows; i++ {
		result = append(result, line(i))
	}
	return result
}

func line(num int) []int {
	if num == 1 {
		return []int{1}
	}
	if num == 2 {
		return []int{1, 1}
	}
	var (
		result   []int
		previous []int
	)
	previous = line(num - 1)
	for i := 0; i < num; i++ {
		if i == 0 || i == num-1 {
			result = append(result, 1)
			continue
		}
		result = append(result, previous[i-1]+previous[i])
	}
	return result
}
