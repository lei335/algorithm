### 切片
go 通过切片模拟栈和队列 

栈

```go
// 创建栈
stack := make([]int, 0)
// push 压入
stack = append(stack, 10)
// pop 弹出
v := stack[len(stack) - 1]  // 取出弹出值
stack = stack[:len(stack)-1] // 弹出
// 检查栈空
len(stack) == 0
```

队列

```go
// 创建队列
queue := make([]int, 0)
// enqueue 入队
queue = append(queue, 10)
// dequeue 出队
v := queue[0]
queue = queue[1:]
// 长度0表示队列为空
len(queue) == 0
```
