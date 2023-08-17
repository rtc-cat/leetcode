#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn jump(nums: Vec<i32>) -> i32 {
        let mut cur = 0 as usize;
        let mut steps = 0;

        loop {
            if cur >= nums.len() - 1 {
                break;
            }
            if nums[cur] == 0 {
                break;
            }
            let mut max = cur;
            let mut max_v = 0;
            for i in 1..=nums[cur] {
                let next = cur + i as usize;
                if next >= nums.len() - 1 {
                    return steps + 1;
                }
                let v = next + nums[next] as usize;
                if v > max_v {
                    max = next;
                    max_v = v;
                }
            }
            cur = max;
            steps += 1;
        }
        steps
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(2, Solution::jump(vec![2, 3, 1, 1, 4]));
        assert_eq!(2, Solution::jump(vec![2, 3, 0, 1, 4]));
        assert_eq!(3, Solution::jump(vec![1, 2, 1, 1, 1]));
    }
}
