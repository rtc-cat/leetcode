#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn is_valid(s: String) -> bool {
        let mut stack = std::collections::VecDeque::new();
        let chars: Vec<_> = s.chars().collect();

        let left = vec!['(', '[', '{'];
        let right = vec![')', ']', '}'];

        for i in 0..chars.len() {
            if left.contains(&chars[i]) {
                stack.push_back(chars[i]);
            }
            if right.contains(&chars[i]) {
                if stack.len() == 0 {
                    return false;
                }
                let c = stack.pop_back().unwrap();
                let left_index = left.iter().position(|&x| x == c).unwrap_or(4);
                if left_index == 4 {
                    return false;
                }
                let right_index = right.iter().position(|&x| x == chars[i]).unwrap();
                if left_index != right_index {
                    return false;
                }
            }
        }
        if stack.len() != 0 {
            return false;
        }
        true
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(true, Solution::is_valid(String::from("()")));
        assert_eq!(true, Solution::is_valid(String::from("()[]{}")));
        assert_eq!(false, Solution::is_valid(String::from("(]")));
        assert_eq!(false, Solution::is_valid(String::from("]")));
        assert_eq!(false, Solution::is_valid(String::from("[")));
    }
}
