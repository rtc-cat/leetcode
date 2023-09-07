#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn is_isomorphic(s: String, t: String) -> bool {
        let mut s_t = std::collections::HashMap::new();
        let mut t_s = std::collections::HashMap::new();
        let s_chars: Vec<char> = s.chars().collect();
        let t_chars: Vec<char> = t.chars().collect();
        for i in 0..s_chars.len() {
            let s_ch = s_chars[i];
            let t_ch = t_chars[i];

            if t_ch != *s_t.entry(s_ch).or_insert(t_ch) {
                return false;
            }
            if s_ch != *t_s.entry(t_ch).or_insert(s_ch) {
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
            Solution::is_isomorphic(String::from("egg"), String::from("add"))
        );
        assert_eq!(
            false,
            Solution::is_isomorphic(String::from("foo"), String::from("bar"))
        );
        assert_eq!(
            true,
            Solution::is_isomorphic(String::from("paper"), String::from("title"))
        );
    }
}
