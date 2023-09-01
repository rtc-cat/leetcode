#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn two_sum(numbers: Vec<i32>, target: i32) -> Vec<i32> {
        let (mut left, mut right) = (0, numbers.len() - 1);
        loop {
            if left > right {
                break;
            }
            let result = numbers[left] + numbers[right];
            if result == target {
                return vec![(left + 1) as i32, (right + 1) as i32];
            }
            if result < target {
                left += 1;
            }
            if result > target {
                right -= 1;
            }
        }
        vec![0, 0]
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(vec![1, 2], Solution::two_sum(vec![2, 7, 11, 15], 9));
        assert_eq!(vec![1, 3], Solution::two_sum(vec![2, 3, 4], 6));
        assert_eq!(vec![1, 2], Solution::two_sum(vec![-1, 0], -1));
    }
}
