#!/usr/bin/env zsh
PROBLEM=$1
echo "运行第${PROBLEM}题测试用例"
cargo test s${PROBLEM} -- --color always --nocapture
