package main

// 统计所有小于非负整数 n 的质数的数量。

// 示例:

// 输入: 10
// 输出: 4
// 解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。

func main() {

}

// 厄拉多塞筛选法
// https://blog.csdn.net/gavinming/article/details/7212980
// 找到一个质数后,将质数的倍数,全部划去,下一个数字就是下一个质数
func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	all := make([]bool, n)
	for i := 2; i < n; i++ {
		all[i] = true
	}
	// 标记
	for i := 2; i*i < n; i++ {
		// 如果这个数是质数,那么把能被它整除的数全都去掉
		if all[i] {
			for j := 2; i*j < n; j++ {
				all[i*j] = false
			}
		}
	}
	var result int
	for i := 0; i < len(all); i++ {
		if all[i] {
			result++
		}
	}
	return result
}
