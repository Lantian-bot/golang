package main

import (
	"basic/task1/merge"
	"basic/task1/mergeTwoLists"
	"basic/task1/mySqrt"
	"basic/task1/permute"
	"basic/task1/removeDuplicates"
	"basic/task1/reverseString"
	"basic/task1/rob"
	"basic/task1/singleNumber"
	"fmt"
)

func main() {
	// 1.只出现一次的数字
	//	示例 1 ：
	var nums1 = [3]int{2, 2, 1}
	s1 := singleNumber.SingleNumber(nums1[:]) // 左侧变量是数组，右侧函数入参是切片，不同类型所以使用该语法转换
	fmt.Println(s1)
	// 示例 2 ：
	var nums2 = [5]int{4, 1, 2, 1, 2}
	s2 := singleNumber.SingleNumber(nums2[:])
	fmt.Println(s2)
	// 示例 3 ：
	var nums3 = [1]int{1}
	s3 := singleNumber.SingleNumber(nums3[:])
	fmt.Println(s3)
	// 2.打家劫舍
	var nums4 = [5]int{2, 11, 9, 3, 1}
	s4 := rob.Rob(nums4[:])
	fmt.Println(s4)
	// 3.合并2个有序链表
	l1 := mergeTwoLists.SliceToList([]int{1, 2, 4})
	l2 := mergeTwoLists.SliceToList([]int{1, 3, 4})
	l3 := mergeTwoLists.MergeTwoLists(l1, l2)
	fmt.Println(mergeTwoLists.ListToSlice(l3))
	// 4.全排列
	var nums5 = [3]int{1, 2, 3}
	s5 := permute.Permute(nums5[:])
	fmt.Println(s5)
	// 5.反转字符串
	var s6 = []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(string(s6))
	s := reverseString.ReverseString(s6)
	fmt.Println(string(s))
	// 6.X的平方根
	x := mySqrt.MySqrt(11)
	fmt.Println(x)
	// 7.删除有序数组中的重复项
	var nums7 = [10]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	s7 := removeDuplicates.RemoveDuplicates(nums7[:])
	fmt.Println(s7)
	// 8.合并区间
	var intervals = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	s8 := merge.Merge(intervals[:])
	fmt.Println(s8)
}
