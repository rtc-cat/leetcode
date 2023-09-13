#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn merge(mut intervals: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        if intervals.len() == 0 {
            return vec![];
        }
        if intervals.len() == 1 {
            return intervals;
        }
        intervals.sort_by(|a, b| a[0].cmp(&b[0]));
        let mut merged = vec![intervals[0].clone()];
        for i in 1..intervals.len() {
            let length = merged.len();
            let last = merged.get_mut(length - 1).unwrap();
            if intervals[i][0] > last[1] {
                merged.push(intervals[i].clone());
            } else {
                last[1] = std::cmp::max(last[1], intervals[i][1]);
            }
        }
        merged
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(
            vec![vec![1, 3], vec![4, 7],],
            Solution::merge(vec![
                vec![2, 3],
                vec![2, 2],
                vec![3, 3],
                vec![1, 3],
                vec![5, 7],
                vec![2, 2],
                vec![4, 6]
            ])
        );
        assert_eq!(
            vec![vec![1, 6], vec![8, 10], vec![15, 18]],
            Solution::merge(vec![vec![1, 3], vec![2, 6], vec![8, 10], vec![15, 18]])
        );
        assert_eq!(
            vec![vec![1, 5]],
            Solution::merge(vec![vec![1, 4], vec![4, 5]])
        );
    }
}
