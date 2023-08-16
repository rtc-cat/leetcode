#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn can_jump(nums: Vec<i32>) -> bool {
        let mut max: usize = 0;
        for i in 0..nums.len() {
            if i <= max {
                max = std::cmp::max(max, i + nums[i] as usize);
                if max >= nums.len() - 1 {
                    return true;
                }
            }
        }
        false
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(true, Solution::can_jump(vec![2, 3, 1, 1, 4]));
        assert_eq!(false, Solution::can_jump(vec![3, 2, 1, 0, 4]));
    }
}
