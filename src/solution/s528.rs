use rand::Rng;
#[allow(dead_code)]
struct Solution {
    prefix_sum: Vec<i32>,
}

impl Solution {
    #[allow(dead_code)]
    fn new(mut w: Vec<i32>) -> Self {
        for i in 1..w.len() {
            w[i] += w[i - 1];
        }

        Self { prefix_sum: w }
    }

    #[allow(dead_code)]
    fn pick_index(&self) -> i32 {
        let x = rand::thread_rng().gen_range(1..self.prefix_sum.last().unwrap() + 1);
        match self.prefix_sum.binary_search(&x) {
            Ok(i) => i as i32,
            Err(i) => i as i32,
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_solution() {
        let s = Solution::new(vec![1]);
        assert_eq!(s.pick_index(), 0);

        // 1 -> sum(vec)
        // left[i] = left[i-1] + w[i]
        // right = left + self - 1
        // [1,3] [4,5] [6,7]
        let s = Solution::new(vec![3, 2, 2]);

        for _ in 0..12 {
            println!("{}", s.pick_index());
        }
        // assert_eq!(s.pick_index(), 1);
        // assert_eq!(s.pick_index(), 1);
        // assert_eq!(s.pick_index(), 1);
        // assert_eq!(s.pick_index(), 0);
    }
}
