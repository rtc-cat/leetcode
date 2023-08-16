#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        let mut last = nums[0];
        let mut count = 0;
        for n in nums {
            if n == last {
                count += 1;
            } else {
                count -= 1;
            }
            if count == 0 {
                last = n;
                count = 1
            }
        }
        last
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(3, Solution::majority_element(vec![3, 2, 3]));
        assert_eq!(2, Solution::majority_element(vec![2, 2, 1, 1, 1, 2, 2]));
    }
}
