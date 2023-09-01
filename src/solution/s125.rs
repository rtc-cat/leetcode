#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn is_palindrome(s: String) -> bool {
        if s.len() == 0 {
            return true;
        }
        let (mut left, mut right) = (0, s.len() as i32 - 1);
        let chars: Vec<char> = s.chars().collect();
        loop {
            if left > right {
                break;
            }
            if !chars[left as usize].is_alphanumeric() {
                left += 1;
                continue;
            }
            if !chars[right as usize].is_alphanumeric() {
                right -= 1;
                continue;
            }

            if chars[left as usize].to_ascii_lowercase()
                != chars[right as usize].to_ascii_lowercase()
            {
                return false;
            }
            left += 1;
            right -= 1;
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
            Solution::is_palindrome(String::from("A man, a plan, a canal: Panama"))
        );
        assert_eq!(false, Solution::is_palindrome(String::from("race a car")));
        assert_eq!(true, Solution::is_palindrome(String::from(" ")));
        assert_eq!(false, Solution::is_palindrome(String::from("0P")));
    }
}
