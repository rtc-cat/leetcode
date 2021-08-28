use super::linked_list::*;
#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        // 改为变量
        let (mut l1, mut l2) = (l1, l2);
        let mut head = Some(Box::new(ListNode::new(0)));
        let mut current = head.as_mut();
        let mut carry = 0;
        while l1.is_some() || l2.is_some() {
            let mut sum = carry;
            if let Some(v) = l1 {
                sum += v.val;
                l1 = v.next;
            }
            if let Some(v) = l2 {
                sum += v.val;
                l2 = v.next;
            }
            carry = sum / 10;
            if let Some(cur) = current {
                cur.next = Some(Box::new(ListNode::new(sum % 10)));
                current = cur.next.as_mut();
            }
        }
        if carry > 0 {
            current.unwrap().next = Some(Box::new(ListNode::new(carry)));
        }
        head.unwrap().next
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_solution() {
        assert_eq!(
            Solution::add_two_numbers(to_list(vec![2, 4, 3]), to_list(vec![5, 6, 4])),
            to_list(vec![7, 0, 8])
        );

        assert_eq!(
            Solution::add_two_numbers(to_list(vec![9, 9, 9, 9]), to_list(vec![9, 9, 9, 9, 9, 9])),
            to_list(vec![8, 9, 9, 9, 0, 0, 1])
        );

        assert_eq!(
            Solution::add_two_numbers(to_list(vec![0]), to_list(vec![0])),
            to_list(vec![0])
        )
    }
}
