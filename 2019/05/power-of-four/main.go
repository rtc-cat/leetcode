package main

// 给定一个整数 (32 位有符号整数)，请编写一个函数来判断它是否是 4 的幂次方。

// 示例 1:

// 输入: 16
// 输出: true
// 示例 2:

// 输入: 5
// 输出: false
// 进阶：
// 你能不使用循环或者递归来完成本题吗？

func main() {

}

func isPowerOfFour(num int) bool {
	// 先排除不是2的幂
	if num < 0 {
		return false
	}
	// 16= 10000
	// 15= 01111
	if num&(num-1) != 0 {
		return false
	}
	// 判断是不是在奇数位
	// 0101 0101 0101 0101 0101 0101 0101 0101
	//0x55555555
	if num&0x55555555 == 0 {
		return false
	}
	return true
}
