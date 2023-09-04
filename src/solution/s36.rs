#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn is_valid_sudoku(board: Vec<Vec<char>>) -> bool {
        // 先判断行
        let mut nums = std::collections::HashSet::<char>::with_capacity(9);
        for row in 0..9 {
            for col in 0..9 {
                if board[row][col] != '.' {
                    if !nums.insert(board[row][col]) {
                        return false;
                    }
                }
            }
            nums.clear();
        }
        // 再判断列
        for col in 0..9 {
            for row in 0..9 {
                if board[row][col] != '.' {
                    if !nums.insert(board[row][col]) {
                        return false;
                    }
                }
            }
            nums.clear();
        }
        for i in 0..3 {
            for j in 0..3 {
                // 00, 01
                for row in 0..3 {
                    for col in 0..3 {
                        let row_index = i * 3 + row;
                        let col_index = j * 3 + col;
                        if board[row_index][col_index] != '.' {
                            if !nums.insert(board[row_index][col_index]) {
                                return false;
                            }
                        }
                    }
                }
                nums.clear();
            }
        }
        true
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(
            false,
            Solution::is_valid_sudoku(vec![
                vec!['.', '.', '.', '.', '5', '.', '.', '1', '.'],
                vec!['.', '4', '.', '3', '.', '.', '.', '.', '.'],
                vec!['.', '.', '.', '.', '.', '3', '.', '.', '1'],
                vec!['8', '.', '.', '.', '.', '.', '.', '2', '.'],
                vec!['.', '.', '2', '.', '7', '.', '.', '.', '.'],
                vec!['.', '1', '5', '.', '.', '.', '.', '.', '.'],
                vec!['.', '.', '.', '.', '.', '2', '.', '.', '.'],
                vec!['.', '2', '.', '9', '.', '.', '.', '.', '.'],
                vec!['.', '.', '4', '.', '.', '.', '.', '.', '.']
            ])
        );
        assert_eq!(
            true,
            Solution::is_valid_sudoku(vec![
                vec!['5', '3', '.', '.', '7', '.', '.', '.', '.'],
                vec!['6', '.', '.', '1', '9', '5', '.', '.', '.'],
                vec!['.', '9', '8', '.', '.', '.', '.', '6', '.'],
                vec!['8', '.', '.', '.', '6', '.', '.', '.', '3'],
                vec!['4', '.', '.', '8', '.', '3', '.', '.', '1'],
                vec!['7', '.', '.', '.', '2', '.', '.', '.', '6'],
                vec!['.', '6', '.', '.', '.', '.', '2', '8', '.'],
                vec!['.', '.', '.', '4', '1', '9', '.', '.', '5'],
                vec!['.', '.', '.', '.', '8', '.', '.', '7', '9'],
            ])
        );
        assert_eq!(
            false,
            Solution::is_valid_sudoku(vec![
                vec!['8', '3', '.', '.', '7', '.', '.', '.', '.'],
                vec!['6', '.', '.', '1', '9', '5', '.', '.', '.'],
                vec!['.', '9', '8', '.', '.', '.', '.', '6', '.'],
                vec!['8', '.', '.', '.', '6', '.', '.', '.', '3'],
                vec!['4', '.', '.', '8', '.', '3', '.', '.', '1'],
                vec!['7', '.', '.', '.', '2', '.', '.', '.', '6'],
                vec!['.', '6', '.', '.', '.', '.', '2', '8', '.'],
                vec!['.', '.', '.', '4', '1', '9', '.', '.', '5'],
                vec!['.', '.', '.', '.', '8', '.', '.', '7', '9'],
            ])
        );
    }
}
