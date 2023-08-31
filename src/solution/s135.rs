#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn candy(ratings: Vec<i32>) -> i32 {
        if ratings.len() <= 1 {
            return ratings.len() as i32;
        }
        let mut candies = vec![1; ratings.len()];
        // 从左往右遍历
        for i in 1..ratings.len() {
            if ratings[i] > ratings[i - 1] && candies[i] <= candies[i - 1] {
                candies[i] = candies[i - 1] + 1;
            }
        }
        // 再从右往左遍历
        for i in (1..=ratings.len() - 1).rev() {
            if ratings[i - 1] > ratings[i] && candies[i - 1] <= candies[i] {
                candies[i - 1] = candies[i] + 1;
            }
        }
        candies.iter().sum()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(5, Solution::candy(vec![1, 0, 2]));
        assert_eq!(4, Solution::candy(vec![1, 2, 2]));
        assert_eq!(7, Solution::candy(vec![1, 3, 2, 2, 1]));
        assert_eq!(7, Solution::candy(vec![3, 2, 1]));
    }
}
