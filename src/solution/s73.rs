#[allow(dead_code)]
struct Solution {}

struct Point(usize, usize);

impl Solution {
    #[allow(dead_code)]
    pub fn set_zeroes(matrix: &mut Vec<Vec<i32>>) {
        let mut points = vec![];
        for i in 0..matrix.len() {
            for j in 0..matrix[i].len() {
                if matrix[i][j] == 0 {
                    points.push(Point(i, j));
                }
            }
        }

        for point in points {
            for i in 0..matrix[point.0].len() {
                matrix[point.0][i] = 0;
            }
            for i in 0..matrix.len() {
                matrix[i][point.1] = 0;
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        let mut matrix_1 = vec![vec![1, 1, 1], vec![1, 0, 1], vec![1, 1, 1]];
        let result_1 = vec![vec![1, 0, 1], vec![0, 0, 0], vec![1, 0, 1]];
        Solution::set_zeroes(&mut matrix_1);
        assert_eq!(result_1, matrix_1);

        let mut matrix_2 = vec![vec![0, 1, 2, 0], vec![3, 4, 5, 2], vec![1, 3, 1, 5]];
        let result_2 = vec![vec![0, 0, 0, 0], vec![0, 4, 5, 0], vec![0, 3, 1, 0]];
        Solution::set_zeroes(&mut matrix_2);
        assert_eq!(result_2, matrix_2);
    }
}
