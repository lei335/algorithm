package stack

import "strconv"

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
				for k:=len(temp)-1;k>=0;k--{
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
