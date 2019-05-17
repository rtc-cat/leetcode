package main

import "fmt"

// 给定两个字符串 s 和 t，判断它们是否是同构的。

// 如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。

// 所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。

// 示例 1:

// 输入: s = "egg", t = "add"
// 输出: true
// 示例 2:

// 输入: s = "foo", t = "bar"
// 输出: false
// 示例 3:

// 输入: s = "paper", t = "title"
// 输出: true
// 说明:
// 你可以假设 s 和 t 具有相同的长度。

func main() {
	fmt.Println(isIsomorphic("ab", "aa"))
}

func isIsomorphic(s string, t string) bool {
	// 用map保存每一个字符 s到t的映射,如果出现了不一致则不对
	m := make(map[byte]byte, len(s))
	// 这个map保存t中的每个字符是否已经被映射过
	tm := make(map[byte]bool, len(t))
	for i := 0; i < len(s); i++ {
		si := s[i]
		ti := t[i]
		// 如果不存在
		mt, ok := m[si]
		if !ok {
			// 需要判断 si 和ti是否相等
			if si == ti {
				m[si] = ti    // 添加s->t映射关系
				tm[ti] = true // 认为已经被映射过
			} else {
				if tm[ti] { // 如果ti已经被映射过,就不能再映射
					return false
				}
				m[si] = ti
				tm[ti] = true
			}
			continue
		}
		// 如果已经被映射过需要判断是否和之前相同
		if mt != ti { // 如果不同说明不是同构数
			return false
		}
	}
	return true
}
