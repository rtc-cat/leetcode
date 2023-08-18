#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn can_complete_circuit(gas: Vec<i32>, cost: Vec<i32>) -> i32 {
        let (mut i, n) = (0, gas.len());
        while i < n {
            let (mut sum_gas, mut sum_cost, mut cnt) = (0, 0, 0);
            while cnt < n {
                let index = (i + cnt) % n;
                sum_gas += gas[index];
                sum_cost += cost[index];
                if sum_cost > sum_gas {
                    break;
                }
                cnt += 1;
            }
            if cnt == n {
                return i as i32;
            } else {
                i += cnt + 1;
            }
        }
        -1
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(
            3,
            Solution::can_complete_circuit(vec![1, 2, 3, 4, 5], vec![3, 4, 5, 1, 2])
        );
        assert_eq!(
            -1,
            Solution::can_complete_circuit(vec![2, 3, 4], vec![3, 4, 3])
        );
    }
}
