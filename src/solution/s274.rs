#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn h_index(mut citations: Vec<i32>) -> i32 {
        // h的范围 1<= h <= citations.len()
        citations.sort_by(|a, b| a.cmp(b).reverse());
        let mut h = 0;
        for i in 0..citations.len() {
            if citations[i] > h {
                h += 1;
            }
        }
        h
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(3, Solution::h_index(vec![3, 0, 6, 1, 5]));
        assert_eq!(3, Solution::h_index(vec![0, 1, 3, 5, 6]));
        assert_eq!(1, Solution::h_index(vec![1, 3, 1]));
    }
}
