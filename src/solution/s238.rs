#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
        let mut left = vec![0; nums.len()];
        let mut right = vec![0; nums.len()];

        left[0] = 1;
        right[nums.len() - 1] = 1;

        for i in 1..nums.len() {
            left[i] = left[i - 1] * nums[i - 1];
        }

        for i in (0..=nums.len() - 2).rev() {
            right[i] = right[i + 1] * nums[i + 1];
        }

        let mut result = vec![0; nums.len()];
        for i in 0..nums.len() {
            result[i] = left[i] * right[i];
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
            vec![24, 12, 8, 6],
            Solution::product_except_self(vec![1, 2, 3, 4])
        );
        assert_eq!(
            vec![0, 0, 9, 0, 0],
            Solution::product_except_self(vec![-1, 1, 0, -3, 3])
        );
    }
}
