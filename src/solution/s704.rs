#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        if nums.len() <= 0 {
            return -1;
        }
        let (mut left, mut right) = (0, nums.len() - 1);
        while left <= right {
            // 防止溢出
            let mid = left + (right - left) / 2;
            match nums[mid].cmp(&target) {
                std::cmp::Ordering::Equal => return mid as i32,
                std::cmp::Ordering::Less => {
                    left = mid + 1;
                }
                std::cmp::Ordering::Greater => {
                    if mid == 0 {
                        return -1;
                    }
                    right = mid - 1;
                }
            }
        }
        -1
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_solution() {
        assert_eq!(Solution::search(vec![], 1), -1);
        assert_eq!(Solution::search(vec![2, 5], 0), -1);
        assert_eq!(Solution::search(vec![5], -5), -1);
        assert_eq!(Solution::search(vec![-1, 0, 3, 5, 9, 12], 9), 4);
        assert_eq!(Solution::search(vec![-1, 0, 3, 5, 9, 12], 2), -1);
    }
}
