package main

import (
	"fmt"
)

// 给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

// 返回 s 所有可能的分割方案。

// 示例:

// 输入: "aab"
// 输出:
// [
//   ["aa","b"],
//   ["a","a","b"]
// ]

func main() {
	// fmt.Println(partition("aabbaac"))
	fmt.Println(check("abba"))
}

func partition(s string) [][]string {
	var result = make([][]string, 0, 0)
	if len(s) == 0 {
		return result
	}
	chars := []rune(s)
	// 子串长度为1,最起码有一种,就是单个字符
	sigle := make([]string, 0, 0)
	for _, char := range chars {
		sigle = append(sigle, string(char))
	}
	result = append(result, sigle)
	// 当前存储的字符串
	var (
		cache  string
		length = 2
	)
	// 子串长度为length开始
	for length < len(s) {
		for {
			cache = s[:length]
			if check(cache) {
			}
		}
	}

	return result
}

// 检测是否为回文字符串
func check(arr string) bool {
	left := 0
	right := len(arr) - 1
	ok := true
	for left < right {
		if !(arr[left] == arr[right]) {
			ok = false
			break
		}
		left++
		right--
	}
	return ok
}
