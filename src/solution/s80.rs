#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
        let k = 2;
        if nums.len() <= k {
            return nums.len() as i32;
        }
        let mut u = 0;
        for i in 0..nums.len() {
            if u < k || nums[u - k] != nums[i] {
                nums[u] = nums[i];
                u += 1;
            }
        }
        u as i32
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(5, Solution::remove_duplicates(&mut vec![1, 1, 1, 2, 2, 3]));
        assert_eq!(
            7,
            Solution::remove_duplicates(&mut vec![0, 0, 1, 1, 1, 1, 2, 3, 3])
        );
    }
}
