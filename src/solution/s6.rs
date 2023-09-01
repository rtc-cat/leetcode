#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn convert(s: String, num_rows: i32) -> String {
        if num_rows < 2 {
            return s;
        }
        let mut strs = vec![String::from(""); num_rows as usize];
        let (mut index, mut flag) = (0 as i32, 1 as i32);
        for ch in s.chars() {
            strs[index as usize].push(ch);
            if index == num_rows - 1 && flag == 1 {
                flag = -flag;
            }
            if index == 0 && flag == -1 {
                flag = -flag;
            }
            index += flag;
        }
        strs.join("")
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(
            String::from("PAHNAPLSIIGYIR"),
            Solution::convert(String::from("PAYPALISHIRING"), 3)
        );
        assert_eq!(
            String::from("PINALSIGYAHRPI"),
            Solution::convert(String::from("PAYPALISHIRING"), 4)
        );
        assert_eq!(String::from("A"), Solution::convert(String::from("A"), 1));
        assert_eq!(String::from("AB"), Solution::convert(String::from("Ab"), 1));
    }
}
