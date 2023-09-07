#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn word_pattern(pattern: String, s: String) -> bool {
        let words: Vec<_> = s.split(' ').collect();
        if pattern.len() != words.len() {
            return false;
        }
        let chars: Vec<_> = pattern.chars().collect();
        let mut p_s = std::collections::HashMap::new();
        let mut s_p = std::collections::HashMap::new();

        for i in 0..chars.len() {
            let p_char = chars[i];
            let word = words[i];
            if *p_s.entry(p_char).or_insert(word) != word {
                return false;
            }
            if *s_p.entry(word).or_insert(p_char) != p_char {
                return false;
            }
        }
        true
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(
            true,
            Solution::word_pattern(String::from("abba"), String::from("dog cat cat dog"))
        );
        assert_eq!(
            false,
            Solution::word_pattern(String::from("abba"), String::from("dog cat cat fish"))
        );
        assert_eq!(
            false,
            Solution::word_pattern(String::from("aaaa"), String::from("dog cat cat dog"))
        );
    }
}
