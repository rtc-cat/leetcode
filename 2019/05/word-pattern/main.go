package main

import (
	"strings"
)

// 给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。

// 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

// 示例1:

// 输入: pattern = "abba", str = "dog cat cat dog"
// 输出: true
// 示例 2:

// 输入:pattern = "abba", str = "dog cat cat fish"
// 输出: false
// 示例 3:

// 输入: pattern = "aaaa", str = "dog cat cat dog"
// 输出: false
// 示例 4:

// 输入: pattern = "abba", str = "dog dog dog dog"
// 输出: false
// 说明:
// 你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。

func main() {

}

func wordPattern(pattern string, str string) bool {
	arr := strings.Split(str, " ")
	if len(pattern) != len(arr) {
		return false
	}
	// 映射关系
	p2s := make(map[byte]string, len(pattern))
	s2p := make(map[string]bool, len(pattern))
	for i := 0; i < len(pattern); i++ {
		r, ok := p2s[pattern[i]]
		if !ok {
			p2s[pattern[i]] = arr[i]
			if s2p[arr[i]] {
				return false
			}
			s2p[arr[i]] = true
			continue
		}
		if r != arr[i] {
			return false
		}
	}
	return true
}
