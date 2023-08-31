#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn length_of_last_word(s: String) -> i32 {
        let words: Vec<&str> = s.split_whitespace().collect();
        if words.len() == 0 {
            return 0;
        }
        words[words.len() - 1].len() as i32
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(
            5,
            Solution::length_of_last_word(String::from("Hello World"))
        );
        assert_eq!(
            4,
            Solution::length_of_last_word(String::from("   fly me   to   the moon  "))
        );
        assert_eq!(
            6,
            Solution::length_of_last_word(String::from("luffy is still joyboy"))
        );
    }
}
