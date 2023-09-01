#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn three_sum(mut nums: Vec<i32>) -> Vec<Vec<i32>> {
        if nums.len() < 3 {
            return vec![];
        }
        let mut result = Vec::<Vec<i32>>::new();
        nums.sort();
        for i in 0..nums.len() {
            if nums[i] > 0 {
                break;
            }
            if i > 0 && nums[i] == nums[i - 1] {
                continue;
            }
            let (mut left, mut right) = (i + 1, nums.len() - 1);
            while left < right {
                let r = nums[i] + nums[left] + nums[right];
                if r == 0 {
                    result.push(vec![nums[i], nums[left], nums[right]]);
                    while left < right && nums[left] == nums[left + 1] {
                        left += 1;
                    }
                    while left < right && nums[right] == nums[right - 1] {
                        right -= 1;
                    }
                    left += 1;
                    right -= 1;
                }
                if r > 0 {
                    right -= 1;
                }
                if r < 0 {
                    left += 1;
                }
            }
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
            vec![vec![-1, -1, 2], vec![-1, 0, 1]],
            Solution::three_sum(vec![-1, 0, 1, 2, -1, -4])
        );
        assert_eq!(Vec::<Vec<i32>>::new(), Solution::three_sum(vec![0, 1, 1]));
        assert_eq!(vec![vec![0, 0, 0]], Solution::three_sum(vec![0, 0, 0]));
    }
}
