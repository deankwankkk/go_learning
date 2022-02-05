// Package tree
// go语言仅支持封装，不支持继承和多态
package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// CreateNode
// go中不需要构造函数，可以用自定义工厂函数替代
func CreateNode(value int) *Node {
	// 返回的是局部变量的地址
	return &Node{Value: value}
}

// Print
// 结构定义方法
// 等同func print(node treeNode) {}
// 调用方式不一样，以下使用node.print()调用
func (node Node) Print() {
	fmt.Println(node.Value)
}

// SetValue
// 只有使用指针才可以改变结构内容
// nil也可以调用结构的方法
func (node *Node) SetValue(value int) {
	if nil == node {
		fmt.Println("error")
		return
	}
	node.Value = value
}
