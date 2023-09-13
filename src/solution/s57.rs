#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn insert(intervals: Vec<Vec<i32>>, new_interval: Vec<i32>) -> Vec<Vec<i32>> {
        let mut result = vec![];
        let mut merged = false;
        let (mut left, mut right) = (new_interval[0], new_interval[1]);
        for i in 0..intervals.len() {
            if intervals[i][1] < left {
                result.push(intervals[i].clone());
            } else if intervals[i][0] > right {
                if !merged {
                    result.push(vec![left, right]);
                    merged = true;
                }
                result.push(intervals[i].clone());
            } else {
                left = std::cmp::min(left, intervals[i][0]);
                right = std::cmp::max(right, intervals[i][1]);
            }
        }
        if !merged {
            result.push(vec![left, right]);
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
            vec![vec![1, 5], vec![6, 9]],
            Solution::insert(vec![vec![1, 3], vec![6, 9]], vec![2, 5])
        );
        assert_eq!(
            vec![vec![1, 2], vec![3, 10], vec![12, 16]],
            Solution::insert(
                vec![
                    vec![1, 2],
                    vec![3, 5],
                    vec![6, 7],
                    vec![8, 10],
                    vec![12, 16]
                ],
                vec![4, 8]
            )
        );
        assert_eq!(vec![vec![5, 7]], Solution::insert(vec![], vec![5, 7]));
        assert_eq!(
            vec![vec![1, 5]],
            Solution::insert(vec![vec![1, 5]], vec![2, 3])
        );
    }
}
