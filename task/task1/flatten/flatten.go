package flatten

//扁平化多级双向链表：
//多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。
//这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。
//给定位于列表第一级的头节点，请扁平化列表，即将这样的多级双向链表展平成普通的双向链表，使所有结点出现在单级双链表中。
//可以定义一个结构体来表示链表节点，包含 val、prev、next 和 child 指针，然后使用递归的方法来扁平化链表，先处理当前节点的子链表，再将子链表插入到当前节点和下一个节点之间。
/*
示例 1：
输入：head = [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
输出：[1,2,3,7,8,11,12,9,10,4,5,6]
示例 2：
输入：head = [1,2,null,3]
输出：[1,3,2]
示例 3：
输入：head = []
输出：[]
说明：输入中可能存在空列表。
*/

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	flattenDFS(root)
	return root
}

func flattenDFS(curr *Node) (tail *Node) {
	var prev *Node
	for curr != nil {
		next := curr.Next // 保存下一个节点
		if curr.Child != nil {
			// 递归处理子链表，得到其尾节点
			childTail := flattenDFS(curr.Child)
			// 将子链表插入到当前节点之后
			curr.Next = curr.Child
			curr.Child.Prev = curr
			// 将子链表的尾节点与原next节点连接
			if next != nil {
				next.Prev = childTail
			}
			childTail.Next = next
			// 清除子节点指针
			curr.Child = nil
		}
		prev = curr // 更新prev为当前节点
		curr = next // 处理下一个节点
	}
	return prev
}
