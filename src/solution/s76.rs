#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn min_window(s: String, t: String) -> String {
        let mut chars_count = std::collections::HashMap::<char, i32>::new();
        for ch in t.chars() {
            *chars_count.entry(ch).or_insert(0) += 1;
        }
        // 滑动窗口遍历左边字符串, 如果map中还有值就
        let s_chars: Vec<char> = s.chars().collect();
        let (mut left, mut right) = (0, 0);
        let mut result = String::new();
        let mut current = String::new();
        loop {
            if Solution::is_empty(&chars_count) {
                if current.len() < result.len() || result.is_empty() {
                    result = current.clone();
                }
                let left_char = s_chars[left];
                if chars_count.contains_key(&left_char) {
                    chars_count.entry(left_char).and_modify(|count| *count += 1);
                }
                current.remove(0);
                left += 1;
                continue;
            }

            if right + 1 > s.len() {
                break;
            }

            let ch = s_chars[right];
            if chars_count.contains_key(&ch) {
                chars_count
                    .entry(ch)
                    .and_modify(|count| *count -= 1)
                    .or_insert(0);
            }
            current.push(ch);
            right += 1;
        }
        result
    }

    pub fn is_empty(chars_count: &std::collections::HashMap<char, i32>) -> bool {
        if chars_count.is_empty() {
            return true;
        }
        chars_count.values().all(|&v| v <= 0)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(
            String::from("cwae"),
            Solution::min_window(String::from("cabwefgewcwaefgcf"), String::from("cae"))
        );
        assert_eq!(
            String::from("BANC"),
            Solution::min_window(String::from("ADOBECODEBANC"), String::from("ABC"))
        );
        assert_eq!(
            String::from("a"),
            Solution::min_window(String::from("a"), String::from("a"))
        );
        assert_eq!(
            String::from(""),
            Solution::min_window(String::from("a"), String::from("aa"))
        );
    }
}
