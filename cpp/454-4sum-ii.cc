/**
 * @file 454-4sum-ii.cc
 * @author krizsx (cb.kris.bj@gmail.com)
 * @brief https://leetcode-cn.com/problems/4sum-ii/
 * @version 0.1
 * @date 2020-11-27 22:05:00
 * 
 * @copyright Copyright (c) 2020
 * 
 */
#include <iostream>
#include <unordered_map>
#include <vector>
using namespace std;

class Solution {
   public:
    int fourSumCount(vector<int>& A, vector<int>& B, vector<int>& C, vector<int>& D) {
        unordered_map<int, int> m;
        for (auto i : A) {
            for (auto j : B) {
                m[i + j]++;
            }
        }

        int result = 0;
        for (auto i : C) {
            for (auto j : D) {
                if (m.count(-i - j)) {
                    result += m[-i - j];
                }
            }
        }
        return result;
    }
};

int main() {
}