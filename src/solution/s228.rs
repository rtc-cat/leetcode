#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn summary_ranges(nums: Vec<i32>) -> Vec<String> {
        if nums.len() == 0 {
            return vec![];
        }
        let mut start = 0;
        let mut end = 0;
        let mut result = vec![];
        for i in 0..nums.len() {
            let j = i + 1;
            if j < nums.len() {
                if nums[j] == nums[i] + 1 {
                    end = j;
                } else {
                    if start == end {
                        result.push(nums[i].to_string());
                    } else {
                        result.push(format!("{}->{}", nums[start], nums[end]));
                    }
                    start = j;
                    end = j;
                }
            }
        }
        if start == end {
            result.push(nums[end].to_string());
        } else {
            result.push(format!("{}->{}", nums[start], nums[end]));
        }
        result
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::*;

    #[test]
    fn test_solution() {
        assert_eq!(
            vec_of_strings!["0->2", "4->5", "7"],
            Solution::summary_ranges(vec![0, 1, 2, 4, 5, 7])
        );
        assert_eq!(
            vec_of_strings!["0", "2->4", "6", "8->9"],
            Solution::summary_ranges(vec![0, 2, 3, 4, 6, 8, 9])
        );
    }
}
