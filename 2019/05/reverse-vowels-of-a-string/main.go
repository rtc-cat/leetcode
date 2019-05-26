package main

// 编写一个函数，以字符串作为输入，反转该字符串中的元音字母。

// 示例 1:

// 输入: "hello"
// 输出: "holle"
// 示例 2:

// 输入: "leetcode"
// 输出: "leotcede"
// 说明:
// 元音字母不包含字母"y"。

func main() {

}

func reverseVowels(s string) string {
	// 只处理 a e i o u A E I O U
	m := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	b := []byte(s)
	// 两个指针左右开始
	for i, j := 0, len(b)-1; i < j; {
		// 如果两个都是元音,则交换
		if m[b[i]] && m[b[j]] {
			b[i], b[j] = b[j], b[i]
			// 移动两个指针
			i, j = i+1, j-1
			continue
		}
		if !m[b[i]] {
			i++
		}
		if !m[b[j]] {
			j--
		}
	}
	return string(b)
}
