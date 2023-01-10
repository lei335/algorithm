package list

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 给定一个排序链表，删除重复元素，使得每个元素只出现一次
func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil {
		for current.Next != nil && current.Val == current.Next.Val {
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
	if head == nil {
		return head
	}
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	// 有可能删除原本的头节点，所以用dummy节点来遍历
	head = dummy

	var rmVal int
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			// 记录被删掉的值，用于后续节点比对
			rmVal = head.Next.Val
			for head.Next != nil && head.Next.Val == rmVal {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return dummy.Next
}

// 反转一个单链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}

// 反转从位置m到n的链表，请使用一趟扫描完成反转
// 思路：先遍历到m处，再开始反转，到n处结束反转，再拼接后续，注意指针处理
func reverseBetween(head *ListNode, m, n int) *ListNode {
	// 思路：输入1->2->3->4->5->NULL, m=2,n=4
	if head == nil {
		return head
	}
	// 以防头部变化，所以要使用dummy node
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	head = dummy
	// 最开始：0->1->2->3->4->5->nil
	var pre *ListNode
	i := 0
	for i < m {
		pre = head
		head = head.Next
		i++
	}
	// 遍历之后：0->1(pre)->2(head)->3->4->5->nil
	j := i // 此时i=m
	var next *ListNode
	var mid = head
	for head != nil && j <= n {
		// 第一次循环 0->1(pre) nil<-2(next) 3(head)->4->5->nil
		temp := head.Next
		head.Next = next
		next = head
		head = temp
		j++
	}
	// 循环执行到n
	// 循环结束：0(dummy)->1(pre) nil<-2(mid)<-3<-4(next) 5(head)->nil
	pre.Next = next
	mid.Next = head
	return dummy.Next
}

// 上述算法的自写
func reverseBetween2(head *ListNode, m, n int) *ListNode {
	if head == nil {
		return head
	}
	// 以防翻转头结点，使用dummy node
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	head = dummy
	var pre *ListNode
	for i := 0; i < m && head != nil; i++ {
		pre = head
		head = head.Next
	}
	j := m
	// 开始翻转
	var next *ListNode
	revStart := head
	for head != nil && j <= n {
		temp := head.Next
		head.Next = next
		next = head
		head = temp
		j++
	}
	// 开始链接翻转的部分和未翻转的部分
	pre.Next = next
	revStart.Next = head
	return dummy.Next
}

// 将两个升序链表合并为一个新的升序链表并返回，新链表是通过拼接给定的两个链表的所有节点组成的
// 思路：通过dummy node链表，连接各个元素
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	head := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	for l1 != nil {
		head.Next = l1
		head = head.Next
		l1 = l1.Next
	}
	for l2 != nil {
		head.Next = l2
		head = head.Next
		l2 = l2.Next
	}
	return dummy.Next
}

// 给定一个链表和一个特定值x，对链表进行分隔，使得所有小于x的节点都在大于或等于x的节点之前
// 思路：将大于x的节点，放到另外一个链表，最后连接两个链表
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	headDummy := &ListNode{Val: 0}
	tailDummy := &ListNode{Val: 0}
	tail := tailDummy
	headDummy.Next = head
	head = headDummy

	for head.Next != nil {
		if head.Next.Val < x {
			head = head.Next
		} else {
			// 移除大于等于x的节点，放到另外一个链表
			t := head.Next
			head.Next = head.Next.Next
			tail.Next = t
			tail = tail.Next
		}
	}
	tail.Next = nil // 注意：这里的设置很重要，不然下面可能会报错，因为有可能tail.Next没有初始化
	head.Next = tailDummy.Next
	return headDummy.Next
}

// 哑巴节点dummy node使用的场景：当头节点不确定的时候，就使用哑巴节点

// 在O(nlogn)时间复杂度和常数级空间复杂度下，对链表进行排序
// 思路：归并排序，找中点和合并操作；和分治法思想相同（返回条件，分治，合并结果）
func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}
func findMiddle(head *ListNode) *ListNode {
	// 1->2->3->4->5
	slow := head
	fast := head.Next
	// 快指针会先为nil
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
func mergeSort(head *ListNode) *ListNode {
	// 如果只有一个节点，就直接返回这个节点
	if head == nil || head.Next == nil {
		return head
	}
	// find middle
	middle := findMiddle(head)
	// 断开中间节点
	tail := middle.Next
	middle.Next = nil
	// 归并排序
	left := mergeSort(head)
	right := mergeSort(tail)
	result := mergeTwoLists(left, right)
	return result
}

// 上述注意点:
// 快慢指针判断fast及fast.Next是否为nil值
// 递归mergeSort时，需要断开中间节点
// 递归返回条件为head为nil或者head.Next为nil

// 给定一个单链表L：L1->L2->...-> Ln-1 -> Ln，将其重新排列后变为：L1 -> Ln -> L2 -> Ln-1 -> L3 ->Ln-2 -> L4 -> Ln-3 ...
// 思路：找到中点断开，翻转后面部分，然后合并前后两个链表
func reorderList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	middle := findMiddle(head)
	tail := reverseList(middle.Next)
	middle.Next = nil
	head = mergeTwoLists2(head, tail)
	return head
}
func mergeTwoLists2(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	head := dummy
	toggle := true
	for l1 != nil && l2 != nil {
		if toggle {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
		toggle = !toggle
	}
	for l1 != nil {
		head.Next = l1
		l1 = l1.Next
		head = head.Next
	}
	for l2 != nil {
		head.Next = l2
		l2 = l2.Next
		head = head.Next
	}
	return dummy.Next
}

// 给定一个链表，判断链表中是否有环
// 思路：快慢指针，最终快慢指针相同则有环，因为终会相遇
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		// 比较指针是否相等，不要比较val
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}

// 给定一个链表，返回链表开始入环的第一个节点，如果链表无环，则返回null
// 思路：快慢指针，快慢指针相遇之后，快指针经过的节点数（快指针每次走两个）是慢指针节点数（慢指针每次走一个）的两倍
// 则相遇点到入环第一点的距离刚好等于头节点到入环第一点的距离
// 相遇之后，快指针回到头节点往前走，慢指针从相遇点往前走，两个指针步调一致，则再次相遇点既为入环第一点
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	fast := head.Next
	slow := head
	for fast != nil && slow != nil {
		if fast == slow {
			// 快指针从头开始
			fast = head
			slow = slow.Next // 注意这一点
			for slow != fast {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return nil
}

// 坑点：指针比较时直接比较对象，不要用值比较，链表中有可能存在重复值情况
// 坑点：第一次相交后，慢指针需要从下一个节点开始和头指针一起匀速移动

// 上述问题的另外一种方式：fast=head; slow=head
func detectCycle2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			// 指针重新从头开始移动
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}

// 这两种方式不同点在于，一般用fast=head.Next较多，因为这样可以知道中点的上一个节点，可以用来删除等操作
// fast如果初始化为head.Next，则中点在slow.Next
// fast初始化为head，则中点在slow

// 判断一个链表是否为回文链表
// 回文链表：链表正序和逆序输出结果一致
// 思路：一个链表为回文链表，则它关于中轴线是左右对称的，也就是后半段翻转后和前半段的值一样
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	// 找中点
	// fast如果初始化为head.Next, 则中点在slow.Next
	// fast初始化为head,则中点在slow
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 中点为slow.Next
	tail := reverseList(slow.Next)
	slow.Next = nil

	for head != nil && tail != nil {
		if head.Val != tail.Val {
			return false
		}
		head = head.Next
		tail = tail.Next
	}
	return true
}

// 给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。要求返回这个链表的深拷贝
// 浅拷贝：在拷贝一种数据结构的时候（结构体、类、map...），如果拷贝的只是这个数据结构的引用
// 即修改p的值，但浅拷贝的s里的值也发生了变化，这说明s和p实际上都是指向的同一块内存地址，拷贝s到P的时候，其实只是拷贝了一个指针
// 深拷贝：不是只拷贝一个指针，而是要拷贝指针指向的内存中的所有数据。对于map的拷贝，流行做法是通过序列化和反序列化
// 思路：1. hash表存储指针；2. 复制节点跟在原节点后面
func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	// 复制节点，紧挨到后面
	cur := head
	for cur != nil {
		clone := &Node{Val: cur.Val, Next: cur.Next}
		temp := cur.Next
		cur.Next = clone
		cur = temp
	}
	// 处理random指针
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	// 分离两个链表
	cur = head
	cloneHead := cur.Next
	for cur != nil && cur.Next != nil {
		temp := cur.Next
		cur.Next = cur.Next.Next
		cur = temp
	}
	// 原始链表头：head
	// 克隆的链表头：cloneHead
	return cloneHead
}

// 总结：链表需要掌握的一些点
// null/nil异常处理
// dummy node哑巴节点的使用
// 快慢指针
// 插入一个节点到排序链表
// 从一个链表中移除一个节点
// 翻转链表
// 合并两个链表
// 找到链表的中间节点
