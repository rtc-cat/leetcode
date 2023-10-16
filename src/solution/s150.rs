#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn eval_rpn(tokens: Vec<String>) -> i32 {
        let mut stack = std::collections::VecDeque::<String>::new();
        for token in tokens {
            match token.as_str() {
                "+" => {
                    let a: i32 = stack.pop_back().unwrap().parse().unwrap();
                    let b: i32 = stack.pop_back().unwrap().parse().unwrap();
                    let result = b + a;
                    stack.push_back(result.to_string());
                }
                "-" => {
                    let a: i32 = stack.pop_back().unwrap().parse().unwrap();
                    let b: i32 = stack.pop_back().unwrap().parse().unwrap();
                    let result = b - a;
                    stack.push_back(result.to_string());
                }
                "*" => {
                    let a: i32 = stack.pop_back().unwrap().parse().unwrap();
                    let b: i32 = stack.pop_back().unwrap().parse().unwrap();
                    let result = b * a;
                    stack.push_back(result.to_string());
                }
                "/" => {
                    let a: i32 = stack.pop_back().unwrap().parse().unwrap();
                    let b: i32 = stack.pop_back().unwrap().parse().unwrap();
                    let result = b / a;
                    stack.push_back(result.to_string());
                }
                _ => {
                    stack.push_back(token);
                }
            }
        }
        stack.pop_back().unwrap().parse().unwrap()
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::*;

    #[test]
    fn test_solution() {
        assert_eq!(
            9,
            Solution::eval_rpn(vec_of_strings!["2", "1", "+", "3", "*"])
        );
        assert_eq!(
            6,
            Solution::eval_rpn(vec_of_strings!["4", "13", "5", "/", "+"])
        );
        assert_eq!(
            22,
            Solution::eval_rpn(vec_of_strings![
                "10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"
            ])
        );
    }
}
