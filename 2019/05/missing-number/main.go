package main

// 给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。

// 示例 1:

// 输入: [3,0,1]
// 输出: 2
// 示例 2:

// 输入: [9,6,4,2,3,5,7,0,1]
// 输出: 8
// 说明:
// 你的算法应具有线性时间复杂度。你能否仅使用额外常数空间来实现?

func main() {

}

func missingNumber(nums []int) int {
	// 0,1,2,3 len=4
	// 0,1,x,3 len=3
	n := len(nums)
	result := (n * (1 + n)) / 2
	for _, i := range nums {
		result -= i
	}
	return result
}

func missingNumber2(nums []int) int {
	// 异或运算,因为index和value是一一对应
	result := 0
	for i := range nums {
		result ^= i
		result ^= nums[i]
	}
	result ^= len(nums)
	return result
}
