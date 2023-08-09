#!/usr/bin/env zsh
PROBLEM=$1
echo "运行第${PROBLEM}题测试用例"
cargo test solution::s${PROBLEM}::tests::test_solution -- --color always --nocapture --exact
