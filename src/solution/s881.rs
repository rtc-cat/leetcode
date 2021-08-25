// https://leetcode-cn.com/problems/boats-to-save-people/
#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn num_rescue_boats(people: Vec<i32>, limit: i32) -> i32 {
        // 贪心, 尽量给体重最大的塞更多的同伴
        let mut people = people;
        people.sort();
        let (mut light, mut heavy) = (0, people.len() - 1);
        let mut result = 0;
        while light <= heavy {
            // 如果是 0,0,0,0 那么heavy会一直减,到-1,所以需要加这个判断
            if light == heavy {
                result += 1;
                break;
            }
            if people[light] + people[heavy] > limit {
                heavy -= 1;
            } else {
                light += 1;
                heavy -= 1;
            }
            result += 1;
        }
        result
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_solution() {
        assert_eq!(Solution::num_rescue_boats(vec![1, 2], 3), 1);
        assert_eq!(Solution::num_rescue_boats(vec![3, 2, 2, 1], 3), 3);
        assert_eq!(Solution::num_rescue_boats(vec![3, 5, 3, 4], 5), 4);
    }
}
