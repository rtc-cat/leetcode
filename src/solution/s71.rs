#[allow(dead_code)]
struct Solution {}

impl Solution {
    #[allow(dead_code)]
    pub fn simplify_path(path: String) -> String {
        let mut result = Vec::new();
        let splited: Vec<_> = path.split("/").collect();

        result.push("");
        for i in 0..splited.len() {
            let dir = splited[i];
            if dir.len() == 0 {
                continue;
            }
            if dir == "." {
                continue;
            }
            if dir == ".." {
                if result.len() > 1 {
                    result.pop();
                }
            } else {
                result.push(dir);
            }
        }
        if result.len() == 1 {
            return "/".to_string();
        }
        result.join("/")
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        assert_eq!(
            String::from("/home"),
            Solution::simplify_path(String::from("/home/"))
        );
        assert_eq!(
            String::from("/"),
            Solution::simplify_path(String::from("/../"))
        );
        assert_eq!(
            String::from("/home/foo"),
            Solution::simplify_path(String::from("/home//foo/"))
        );
        assert_eq!(
            String::from("/c"),
            Solution::simplify_path(String::from("/a/./b/../../c/"))
        );
    }
}
