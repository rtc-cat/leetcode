#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
        if nums.len() <= 1 {
            return nums.len() as i32;
        }
        let mut slow = 1 as usize;
        for fast in 1..nums.len() {
            if nums[fast] > nums[fast - 1] {
                nums[slow] = nums[fast];
                slow += 1;
            }
        }
        slow as i32
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(2, Solution::remove_duplicates(&mut vec![1, 1, 2]));
        assert_eq!(
            5,
            Solution::remove_duplicates(&mut vec![0, 0, 1, 1, 1, 2, 2, 3, 3, 4])
        );
    }
}
