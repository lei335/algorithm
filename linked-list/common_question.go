package list

type ListNode struct {
	Val int
	Next *ListNode
}

// 给定一个排序链表，删除重复元素，使得每个元素只出现一次
func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current!=nil{
		for current.Next!=nil && current.Val==current.Next.Val{
			current.Next = current.Next.Next
		}
		current = current.Next
	}
	return head
}

// 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中，没有重复出现的数字
// 链表头节点可能被删除
// 注意点：A->B->C删除B,A.Next=C. 删除用一个dummy node来辅助，从而允许头节点可变。访问X.Next、X.Val，前提要保证X!=nil
func deleteDuplicates2(head *ListNode) *ListNode {
	if head==nil{
		return head
	}
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	// 有可能删除原本的头节点，所以用dummy节点来遍历
	head = dummy

	var rmVal int
	for head.Next!=nil && head.Next.Next!=nil{
		if head.Next.Val==head.Next.Next.Val{
			// 记录被删掉的值，用于后续节点比对
			rmVal = head.Next.Val
			for head.Next!=nil&&head.Next.Val==rmVal{
				head.Next = head.Next.Next
			}
		}else{
			head = head.Next
		}
	}
	return dummy.Next
}

// 反转一个单链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head!=nil{
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}

// 反转从位置m到n的链表，请使用一趟扫描完成反转
// 思路：先遍历到m处，再开始反转，到n处结束反转，再拼接后续，注意指针处理