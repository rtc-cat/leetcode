#[allow(dead_code)]
struct Solution {}

impl Solution {
    // I             1
    // V             5
    // X             10
    // L             50
    // C             100
    // D             500
    // M             1000
    #[allow(dead_code)]
    pub fn roman_to_int(s: String) -> i32 {
        let values = std::collections::HashMap::from([
            ('I', 1),
            ('V', 5),
            ('X', 10),
            ('L', 50),
            ('C', 100),
            ('D', 500),
            ('M', 1000),
        ]);
        let chars: Vec<char> = s.chars().collect();

        let mut result = 0;

        for i in 0..chars.len() {
            if i < chars.len() - 1 {
                let current = values.get(&chars[i]).unwrap();
                let next = values.get(&chars[i + 1]).unwrap();
                if current >= next {
                    result += current;
                } else {
                    result -= current;
                }
            }
        }
        result += values.get(&chars[chars.len() - 1]).unwrap();
        result
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(3, Solution::roman_to_int(String::from("III")));
        assert_eq!(4, Solution::roman_to_int(String::from("IV")));
        assert_eq!(9, Solution::roman_to_int(String::from("IX")));
        assert_eq!(58, Solution::roman_to_int(String::from("LVIII")));
        assert_eq!(1994, Solution::roman_to_int(String::from("MCMXCIV")));
    }
}
