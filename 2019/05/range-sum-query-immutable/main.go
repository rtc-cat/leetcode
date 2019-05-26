package main

// 给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。

// 示例：

// 给定 nums = [-2, 0, 3, -5, 2, -1]，求和函数为 sumRange()

// sumRange(0, 2) -> 1
// sumRange(2, 5) -> -1
// sumRange(0, 5) -> -3
// 说明:

// 你可以假设数组不可变。
// 会多次调用 sumRange 方法。

func main() {

}

// NumArray 数组
type NumArray struct {
	nums []int
}

// Constructor 构造器
func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{nums: nums}
	}
	// sum是从第一个数字累加到第i个位置的和
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
	}
	return NumArray{nums: sum}
}

// SumRange 求和函数 (0,2)=sum[2]
// (2,5)=sum[5]-sum[1]
func (na *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return na.nums[j]
	}
	return na.nums[j] - na.nums[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
