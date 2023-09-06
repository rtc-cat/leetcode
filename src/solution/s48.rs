#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn rotate(matrix: &mut Vec<Vec<i32>>) {
        let n = matrix.len();
        for i in 0..n / 2 {
            for j in 0..(n + 1) / 2 {
                // 把第一个拿出来
                let temp = matrix[i][j];
                matrix[i][j] = matrix[n - j - 1][i];
                matrix[n - j - 1][i] = matrix[n - i - 1][n - j - 1];
                matrix[n - i - 1][n - j - 1] = matrix[j][n - i - 1];
                matrix[j][n - i - 1] = temp;
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        let mut matrix_1 = vec![vec![1, 2, 3], vec![4, 5, 6], vec![7, 8, 9]];
        let result_1 = vec![vec![7, 4, 1], vec![8, 5, 2], vec![9, 6, 3]];
        Solution::rotate(&mut matrix_1);
        assert_eq!(result_1, matrix_1);

        let mut matrix_2 = vec![
            vec![5, 1, 9, 11],
            vec![2, 4, 8, 10],
            vec![13, 3, 6, 7],
            vec![15, 14, 12, 16],
        ];
        let result_2 = vec![
            vec![15, 13, 2, 5],
            vec![14, 3, 4, 1],
            vec![12, 6, 8, 9],
            vec![16, 7, 10, 11],
        ];
        Solution::rotate(&mut matrix_2);
        assert_eq!(result_2, matrix_2);
    }
}
