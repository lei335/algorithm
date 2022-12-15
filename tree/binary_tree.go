package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树 前序遍历 递归
func preorderTraversalRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	// 先访问根节点，再前序遍历左子树，再前序遍历右子树
	fmt.Println(root.Val)
	preorderTraversalRecursion(root.Left)
	preorderTraversalRecursion(root.Right)
}

// 二叉树 前序遍历 非递归
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			// 先访问根节点
			result = append(result, root.Val)
			stack = append(stack, root)
			// 再前序访问左子树
			root = root.Left
		}
		// pop 回退访问右子树
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

// 二叉树 中序遍历 递归
func inorderTraversalRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversalRecursion(root.Left)
	fmt.Println(root.Val)
	inorderTraversalRecursion(root.Right)
}

// 二叉树 中序遍历 非递归
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(res) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left // 一直向左
		}
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 访问根节点
		res = append(res, node.Val)
		// 访问右子树
		root = node.Right
	}
	return res
}

// 二叉树 后序遍历 递归
func postorderTraversalRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	postorderTraversalRecursion(root.Left)
	postorderTraversalRecursion(root.Right)
	fmt.Println(root.Val)
}

// 二叉树 后序遍历 非递归
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left // 一直向左访问
		}
		// 先记录，先不弹出
		node := stack[len(stack)-1]
		// 根节点必须在右节点弹出之后，再弹出
		if node.Right == nil || node.Right == lastVisit {
			// 弹出
			stack = stack[:len(stack)-1]
			res = append(res, node.Val)
			// 标记当前这个节点已经访问过
			lastVisit = node
		} else {
			// 先向右访问
			root = node.Right
		}
	}
	return res
}