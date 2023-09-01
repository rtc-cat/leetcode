#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn str_str(haystack: String, needle: String) -> i32 {
        if haystack.len() < needle.len() {
            return -1;
        }
        let haystack_chars: Vec<char> = haystack.chars().collect();
        let needle_chars: Vec<char> = needle.chars().collect();
        let mut index = 0;
        loop {
            if index >= haystack_chars.len() {
                break;
            }
            let mut matched = true;
            for i in 0..needle_chars.len() {
                let begin = index + i;
                if begin >= haystack_chars.len() {
                    return -1;
                }
                if haystack_chars[begin] != needle_chars[i] {
                    index = index + 1;
                    matched = false;
                    break;
                }
            }
            if matched {
                return index as i32;
            }
        }
        -1
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        // assert_eq!(
        //     0,
        //     Solution::str_str(String::from("sadbutsad"), String::from("sad"))
        // );
        // assert_eq!(
        //     -1,
        //     Solution::str_str(String::from("leetcode"), String::from("leeto"))
        // );
        // assert_eq!(
        //     -1,
        //     Solution::str_str(String::from("aaa"), String::from("aaaa"))
        // );
        assert_eq!(
            4,
            Solution::str_str(String::from("mississippi"), String::from("issip"))
        );
    }
}
