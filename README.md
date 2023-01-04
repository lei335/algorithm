# algorithm
algorithm practice

众所周知，代码不能只看，要多写。
所以这个仓库记录自己对算法知识的练习（用 go 语言实现）。

注意：
一旦报错“超出时间限制”，就需要检查"for"循环，一定是陷入了无限循环中。
在合并两个链表并排序时，把剩余的单个链表合并上来，用这种"for"循环的方式：
```go
for left!=nil{
    head.Next=left
    left=left.Next
    head=head.Next
}
for right!=nil{
    head.Next=right
    right=right.Next
    head=head.Next
}
```
会比用"if"的方式：
```go
if left!=nil{
    head.Next=left
}
if right!=nil{
    head.Next=right
}
```
要快一些，并且内存消耗也要小一些。