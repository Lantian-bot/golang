package reverseString

// 5.反转字符串
//编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
//不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
//可以使用 for 循环和两个指针，一个指向字符串的开头，一个指向字符串的结尾，然后交换两个指针所指向的字符，直到两个指针相遇。
/*
示例 1：
输入：s = ["h","e","l","l","o"]
输出：["o","l","l","e","h"]
示例 2：
输入：s = ["H","a","n","n","a","h"]
输出：["h","a","n","n","a","H"]
*/

func ReverseString(s []byte) []byte {
	left, right := 0, len(s)-1
	// 双指针向中间移动，直到相遇
	for left < right {
		// 交换左右指针的字符
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return s
}
