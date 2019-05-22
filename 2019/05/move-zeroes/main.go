package main

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

// 示例:

// 输入: [0,1,0,3,12]
// 输出: [1,3,12,0,0]
// 说明:

// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。

func main() {

}

func moveZeroes(nums []int) {
	// 两个指针,一个是零指针,一个是非零指针
	zero := 0
	non := 0
	for ; non < len(nums); non = non + 1 {
		// 如果非零指针指着的也是零,那么就开始下一次循环
		if nums[non] == 0 {
			continue
		}
		// 如果不是的话,就交换位置
		// 这样可以把所有非零的数都移动到前面
		nums[zero] = nums[non]
		zero++
	}
	for i := zero; i < len(nums); i++ {
		nums[i] = 0
	}
}
