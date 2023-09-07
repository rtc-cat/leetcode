#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        // 排序?
        let mut result = std::collections::HashMap::<String, Vec<String>>::new();
        for i in 0..strs.len() {
            let mut sorted: Vec<_> = strs[i].chars().collect();
            sorted.sort();
            let s: String = sorted.into_iter().collect();
            result
                .entry(s)
                .and_modify(|v| v.push(strs[i].clone()))
                .or_insert(vec![strs[i].clone()]);
        }
        result.values().cloned().collect()
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::vec_of_strings;

    #[test]
    fn test_solution() {
        assert_eq!(
            vec![
                vec_of_strings!["ate", "eat", "tea"],
                vec_of_strings!["nat", "tan"],
                vec_of_strings!["bat"],
            ],
            Solution::group_anagrams(vec_of_strings!["eat", "tea", "tan", "ate", "nat", "bat"])
        );
        assert_eq!(
            vec![vec_of_strings![""]],
            Solution::group_anagrams(vec_of_strings![""])
        );
        assert_eq!(
            vec![vec_of_strings!["a"]],
            Solution::group_anagrams(vec_of_strings!["a"])
        );
    }
}
