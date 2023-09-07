#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn is_happy(n: i32) -> bool {
        let mut repeated = std::collections::HashSet::new();
        let mut num = n;
        repeated.insert(num);
        loop {
            let mut result = 0;
            while num > 0 {
                let x = num % 10;
                result += x * x;
                num = num / 10;
            }
            if result == 1 {
                return true;
            }
            if !repeated.insert(result) {
                return false;
            }
            num = result
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(true, Solution::is_happy(19));
        assert_eq!(false, Solution::is_happy(2));
    }
}
