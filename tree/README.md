### 结构体
```go
type myNode struct {}

var node myNode
node1 := myNode{}
```

### 封装
> go语言中，封装是使用CamelCase，首字母大写：public，首字母小写：private

```go
// 扩充系统类型或已有的类型
// 定义别名 参考queue --- 最简单

// 使用组合 --- 最常用
type myNode struct {
	node *tree.Node
}

left := myNode{node.left}

// 使用内嵌方式 embedding --- 代码最少
// 可以将需要扩展的类型，写入struct中，可以直接使用Node的变量和函数
type myTreeNode struct {
	*tree.Node
}
```

### 包
> 每个目录一个包；main包包含可执行入口；为结构定义的方法必须放在同一个包内，但可以是不同文件
