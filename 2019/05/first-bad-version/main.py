# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool
# def isBadVersion(version):


def isBadVersion(version):
    return


class Solution:
    def firstBadVersion(self, n):
        """
        :type n: int
        :rtype: int
        [1:n)区间
        """
        # 二分查找
        low = 1
        high = n
        mid = low + (high - low) // 2
        while low < high:
            if (isBadVersion(mid)):
                high = mid
            else:
                low = mid + 1
            mid = low + (high - low) // 2
        return low
