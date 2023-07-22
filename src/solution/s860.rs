#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn lemonade_change(bills: Vec<i32>) -> bool {
        let (mut five, mut ten) = (0, 0);
        for bill in bills.iter() {
            match bill {
                5 => {
                    five += 1;
                }
                10 => {
                    // 1个5都没有的话, 就返回false
                    if five < 1 {
                        return false;
                    }
                    five -= 1;
                    ten += 1;
                }
                20 => {
                    // 3个5, 或者1个10,1个5
                    if ten > 0 && five > 0 {
                        five -= 1;
                        ten -= 1;
                    } else if five > 2 {
                        five -= 3
                    } else {
                        return false;
                    }
                }
                _ => {}
            }
        }
        return true;
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_solution() {
        assert_eq!(true, Solution::lemonade_change(vec![5, 5, 5, 10, 20]));
        assert_eq!(false, Solution::lemonade_change(vec![5, 5, 10, 10, 20]));
    }
}
