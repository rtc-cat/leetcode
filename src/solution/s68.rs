use std::vec;

#[allow(unused_macros)]
macro_rules! vec_of_strings {
    // match a list of expressions separated by comma:
    ($($str:expr),*) => ({
        // create a Vec with this list of expressions,
        // calling String::from on each:
        vec![$(String::from($str),)*] as Vec<String>
    });
}

#[allow(dead_code)]
struct Solution {}

// 给定一个单词数组和一个长度 maxWidth，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。

// 你应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。

// 要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。

// 文本的最后一行应为左对齐，且单词之间不插入额外的空格。

// 说明:

// 单词是指由非空格字符组成的字符序列。
// 每个单词的长度大于 0，小于等于 maxWidth。
// 输入单词数组 words 至少包含一个单词。

impl Solution {
    #[allow(dead_code)]
    pub fn full_justify(words: Vec<String>, max_width: i32) -> Vec<String> {
        // 首先确定每一行需要多少个单词
        let mut multi_line: Vec<Vec<String>> = vec![vec![]; words.len()]; // 每一行的单词
        let mut index = 0; // 当前行 索引
        let mut width = 0; // 当前行字符宽度
        for w in words.iter() {
            let line = &mut multi_line[index];
            width += w.len();
            line.push(w.to_owned());
            // 最终的最小宽度, 应该是每个单词的长度, 加上单词间的空格
            let line_min_width = width + line.len() - 1;
            // 如果能够装得下, 则继续下一个单词, 还在这一行继续加
            if line_min_width <= max_width as usize {
                continue;
            }
            // 如果装不下的话, 就需要下一行来装
            index += 1;
            // 把这一行的单词弹出
            line.pop();
            // 宽度变为这个单词
            width = w.len();
            // 下一行存入这个单词
            let next_line = &mut multi_line[index];
            next_line.push(w.to_owned());
        }
        // 记录最后一行的index
        let last_index = index;
        // 上面这个循环结束后, 每一行多少个单词已经确定, 接下来需要确定添加多少空格
        // 添加空格, 应该按照, 除了最后一个单词外的其他单词 依次往后添加空格
        // println!("{:?}", multi_line);
        // 最后一行要左对齐
        let mut result = vec![];
        for (i, line) in multi_line.iter_mut().enumerate() {
            if line.len() == 0 {
                continue;
            }

            // 先获取一下当前的长度
            let mut width = 0;
            for w in line.iter() {
                width += w.len();
            }
            // 计算需要多少空格
            width = max_width as usize - width;
            // 依次添加空格
            let mut index = 0;
            // 如果是最后一行需要特殊处理,左对齐
            if i == last_index {
                // 每个单词先添加一个空格, 然后直接在最后一个字符串补上所有空格
                while width > 0 {
                    line[index].push(' ');
                    width -= 1;
                    index += 1;
                    // 如果到最后一个
                    if index >= line.len() {
                        while width > 0 {
                            line[index - 1].push(' ');
                            width -= 1;
                        }
                    }
                }
            } else {
                while width > 0 {
                    line[index].push(' ');
                    width -= 1;
                    index += 1;
                    if index >= line.len() - 1 {
                        index = 0;
                    }
                }
            }

            // 把最终的结果组合成一个字符串, 写入结果中
            let mut line_str = String::new();
            for w in line.iter() {
                line_str += w;
            }
            result.push(line_str)
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
            Solution::full_justify(
                vec_of_strings![
                    "This",
                    "is",
                    "an",
                    "example",
                    "of",
                    "text",
                    "justification."
                ],
                16
            ),
            vec!["This    is    an", "example  of text", "justification.  "]
        );
        assert_eq!(
            Solution::full_justify(
                vec_of_strings!["What", "must", "be", "acknowledgment", "shall", "be"],
                16
            ),
            vec!["What   must   be", "acknowledgment  ", "shall be        "]
        );
        assert_eq!(
            Solution::full_justify(
                vec_of_strings![
                    "Science",
                    "is",
                    "what",
                    "we",
                    "understand",
                    "well",
                    "enough",
                    "to",
                    "explain",
                    "to",
                    "a",
                    "computer.",
                    "Art",
                    "is",
                    "everything",
                    "else",
                    "we",
                    "do"
                ],
                20
            ),
            vec![
                "Science  is  what we",
                "understand      well",
                "enough to explain to",
                "a  computer.  Art is",
                "everything  else  we",
                "do                  "
            ]
        );
    }
}
