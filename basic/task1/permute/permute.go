package permute

// 全排列
//给定一个不含重复数字的数组 nums ，返回其所有可能的全排列。 可以使用回溯算法，
//定义一个函数来进行递归操作，在函数中通过交换数组元素的位置来生成不同的排列，使用 for 循环遍历数组，每次选择一个元素作为当前排列的第一个元素，然后递归调用函数处理剩余的元素。
/*
示例 1：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：
输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：
输入：nums = [1]
输出：[[1]]
*/

func Permute(nums []int) [][]int {
	var result [][]int
	backtrack(0, nums, &result)
	return result
}

func backtrack(k int, nums []int, result *[][]int) {
	if k == len(nums) {
		// 复制当前排列到结果中
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*result = append(*result, tmp)
		return
	}
	// 遍历所有可能的交换位置
	for i := k; i < len(nums); i++ {
		// 交换当前元素和第i个元素
		nums[k], nums[i] = nums[i], nums[k]
		// 递归生成后续位置的排列
		backtrack(k+1, nums, result)
		// 恢复交换，回溯
		nums[k], nums[i] = nums[i], nums[k]
	}
}
