#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn num_jewels_in_stones(jewels: String, stones: String) -> i32 {
        let mut char_set = std::collections::HashSet::new();
        for c in jewels.chars() {
            char_set.insert(c);
        }
        let mut result = 0;
        for c in stones.chars() {
            if char_set.contains(&c) {
                result += 1;
            }
        }
        result
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_solution() {
        assert_eq!(
            3,
            Solution::num_jewels_in_stones(String::from("aA"), String::from("aAAbbbb"))
        );
        assert_eq!(
            0,
            Solution::num_jewels_in_stones(String::from("z"), String::from("ZZ"))
        );
    }
}
