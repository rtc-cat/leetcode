#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn contains_nearby_duplicate(nums: Vec<i32>, k: i32) -> bool {
        let mut indexes = std::collections::HashMap::new();
        for i in 0..nums.len() {
            let last = *indexes.entry(nums[i]).or_insert(i);
            if last != i && (i - last) <= k as usize {
                return true;
            }
            indexes.insert(nums[i], i);
        }
        false
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(
            true,
            Solution::contains_nearby_duplicate(vec![1, 2, 3, 1], 3)
        );
        assert_eq!(
            true,
            Solution::contains_nearby_duplicate(vec![1, 0, 1, 1], 1)
        );
        assert_eq!(
            false,
            Solution::contains_nearby_duplicate(vec![1, 2, 3, 1, 2, 3], 2)
        );
    }
}
