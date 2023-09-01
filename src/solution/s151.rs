#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn reverse_words(s: String) -> String {
        let mut words = Vec::new();
        let mut current_word = String::new();
        for ch in s.chars() {
            if ch == ' ' {
                if current_word.len() > 0 {
                    words.push(current_word.clone());
                    current_word.clear();
                }
                continue;
            }
            current_word.push(ch);
        }
        if current_word.len() > 0 {
            words.push(current_word.clone());
        }
        words.reverse();
        words.join(" ")
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(
            String::from("the sky is blue"),
            Solution::reverse_words(String::from("blue is sky the"))
        );
        assert_eq!(
            String::from("world hello"),
            Solution::reverse_words(String::from("  hello world  "))
        );
        assert_eq!(
            String::from("example good a"),
            Solution::reverse_words(String::from("a good   example"))
        );
    }
}
