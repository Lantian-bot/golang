package mergeTwoLists

//合并2个有序链表
//将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//可以定义一个函数，接收两个链表的头节点作为参数，在函数内部使用双指针法，通过比较两个链表节点的值，将较小值的节点添加到新链表中，
//直到其中一个链表为空，然后将另一个链表剩余的节点添加到新链表中。
/*
示例 1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
示例 2：
输入：l1 = [], l2 = []
输出：[]
示例 3：
输入：l1 = [], l2 = [0]
输出：[0]
*/

// 链表的结构：每个节点有一个Val字段和一个Next指针
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // 创建哑节点作为新链表的起点
	current := dummy     // 当前指针用于构建新链表

	// 双指针遍历两个链表
	for l1 != nil && l2 != nil {
		// 比较l1和l2的当前节点值，将较小的那个连接到current的Next，并移动相应的指针
		if l1.Val < l2.Val {
			current.Next = l1 // 连接较小值的节点
			l1 = l1.Next      // 移动l1指针
		} else {
			current.Next = l2 // 连接较小值的节点
			l2 = l2.Next      // 移动l2指针
		}
		current = current.Next // 移动当前指针
	}

	// 处理剩余节点（最多只有一个链表非空）
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next // 返回哑节点的下一个节点（真正的新链表头）
}

// 将切片转换为链表，方便测试时快速构造输入
func SliceToList(s []int) *ListNode {
	dummy := &ListNode{} // 哑节点简化操作
	current := dummy
	for _, val := range s {
		current.Next = &ListNode{Val: val} // 创建新节点并链接
		current = current.Next
	}
	return dummy.Next // 返回真正的头节点
}

// 将链表转换为切片，方便验证结果：
func ListToSlice(head *ListNode) []int {
	var slice []int
	for head != nil {
		slice = append(slice, head.Val)
		head = head.Next
	}
	return slice
}
