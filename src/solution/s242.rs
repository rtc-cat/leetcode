#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn is_anagram(s: String, t: String) -> bool {
        let mut s_chars = s
            .chars()
            .fold(std::collections::HashMap::new(), |mut m, c| {
                *m.entry(c).or_insert(0) += 1;
                m
            });
        let t_chars: Vec<_> = t.chars().collect();

        for i in 0..t_chars.len() {
            if *s_chars
                .entry(t_chars[i])
                .and_modify(|c| *c -= 1)
                .or_insert(-1)
                < 0
            {
                return false;
            }
        }
        s_chars.values().all(|&value| value == 0)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(
            true,
            Solution::is_anagram(String::from("anagram"), String::from("nagaram"))
        );
        assert_eq!(
            false,
            Solution::is_anagram(String::from("rat"), String::from("car"))
        );
    }
}
