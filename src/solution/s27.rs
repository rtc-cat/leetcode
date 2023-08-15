#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn remove_element(nums: &mut Vec<i32>, val: i32) -> i32 {
        let mut left = 0 as usize;
        for i in 0..nums.len() {
            if nums[i] != val {
                nums[left] = nums[i];
                left += 1;
            }
        }
        return left as i32;
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_solution() {
        assert_eq!(2, Solution::remove_element(&mut vec![3, 2, 2, 3], 3));
        assert_eq!(
            5,
            Solution::remove_element(&mut vec![0, 1, 2, 2, 3, 0, 4, 2], 2)
        );
    }
}
