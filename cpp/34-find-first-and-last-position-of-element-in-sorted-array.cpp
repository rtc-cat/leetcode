// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

#include <iostream>
#include <vector>
using namespace std;

class Solution
{
public:
    vector<int> searchRange(vector<int> &nums, int target)
    {
        // 左边开始
        bool exists = true;
        int left = 0;
        int right = nums.size() - 1;
        for (;;)
        {
            if (left > right)
            {
                exists = false;
                break;
            }
            if (nums[left] == target && nums[right] == target)
            {
                break;
            }
            if (nums[left] != target)
            {
                left++;
            };
            if (nums[right] != target)
            {
                right--;
            }
        }

        return exists ? vector<int>{left, right} : vector<int>{-1, -1};
    }
};

int main()
{
    Solution solution;
    vector<int> nums{5, 7, 7, 8, 8, 10};
    auto result = solution.searchRange(nums, 8);
    for (auto &n : result)
    {
        std::cout << n << std::endl;
    }
}