package mySqrt

// X的平方根
//实现 int sqrt(int x) 函数。计算并返回 x 的平方根，其中 x 是非负整数。由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
//可以使用二分查找法来解决，定义左右边界 left 和 right，然后通过 while 循环不断更新中间值 mid，直到找到满足条件的平方根或者确定不存在精确的平方根。
/*
示例 1：
输入：x = 4
输出：2
示例 2：
输入：x = 8
输出：2
解释：8
*/

func MySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left, right := 1, x
	ans := 0
	for left <= right {
		mid := left + (right-left)/2
		if mid <= x/mid {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
