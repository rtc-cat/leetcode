#[allow(dead_code)]
struct Solution {}

impl Solution {
    // 如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
    // 如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
    // 如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
    // 如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
    #[allow(dead_code)]
    pub fn game_of_life(board: &mut Vec<Vec<i32>>) {
        let current = board.clone();

        for i in 0..current.len() {
            for j in 0..current[i].len() {
                let living_counts = Solution::get_living_counts(&current, i, j);
                if current[i][j] == 1 {
                    if living_counts < 2 || living_counts > 3 {
                        board[i][j] = 0;
                    }
                } else {
                    if living_counts == 3 {
                        board[i][j] = 1;
                    }
                }
            }
        }
    }

    pub fn get_living_counts(board: &Vec<Vec<i32>>, row: usize, col: usize) -> i32 {
        let mut result = 0;
        if board.len() == 0 {
            return result;
        }
        let max_row = board.len() - 1;
        let max_col = board[0].len() - 1;
        // 左上
        if row > 0 && col > 0 {
            if board[row - 1][col - 1] == 1 {
                result += 1;
            }
        }
        // 上
        if row > 0 {
            if board[row - 1][col] == 1 {
                result += 1;
            }
        }
        // 右上
        if row > 0 && col < max_col {
            if board[row - 1][col + 1] == 1 {
                result += 1;
            }
        }
        // 左
        if col > 0 {
            if board[row][col - 1] == 1 {
                result += 1;
            }
        }
        // 右
        if col < max_col {
            if board[row][col + 1] == 1 {
                result += 1;
            }
        }
        // 左下
        if row < max_row && col > 0 {
            if board[row + 1][col - 1] == 1 {
                result += 1;
            }
        }
        // 下
        if row < max_row {
            if board[row + 1][col] == 1 {
                result += 1;
            }
        }
        // 右下
        if row < max_row && col < max_col {
            if board[row + 1][col + 1] == 1 {
                result += 1;
            }
        }
        result
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        let mut matrix_1 = vec![vec![0, 1, 0], vec![0, 0, 1], vec![1, 1, 1], vec![0, 0, 0]];
        let result_1 = vec![vec![0, 0, 0], vec![1, 0, 1], vec![0, 1, 1], vec![0, 1, 0]];
        Solution::game_of_life(&mut matrix_1);
        assert_eq!(result_1, matrix_1);

        let mut matrix_2 = vec![vec![1, 1], vec![1, 0]];
        let result_2 = vec![vec![1, 1], vec![1, 1]];
        Solution::game_of_life(&mut matrix_2);
        assert_eq!(result_2, matrix_2);
    }
}
