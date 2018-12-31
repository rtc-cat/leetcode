package main

import "fmt"

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

// 注意：给定 n 是一个正整数。

// 示例 1：

// 输入： 2
// 输出： 2
// 解释： 有两种方法可以爬到楼顶。
// 1.  1 阶 + 1 阶
// 2.  2 阶
// 示例 2：

// 输入： 3
// 输出： 3
// 解释： 有三种方法可以爬到楼顶。
// 1.  1 阶 + 1 阶 + 1 阶
// 2.  1 阶 + 2 阶
// 3.  2 阶 + 1 阶
func main() {
	fmt.Println(climbStairsB(4))
}

// 到达n的位置可以转换为到达n-1和n-2分别要多少种,以此类推
// 递归解法
// emmm...超时了
func climbStairs(n int) int {
	// 递归退出条件
	if n <= 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// 递归转换为循环
func climbStairsB(n int) int {
	if n == 1 {
		return 1
	}

	i, j := 1, 2
	for n > 2 {
		i, j = j, i+j
		n--
	}
	return j
}
