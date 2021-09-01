#[allow(dead_code)]
struct Solution {}

// 这里有 n 个航班，它们分别从 1 到 n 进行编号。
// 有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi]
// 意味着在从 firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。
// 请你返回一个长度为 n 的数组 answer，其中 answer[i] 是航班 i 上预订的座位总数。
impl Solution {
    #[allow(dead_code)]
    pub fn corp_flight_bookings(bookings: Vec<Vec<i32>>, n: i32) -> Vec<i32> {
        let mut result = vec![0; n as usize];
        for booking in bookings.iter() {
            let (first, last, seats) = (booking[0], booking[1], booking[2]);
            for i in first - 1..last {
                result[i as usize] += seats;
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
        assert_eq!(
            Solution::corp_flight_bookings(vec![vec![1, 2, 10], vec![2, 3, 20], vec![2, 5, 25]], 5),
            vec![10, 55, 45, 25, 25]
        );
        assert_eq!(
            Solution::corp_flight_bookings(vec![vec![1, 2, 10], vec![2, 2, 15]], 2),
            vec![10, 25]
        );
    }
}
