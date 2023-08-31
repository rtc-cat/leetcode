#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn longest_common_prefix(strs: Vec<String>) -> String {
        let mut result = String::new();
        for (i, ch) in strs[0].char_indices() {
            for str in strs.iter() {
                if i >= str.len() {
                    return result;
                }
                if str.chars().nth(i).unwrap() != ch {
                    return result;
                }
            }
            result.push(ch);
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
            String::from("fl"),
            Solution::longest_common_prefix(vec![
                String::from("flower"),
                String::from("flow"),
                String::from("flight")
            ])
        );
        assert_eq!(
            String::from(""),
            Solution::longest_common_prefix(vec![
                String::from("dog"),
                String::from("racecar"),
                String::from("car")
            ])
        );
        assert_eq!(
            String::from("c"),
            Solution::longest_common_prefix(vec![String::from("cir"), String::from("car"),])
        );
    }
}
