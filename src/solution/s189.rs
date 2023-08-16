#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn rotate(nums: &mut Vec<i32>, k: i32) {
        let k: usize = k as usize % nums.len();
        if nums.len() <= 1 || k == 0 {
            return;
        }
        Self::reverse(nums, 0, nums.len() - 1);
        Self::reverse(nums, 0, k - 1);
        Self::reverse(nums, k, nums.len() - 1);
    }

    pub fn reverse(nums: &mut Vec<i32>, start: usize, end: usize) {
        let (mut i, mut j) = (start, end);
        while i < j {
            nums.swap(i, j);
            i += 1;
            j -= 1;
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        let mut v1 = vec![1, 2, 3, 4, 5, 6, 7];
        Solution::rotate(&mut v1, 3);
        assert_eq!(vec![5, 6, 7, 1, 2, 3, 4], v1);
    }
}
