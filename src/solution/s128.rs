#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn longest_consecutive(nums: Vec<i32>) -> i32 {
        if nums.len() == 0 {
            return 0;
        }
        let nums_set: std::collections::HashSet<i32> = nums.clone().into_iter().collect();
        let mut result = 1;
        for i in 0..nums.len() {
            let n = nums[i];
            let last = n - 1;
            if nums_set.contains(&last) {
                continue;
            }
            let mut next = n + 1;
            let mut current = 1;
            loop {
                if !nums_set.contains(&next) {
                    break;
                }
                current += 1;
                next = next + 1;
            }
            result = std::cmp::max(result, current);
        }
        result
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(4, Solution::longest_consecutive(vec![100, 4, 200, 1, 3, 2]));
        assert_eq!(
            9,
            Solution::longest_consecutive(vec![0, 3, 7, 2, 5, 8, 4, 6, 0, 1])
        );
    }
}
