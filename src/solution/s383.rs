#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn can_construct(ransom_note: String, magazine: String) -> bool {
        let mut chars = magazine
            .chars()
            .fold(std::collections::HashMap::new(), |mut m, ch| {
                *m.entry(ch).or_insert(0) += 1;
                m
            });

        for ch in ransom_note.chars() {
            chars.entry(ch).and_modify(|c| *c -= 1).or_insert(-1);
            if chars[&ch] < 0 {
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
            Solution::can_construct(String::from("aa"), String::from("aab"))
        );
        assert_eq!(
            false,
            Solution::can_construct(String::from("a"), String::from("b"))
        );
        assert_eq!(
            false,
            Solution::can_construct(String::from("aa"), String::from("ab"))
        );
    }
}
