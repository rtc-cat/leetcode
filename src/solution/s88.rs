#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn merge(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
        let mut i = (nums1.len() - 1) as i32;
        let (mut i1, mut i2) = ((m - 1), (n - 1));
        while i2 >= 0 {
            while i1 >= 0 && nums1[i1 as usize] > nums2[i2 as usize] {
                nums1.swap(i as usize, i1 as usize);
                i1 -= 1;
                i -= 1;
            }
            nums1[i as usize] = nums2[i2 as usize];
            i2 -= 1;
            i -= 1;
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_solution() {
        let mut nums1 = vec![1, 2, 3, 0, 0, 0];
        Solution::merge(&mut nums1, 3, &mut vec![4, 5, 6], 3);
        println!("{:?}", nums1);
    }
}
