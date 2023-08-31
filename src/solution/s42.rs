#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn trap(height: Vec<i32>) -> i32 {
        if height.len() <= 1 {
            return 0;
        }
        // 找到左边和右边最矮的柱子
        let mut left_max = vec![0; height.len()];
        left_max[0] = height[0];
        for i in 1..height.len() {
            left_max[i] = std::cmp::max(left_max[i - 1], height[i]);
        }

        let mut right_max = vec![0; height.len()];
        right_max[height.len() - 1] = height[height.len() - 1];
        for i in (0..=height.len() - 2).rev() {
            right_max[i] = std::cmp::max(right_max[i + 1], height[i]);
        }

        let mut result = 0;
        for i in 0..height.len() {
            result += std::cmp::min(left_max[i], right_max[i]) - height[i];
        }
        result
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(6, Solution::trap(vec![0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]));
        assert_eq!(9, Solution::trap(vec![4, 2, 0, 3, 2, 5]));
    }
}
