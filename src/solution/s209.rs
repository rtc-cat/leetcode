#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn min_sub_array_len(target: i32, nums: Vec<i32>) -> i32 {
        if nums.len() == 1 {
            if nums[0] >= target {
                return 1;
            }
            return 0;
        }
        let mut sum = 0;
        let (mut left, mut right) = (0, 0);
        let mut min = 0;
        sum += nums[right];
        loop {
            match sum.cmp(&target) {
                std::cmp::Ordering::Less => {
                    right += 1;
                    if right > nums.len() - 1 {
                        break;
                    }
                    sum += nums[right];
                }
                _ => {
                    if min == 0 {
                        min = right - left + 1;
                    } else {
                        min = std::cmp::min(min, right - left + 1);
                    }
                    // 向右移动
                    sum -= nums[left];
                    left += 1;
                    if right < left {
                        right = left
                    }
                }
            }
        }
        min as i32
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(2, Solution::min_sub_array_len(7, vec![2, 3, 1, 2, 4, 3]));
        assert_eq!(1, Solution::min_sub_array_len(4, vec![1, 4, 4]));
        assert_eq!(
            0,
            Solution::min_sub_array_len(11, vec![1, 1, 1, 1, 1, 1, 1, 1])
        );
    }
}
