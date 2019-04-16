package main

// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

// 注意：

// num1 和num2 的长度都小于 5100.
// num1 和num2 都只包含数字 0-9.
// num1 和num2 都不包含任何前导零。
// 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。

func main() {

}

func addStrings(num1 string, num2 string) string {
	if num1 == "" || num1 == "0" {
		return num2
	}
	if num2 == "" || num2 == "0" {
		return num1
	}
	return ""
}

func add(a, b string) (result string, carry bool) {
	return
}
