package main

import "learnGo/tree"

func main() {
	//var root tree.Node
	root := tree.Node{Value: 3}

	//fmt.Println(root) // {3 nil nil}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	//nodes := []tree.Node{
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}

	//for _, i :=  range nodes {
	//	fmt.Println(i)
	//}

	//pRoot为nil，但也可以调用setValue()
	//var pRoot *tree.Node
	//pRoot.setValue(200)
}
