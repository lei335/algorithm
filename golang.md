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

注意点：
* 切片作为参数传递，只能查看修改，不能新增或者删除原始视频
* 默认s=s[0:len(s)]，取下限不取上限，数字表示为：[)

### map

基本用法

```go
// 创建
m := make(map[string]int)
// 设置kv
m["hello"]=1
// 删除k
delete(m,"hello")
// 遍历
for k,v := range m{
    println(k,v)
}

注意点：
* map的键值需要是可比较的，不能为slice、map、function、channel
* map值都有默认值，可以直接操作默认值，比如：m[money]++ 值由0变为1
* 比较两个map需要遍历，判断其中的kv是否相同，因为有默认值关系，所以需要检查val和ok两个值。判断两个map是否相同，需要使用`reflect`包的`DeepEqual()`方法

### 标准库

sort
```go
// int排序
sort.Ints([]int{})
// 字符串排序
sort.Strings([]string{})
// 自定义排序
sort.Slice(s,func(i,j int)bool{return s[i]<s[j]})
```

math
```go
// int32 最大最小值
math.MaxInt32 // 实际值：1<<31-1
math.MinInt32 // 实际值：-1<<31
// int64 最大最小值（int 默认就是int64）
math.MaxInt64
math.MinInt64
```

copy
```go
// 删除a[i]，可以用 copy 将 i+1 到末尾的值覆盖到i，然后末尾-1
copy(a[i:],a[i+1:])
a=a[:len(a)-1]

// make 创建长度，则通过索引赋值
a:= make([]int,n)
a[n-1]=x

// make 长度为0，则通过append()赋值
a := make([]int, 0)
a = append(a, x)
```

### 常用技巧

类型转换
```go
// byte 转数字
s := "12345" // s[0]类型是byte
num := int(s[0]-'0') // 1
str := string(s[0]) // "1"
b := byte(num+'0') // '1'
fmt.Printf("%d%s%c\n", num, str, b) // 111

// 字符串转数字
num, _ := strconv.Atoi(str)
str := strconv.Itoa()
```
