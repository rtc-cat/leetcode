#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    fn compress(chars: &mut Vec<char>) -> i32 {
        if chars.len() == 1 {
            return 1;
        }
        0
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        let test_cases = vec![
            ("1个字符", vec!['a'], 1),
            ("通过", vec!['a', 'a', 'b', 'b', 'c', 'c', 'c'], 0),
            (
                "多个字符",
                vec![
                    'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b',
                ],
                0,
            ),
        ];
        for (name, mut chars, result) in test_cases {
            println!("{}测试:", name);
            assert_eq!(Solution::compress(&mut chars), result);
        }
    }
}
