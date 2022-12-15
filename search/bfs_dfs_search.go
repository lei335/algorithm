package search

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS 深度搜索，从上往下，递归法
// 深度搜索，其实思想和前序遍历是一样的
func dfs(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	dfs(root.Left, res)
	dfs(root.Right, res)
}

// DFS 深度搜索，从下往上，分治法
func divideAndConquer(root *TreeNode) []int {
	res := make([]int, 0)
	// 返回条件：null & leaf
	if root == nil {
		return res
	}
	// 分治（Divide）
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)
	// 合并结果（Conquer）
	res = append(res, root.Val)
	res = append(res, left...)
	res = append(res, right...)
	return res
}

// BFS 层次遍历
func levelOrder(root *TreeNode) [][]int {
	// 通过上一层的长度确定下一层的长度
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		length := len(queue) // 记录每一层的长度
		tmp := make([]int, 0)
		for i := 0; i < length; i++ {
			node := queue[i]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[length:]
		res = append(res, tmp)
	}
	return res
}

// ---------------分治法---------------
// 分治法模板： 1. 递归返回条件；2. 分段处理； 3. 合并结果

// 分治法实现归并排序
func MergeSort(nums []int) []int {
	return mergeSort(nums)
}
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	// 分治法：分为两段
	mid := len(nums) / 2
	left := mergeSort(nums[:mid]) // 返回结果用于合并
	right := mergeSort(nums[mid:])
	// 合并两端数据
	res := merge(left, right)
	return res
}
func merge(left, right []int) (res []int) {
	// 两边数组合并游标
	l := 0
	r := 0
	// 注意不能越界
	for l < len(left) && r < len(right) {
		// 谁小合并谁
		if left[l] > right[r] {
			res = append(res, right[r])
			r++
		} else {
			res = append(res, left[l])
			l++
		}
	}
	// 剩余部分合并
	res = append(res, left[l:]...)
	res = append(res, right[r:]...)
	return
}

// 分治法实现快速排序
func QuickSort(nums []int) []int {
	// 思路：每次找一个基准值，把一个数组分为左右两段，左段小于基准值，右段大于基准值，类似分治法没有合并过程
	quickSort(nums, 0, len(nums)-1)
	return nums
}

// 原地交换，所以传入交换索引
func quickSort(nums []int, start, end int) {
	if start < end {
		// 分治法
		pivot := partition(nums, start, end) // 找中间值所处的位置
		quickSort(nums, 0, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

// 分区
func partition(nums []int, start, end int) int {
	p := nums[end] // 找数组最后一位作为基准值
	i := start     // 从头开始的游标
	for j := start; j < end; j++ {
		// 把小于基准值的数字都放在最左边
		if nums[j] < p {
			swap(nums, i, j)
			i++
		}
	}
	// 把中间的值换为用于比较的基准值
	swap(nums, i, end)
	return i
}
func swap(nums []int, i, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}

// 分治法实现找出二叉树的最大深度
func maxDepth(root *TreeNode) int {
	// 返回条件
	if root == nil {
		return 0
	}

	// 分段处理
	// 先处理左子树
	left := maxDepth(root.Left)
	// 再处理右子树
	right := maxDepth(root.Right)

	// 合并结果
	if left > right {
		return left + 1
	}
	return right + 1
}

// 分治法实现判断是否是高度平衡的二叉树
// 思路：用分治法，左边是平衡树 && 右边也是平衡树 && 左右两边高度差<=1。
// 因为需要返回是否是平衡树以及树高度，所以要么返回两个值，要么返回一个值，用其二义性，-1表示非平衡树，>0表示树高度
// 二义性：一个变量有两种含义
func isBalanced(root *TreeNode) bool {
	if balanceMaxDepth(root) == -1 {
		return false
	}
	return true
}

// 这里返回值具有二义性。但一般工程中，建议结果通过两个变量来返回，不建议用一个变量表示两种含义
func balanceMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := balanceMaxDepth(root.Left)
	right := balanceMaxDepth(root.Right)
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

// 分治法实现：给定一个非空二叉树，返回其最大路径和
// 思路：分为三种情况，左子树最大路径和最大，右子树最大路径和最大，左右子树加根节点最大路径和最大
// 需要保存两个变量：保存子树最大路径和，保存左右子树加根节点和，然后比较两个变量选择最大值即可
type ResultType struct {
	SinglePath int // 保存单边最大值
	MaxPath    int // 保存最大值（单边或者两个单边+根的值）
}

func maxPathSum(root *TreeNode) int {
	result := helper(root)
	return result.MaxPath
}
func helper(root *TreeNode) ResultType {
	// check
	if root == nil {
		return ResultType{
			SinglePath: 0,
			MaxPath:    math.MinInt, // 考虑了节点Val值可能为负的情况
		}
	}
	// divide
	left := helper(root.Left)
	right := helper(root.Right)
	// conquer
	result := ResultType{}
	// 求单边最大值
	if left.SinglePath > right.SinglePath {
		result.SinglePath = max(left.SinglePath+root.Val, 0) //考虑了节点Val值可能为负的情况
	} else {
		result.SinglePath = max(right.SinglePath+root.Val, 0)
	}
	// 求两边加根节点最大值
	maxPath := max(right.MaxPath, left.MaxPath)                              // 比较左、右子树最大路径和
	result.MaxPath = max(maxPath, left.SinglePath+right.SinglePath+root.Val) // 比较单边子树最大路径和、左右子树加根节点最大路径和
	return result
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 分治法实现：找到二叉树中两个指定节点的最近公共祖先节点
// 思路：列出所有直接返回的条件，左子树或者右子树中有公共祖先，就先返回子树的祖先，否则返回根节点
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p == root || q == root {
		return root
	}

	// 左子树和右子树单独处理
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 合并结果 conquer
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

// BFS层级遍历二：给定一个二叉树，返回其节点值自底向上的层次遍历，即从叶子节点向上访问到根节点
// 思路，在层级遍历的基础上，翻转一下结果即可
func levelOrderBottom(root *TreeNode) [][]int {
	result := levelOrder(root)
	// 翻转结果
	reverse(result)
	return result
}
func reverse(result [][]int) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

// 给定一个二叉树，返回其节点值的锯齿形层次遍历。即Z字形遍历，一层从左到右，一层从右到左
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	toggle := false
	for len(queue) != 0 {
		list := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if toggle {
			reverseNums(list)
		}
		result = append(result, list)
		toggle = !toggle
	}
	return result
}
func reverseNums(list []int) {
	l := len(list)
	for i := 0; i < l/2; i++ {
		list[i], list[l-i-1] = list[l-i-1], list[i]
	}
}

// 判断一个二叉树是否是有效的二叉搜索树
// 二叉搜索树：要么是一棵空树，要么具有如下性质：
// 左子树上所有节点的值都小于根节点，右子树上所有节点的值均大于根节点，且左右子树又分别都是一个二叉搜索树
// 思路一：中序遍历，节点从小到大是有序的；思路二：分治法，判断左Max < 根 < 右MIN
// v1
func isValidBST(root *TreeNode) bool {
	result := make([]int, 0)
	inOrder(root, &result)
	// check order
	for i := 0; i < len(result)-1; i++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true
}
func inOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, result)
	*result = append(*result, root.Val)
	inOrder(root.Right, result)
}

// v2
type result struct {
	isValid bool
	max     *TreeNode
	min     *TreeNode
}

func isValidBSTV2(root *TreeNode) bool {
	res := helperV2(root)
	return res.isValid
}
func helperV2(root *TreeNode) result {
	res := result{}
	if root == nil {
		res.isValid = true
		return res
	}

	left := helperV2(root.Left)
	right := helperV2(root.Right)

	if !left.isValid || !right.isValid {
		res.isValid = false
		return res
	}

	if left.max != nil && left.max.Val >= root.Val {
		res.isValid = false
		return res
	}

	if right.min != nil && right.min.Val <= root.Val {
		res.isValid = false
		return res
	}

	res.isValid = true
	res.min = root
	if left.min != nil {
		res.min = left.min
	}
	res.max = root
	if right.max != nil {
		res.max = right.max
	}
	return res
}

// 给定一个二叉搜索树的根节点和要插入树中的值，将值插入二叉搜索树并返回插入后二叉搜索树的根节点
// 二叉搜索树定义：一颗空树，或者具有下列性质的二叉树：
// 1.若任意节点的左子树不空，则左子树上所有节点的值均小于它的根节点的值
// 2.若任意节点的右子树不空，则右子树上所有节点的值均大于它的根节点的值
// 3.任意节点的左、右子树也分别为二叉搜索树
// 思路：找到最后一个叶子节点满足插入条件即可
// DFS查找插入位置
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// 总结：掌握二叉树递归和非递归前中后序遍历；掌握DFS的前序遍历和分治法；掌握BFS层次遍历
