package main

import (
	"fmt"
	"strings"
)

// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

// 说明：本题中，我们将空字符串定义为有效的回文串。

// 示例 1:

// 输入: "A man, a plan, a canal: Panama"
// 输出: true
// 示例 2:

// 输入: "race a car"
// 输出: false

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("race a car"))                     // false
	fmt.Println(isPalindrome("0P"))                             // false
	fmt.Println(isPalindrome(" "))                              // true
	fmt.Println(isPalindrome("a"))                              // true
}

func isPalindrome(s string) bool {
	str := strings.ToLower(s)
	// 预处理,把所有非字母的全都清理
	var ra []rune
	for _, r := range str {
		if isLetterOrNumber(byte(r)) {
			ra = append(ra, r)
		}
	}
	s = string(ra)
	left := 0
	right := len(s) - 1
	match := 0
	for left < right {
		rl := s[left]
		rr := s[right]
		if !isLetterOrNumber(rl) {
			left++
			continue
		}
		if !isLetterOrNumber(rr) {
			right--
			continue
		}
		if rl != rr {
			return false
		}
		left++
		right--
		match++
	}
	if match == 0 {
		if len(s) <= 1 {
			return true
		}
		return false
	}
	return true
}

func isLetterOrNumber(b byte) bool {
	if b >= 'a' && b <= 'z' {
		return true
	}
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}
