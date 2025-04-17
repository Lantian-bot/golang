package main

import (
	"basic/task1/singleNumber"
	"fmt"
)

func main() {
	//	示例 1 ：
	var nums1 = [3]int{2, 2, 1}
	fmt.Println(s1)
	// 示例 2 ：
	var nums2 = [5]int{4, 1, 2, 1, 2}
	s2 := singleNumber.SingleNumber(nums2[:])
	fmt.Println(s2)
	// 示例 3 ：
	var nums3 = [1]int{1}
	s3 := singleNumber.SingleNumber(nums3[:])
	fmt.Println(s3)
}
