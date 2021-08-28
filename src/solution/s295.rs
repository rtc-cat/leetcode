use std::collections::BinaryHeap; // 默认是最大堆

// 官方库 最大堆,最小堆 https://doc.rust-lang.org/std/collections/struct.BinaryHeap.html
#[allow(dead_code)]
struct MedianFinder {
    cnt: usize,
    min_hp: BinaryHeap<i32>,
    max_hp: BinaryHeap<i32>,
}
/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MedianFinder {
    /** initialize your data structure here. */
    #[allow(dead_code)]
    fn new() -> Self {
        MedianFinder {
            cnt: 0,
            max_hp: BinaryHeap::new(),
            min_hp: BinaryHeap::new(),
        }
    }

    #[allow(dead_code)]
    fn add_num(&mut self, num: i32) {
        self.cnt += 1;
        self.max_hp.push(-num);
        self.min_hp.push(-self.max_hp.pop().unwrap());
        if self.cnt & 1 == 1 {
            self.max_hp.push(-self.min_hp.pop().unwrap());
        }
    }

    #[allow(dead_code)]
    fn find_median(&self) -> f64 {
        if self.cnt == 0 {
            return 0.0;
        }
        if self.cnt & 1 == 1 {
            -self.max_hp.peek().unwrap() as f64
        } else {
            (self.min_hp.peek().unwrap() - self.max_hp.peek().unwrap()) as f64 / 2.0
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        let mut obj = MedianFinder::new();
        assert_eq!(obj.find_median(), 0.0);

        obj.add_num(1);
        assert_eq!(obj.find_median(), 1.0);

        obj.add_num(2);
        assert_eq!(obj.find_median(), 1.5);

        obj.add_num(3);
        assert_eq!(obj.find_median(), 2.0);
    }
}
