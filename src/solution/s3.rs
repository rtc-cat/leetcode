#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn length_of_longest_substring(s: String) -> i32 {
        if s.len() == 0 {
            return 0;
        }
        let mut max = 0;
        let (mut left, mut right) = (0, 0);
        let mut map = std::collections::HashMap::<char, usize>::new();
        let chars: Vec<char> = s.chars().collect();
        loop {
            let ch = chars[right];
            match map.get(&ch) {
                Some(&last) => {
                    // 把中间的全部删掉
                    for i in left..=last {
                        map.remove(&chars[i]);
                    }
                    left = last + 1;
                    map.insert(ch, right);
                    right += 1;
                }
                None => {
                    map.insert(ch, right);
                    max = std::cmp::max(max, map.len());
                    right += 1;
                }
            }
            if right > chars.len() - 1 {
                break;
            }
        }
        max as i32
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(
            3,
            Solution::length_of_longest_substring(String::from("abcabcbb"))
        );
        assert_eq!(
            1,
            Solution::length_of_longest_substring(String::from("bbbbb"))
        );
        assert_eq!(
            3,
            Solution::length_of_longest_substring(String::from("pwwkew"))
        )
    }
}
