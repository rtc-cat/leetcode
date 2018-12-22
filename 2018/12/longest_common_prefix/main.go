package main

import "fmt"

// 编写一个函数来查找字符串数组中的最长公共前缀。

// 如果不存在公共前缀，返回空字符串 ""。

// 示例 1:

// 输入: ["flower","flow","flight"]
// 输出: "fl"
// 示例 2:

// 输入: ["dog","racecar","car"]
// 输出: ""
// 解释: 输入不存在公共前缀。
// 说明:

// 所有输入只包含小写字母 a-z 。
func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func longestCommonPrefix(strs []string) string {
	// 选出数组中最短的那个字符串
	if len(strs) == 0 {
		return ""
	}
	var min = strs[0]
	for _, s := range strs {
		if len(s) < len(min) {
			min = s
		}
	}
	// 以这个字符串为标准开始对比所有的字符串
	var (
		compare, result string
	)
	for _, r := range min {
		compare += string(r)
		for _, str := range strs {
			if !(str[:len(compare)] == compare) {
				return result
			}
		}
		result = compare
	}
	return result
}
