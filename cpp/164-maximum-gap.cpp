/**
 * @file 164-maximum-gap.cc
 * @author krizsx (cb.kris.bj@gmail.com)
 * @brief https://leetcode-cn.com/problems/maximum-gap/
 * @version 0.1
 * @date 2020-11-27 21:51:32
 * 
 * @copyright Copyright (c) 2020
 * 
 */
#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;

class Solution {
   public:
    int maximumGap(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        int max = 0;
        for (int i = 1; i < nums.size(); i++) {
            int n = nums[i] - nums[i - 1];
            n > max ? max = n : 0;
        }
        return max;
    }
};

int main() {
}