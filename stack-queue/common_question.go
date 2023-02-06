package stack

import (
	"strconv"

	"github.com/study/algorithm/tree"
)

// 栈是允许在同一端进行插入和删除操作的特殊线性表

// 设计一个支持push, pop, top操作，并能在常数时间内检索到最小元素的栈
// 思路：用两个栈实现，一个最小栈始终保证最小值在顶部
type MinStack struct {
	min   []int
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		min:   make([]int, 0),
		stack: make([]int, 0),
	}
}
func (this *MinStack) Push(x int) {
	min := this.GetMin()
	if x < min {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, min)
	}
	this.stack = append(this.stack, x)
}
func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}
func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}
func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 1 << 31
	}
	min := this.min[len(this.min)-1]
	return min
}

// 波兰表达式计算 > 输入：["2", "1", "+", "3", "*"] > 输出：9
// 解释：((2 + 1) * 3) = 9
// 思路： 通过栈保存原来的元素，遇到表达式弹出运算，再推入结果， 重复这个过程
func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		// 遇到表达式则弹出栈中数字进行计算，再把结果推入栈中
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return -1 // 表达式有问题
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var res int
			switch tokens[i] {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				res = a / b
			}
			stack = append(stack, res)
		// 遇到数值，则将其先保存在栈中
		default:
			val, err := strconv.Atoi(tokens[i])
			if err != nil {
				return -1
			}
			stack = append(stack, val)
		}
	}
	return stack[0]
}

// 给定一个经过编码的字符串，返回它解码后的字符串，s = "3[a]2[bc]",返回"aaabcbc". s = "3[a2[c]]",返回"accaccacc".
// s = "2[abc]3[cd]ef",返回"abcabccdcdcdef".
// 思路：通过栈辅助进行操作
func decodeString(s string) string {
	if s == "" {
		return s
	}
	// 先遍历s
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		// 当遇到']'时，就需要回顾前面的字符进行处理了
		case ']':
			// 先将']'前面的字符逆向记录下来
			temp := make([]byte, 0)
			for len(stack) != 0 && stack[len(stack)-1] != '[' {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				temp = append(temp, v)
			}
			// 遇到对应的'['，将其弹出
			stack = stack[:len(stack)-1]
			// 然后找'['前面的数字
			// 先找到前面数字的数位（1位数还是2位数，还是其他位数）
			idx := 1
			for len(stack) >= idx && stack[len(stack)-idx] >= '0' && stack[len(stack)-idx] <= '9' {
				idx++
			}
			// 有了数位后，就可以直接获取字符串格式的数字了
			num := stack[len(stack)-idx+1:] // 注意，idx初始化为1，所以这里还需要加1
			stack = stack[:len(stack)-idx+1]
			count, _ := strconv.Atoi(string(num))
			// 将解码的字符正向放回到stack中
			for j := 0; j < count; j++ {
				for k := len(temp) - 1; k >= 0; k-- {
					stack = append(stack, temp[k])
				}
			}
		// 除了']'之外，都只被放进stack中即可
		default:
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

// 利用栈进行DFS递归搜索模板
// boolean DFS(int root, int target) {
// 	Set<Node> visited;
// 	Stack<Node> s;
// 	add root to s;
// 	while(s is not empty) {
// 		Node cur = the top element in s;
// 		return true if cur is target;
// 		for(Node next:the neighbors of cur) {
// 			if(next is not in visited) {
// 				add next to s;
// 				add next to visited;
// 			}
// 		}
// 		remove cur from s;
// 	}
// 	return false;
// }

// 给定一个二叉树，返回它的中序遍历
// 思路：用栈保存已经访问过的元素，用来原路后进先出地返回
func inorderTraversal(root *tree.TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	sta := make([]*tree.TreeNode, 0)
	for root != nil || len(sta) > 0 {
		// 先把所有左子树节点保存下来
		for root != nil {
			sta = append(sta, root)
			root = root.Left
		}
		// pop 栈中最后一个节点
		cur := sta[len(sta)-1]
		sta = sta[:len(sta)-1]
		// 访问根节点值
		res = append(res, cur.Val)
		// 访问右子树
		root = cur.Right
	}
	return res
}

// 给定一个无向连通图中一个节点的引用，情返回该图的深拷贝（克隆）
// 思路： 递归遍历每一个连接节点，新初始化一个节点，和被克隆的这个节点绑定在一起（用map）
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	return clone(node, visited)
}
func clone(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	// 已经访问过，直接返回
	if v, ok := visited[node]; ok {
		return v
	}
	newNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	visited[node] = newNode
	for i := 0; i < len(node.Neighbors); i++ {
		newNode.Neighbors[i] = clone(node.Neighbors[i], visited)
	}
	return newNode
}

// 给定一个由‘1’（陆地）和‘0’（水）组成的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或者垂直方向上
// 相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。
// 思路：通过深度搜索遍历可能性（注意标记已访问元素）
func numIsland(grid [][]byte) int {
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && dfs(grid, i, j) >= 1 {
				count++
			}
		}
	}
	return count
}
func dfs(grid [][]byte, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] == '1' {
		// 标记已经访问过的（每一个节点只需要访问一次）
		grid[i][j] = 0 //对切片的修改会影响到numIsland函数中的切片值，修改切片的值会覆盖底层数组的值
		return dfs(grid, i-1, j) + dfs(grid, i, j-1) + dfs(grid, i+1, j) + dfs(grid, i, j+1) + 1
	}
	return 0
}

// 给定n个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为1.求在该柱状图中，能够勾勒出来的矩形的最大面积（不能包含柱状图未到达的区域）
// 思路：求以当前柱子为高度的最大矩形面积（从左右两边值开始比较，遇到较小值就停止），再比较得出最大的那个面积值。
// 思路：保证栈中保存的高度是依次递增的
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	stack := make([]int, 0)
	max := 0
	for i := 0; i <= len(heights); i++ {
		var cur int
		if i == len(heights) {
			cur = 0
		} else {
			cur = heights[i]
		}
		// 当前高度小于栈，则将栈内元素都弹出计算面积
		for len(stack) != 0 && cur <= heights[stack[len(stack)-1]] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 依次作为高度计算面积
			h := heights[pop]
			// 计算宽度
			w := i
			if len(stack) != 0 {
				beforeH := stack[len(stack)-1]
				w = i - beforeH - 1
			}
			max = Max(max, h*w)
		}
		// 记录索引即可获得对应高度
		stack = append(stack, i)
	}
	return max
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Queue队列
// 常用于BFS宽度优先搜索

// 使用栈实现队列
type MyQueue struct {
	stack []int
	back  []int
}

// 初始化结构体
func Constructor() MyQueue {
	return MyQueue{
		stack: make([]int, 0),
		back:  make([]int, 0),
	}
}
func (this *MyQueue) Push(x int) {
	for len(this.back) != 0 {
		val := this.back[len(this.back)-1]
		this.back = this.back[:len(this.back)-1]
		this.stack = append(this.stack, val)
	}
	this.stack = append(this.stack, x)
}
func (this *MyQueue) Pop() int {
	for len(this.stack) != 0 {
		val := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		this.back = append(this.back, val)
	}
	if len(this.back) == 0 {
		return 0
	}
	val := this.back[len(this.back)-1]
	this.back = this.back[:len(this.back)-1]
	return val
}
func (this *MyQueue) Peek() int {
	for len(this.stack) != 0 {
		val := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		this.back = append(this.back, val)
	}
	if len(this.back) == 0 {
		return 0
	}
	val := this.back[len(this.back)-1]
	return val
}
func (this *MyQueue) IsEmpty() bool {
	return len(this.stack) == 0 && len(this.back) == 0
}

// 二叉树层次遍历
func levelOrder(root *tree.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := make([]*tree.TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue) // 记录每一层的节点数
		level := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
