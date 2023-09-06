#[allow(dead_code)]
struct Solution {}

#[derive(PartialEq)]
enum Direction {
    Right,
    Down,
    Left,
    Up,
}

impl Solution {
    #[allow(dead_code)]
    pub fn spiral_order(matrix: Vec<Vec<i32>>) -> Vec<i32> {
        if matrix.len() == 0 {
            return vec![];
        }
        let mut direction = Direction::Right;
        let mut result = Vec::<i32>::new();
        let (mut min_row, mut max_row) = (0 as i32, (matrix.len() - 1) as i32);
        let (mut min_col, mut max_col) = (0 as i32, (matrix[0].len() - 1) as i32);
        let mut row: i32 = 0;
        let mut col: i32 = 0;
        loop {
            if min_row > max_row || min_col > max_col {
                break;
            }
            result.push(matrix[row as usize][col as usize]);
            match direction {
                Direction::Right => {
                    if col == max_col {
                        direction = Direction::Down;
                        row += 1;
                        min_row += 1;
                    } else {
                        col += 1;
                    }
                }
                Direction::Down => {
                    if row == max_row {
                        direction = Direction::Left;
                        col -= 1;
                        max_col -= 1;
                    } else {
                        row += 1;
                    }
                }
                Direction::Left => {
                    if col == min_col {
                        direction = Direction::Up;
                        row -= 1;
                        max_row -= 1;
                    } else {
                        col -= 1;
                    }
                }
                Direction::Up => {
                    if row == min_row {
                        direction = Direction::Right;
                        col += 1;
                        min_col += 1;
                    } else {
                        row -= 1;
                    }
                }
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
        assert_eq!(vec![3, 2], Solution::spiral_order(vec![vec![3], vec![2]]));
        assert_eq!(
            vec![1, 2, 3, 6, 9, 8, 7, 4, 5],
            Solution::spiral_order(vec![vec![1, 2, 3], vec![4, 5, 6], vec![7, 8, 9]])
        );
        assert_eq!(
            vec![1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7],
            Solution::spiral_order(vec![
                vec![1, 2, 3, 4],
                vec![5, 6, 7, 8],
                vec![9, 10, 11, 12]
            ])
        )
    }
}
