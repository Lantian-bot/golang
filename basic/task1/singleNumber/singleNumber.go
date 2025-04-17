package singleNumber

//只出现一次的数字
//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
/*
示例 1 ：
输入：nums = [2,2,1]
输出：1
示例 2 ：
输入：nums = [4,1,2,1,2]
输出：4
示例 3 ：
输入：nums = [1]
输出：1
*/

func SingleNumber(nums []int) int {
	countMap := make(map[int]int)
	// 遍历数组，统计每个元素出现的次数，并存入 map
	for _, num := range nums {
		if _, exists := countMap[num]; exists {
			//如果num已经在map中存在，就将对应的值+1
			countMap[num]++
		} else {
			countMap[num] = 1
		}
	}
	// 遍历map，找到出现次数为1的元素
	for num, count := range countMap {
		if count == 1 {
			return num
		}
	}
	return 0
}
