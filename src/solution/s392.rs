#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn is_subsequence(s: String, t: String) -> bool {
        if s.len() == 0 {
            return true;
        }
        let chars: Vec<char> = s.chars().collect();
        let mut index = 0;
        for ch in t.chars() {
            if ch != chars[index] {
                continue;
            }
            if index == chars.len() - 1 {
                return true;
            }
            index += 1;
        }
        return false;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(
            true,
            Solution::is_subsequence(String::from("abc"), String::from("ahbgdc"))
        );
        assert_eq!(
            false,
            Solution::is_subsequence(String::from("axc"), String::from("ahbgdc"))
        );
    }
}
