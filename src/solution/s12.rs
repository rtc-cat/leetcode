#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn int_to_roman(num: i32) -> String {
        let mut input = num;
        let values = vec![
            (1000, "M"),
            (900, "CM"),
            (500, "D"),
            (400, "CD"),
            (100, "C"),
            (90, "XC"),
            (50, "L"),
            (40, "XL"),
            (10, "X"),
            (9, "IX"),
            (5, "V"),
            (4, "IV"),
            (1, "I"),
        ];
        let mut result = String::new();
        for (n, s) in values {
            let mut count = input / n;
            while count > 0 {
                result.push_str(s);
                count -= 1;
            }
            input %= n;
        }
        result
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(String::from("III"), Solution::int_to_roman(3));
        assert_eq!(String::from("IV"), Solution::int_to_roman(4));
        assert_eq!(String::from("IX"), Solution::int_to_roman(9));
        assert_eq!(String::from("LVIII"), Solution::int_to_roman(58));
        assert_eq!(String::from("MCMXCIV"), Solution::int_to_roman(1994));
    }
}
