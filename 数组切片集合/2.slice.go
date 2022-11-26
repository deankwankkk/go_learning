package main

import "fmt"

func main() {
	arr := [6]int{0, 1, 2, 3, 4, 5}

	// getOutOfBoundValue(arr[:])
	sliceOp(arr[:])
}

func basicUse(arr []int) {
	// fmt.Println(arr[:]) // [0 1 2 3 4 5]
	// fmt.Println(arr[2:]) // [2 3 4 5]
	// fmt.Println(arr[:4]) // [0 1 2 3]
}

// slice可以向后扩展，不可以向前扩展
// s[i]不可以超越len(s)，向后扩展不可以超越底层数组cap(s)
func getOutOfBoundValue(arr []int) {
	s1 := arr[:3]
	fmt.Println(cap(arr)) // 6
	fmt.Println(s1[0]) // 0
	// fmt.Println(s1[3]) // index out of range [3] with length 3
	fmt.Println(s1[2:4]) // [2 3]
}

func sliceOp(arr []int) {
	// 创建slice
	newSlice := make([]int, 4)

	// 添加元素
	// slice进行append操作时，如果添加的元素超过cap，则会新分配更大的底层数组，原来的数组会根据是否被引用，决定是否被回收
	s1 := append(newSlice, 6) // [0 0 0 0 6]
	fmt.Println(s1)
	s2 := append(s1, 7)
	fmt.Println(s2) // [0 0 0 0 6 7]
	s3 := append(s2, 8)
	fmt.Println(s3) // [0 0 0 0 6 7 8]

	// 复制slice
	copy(s1, arr)
	fmt.Println(s1) // [0 1 2 3 4]

	// 删除某元素
	// 比如删除s1中的2, append的第二个参数是elems...类型，所以要展开
	// s1[3:]... 等于 3, 4
	s1 = append(s1[:2], s1[3:]...)
	// 删除头元素
	s1 = s1[1:]
	// 删除尾元素
	s1 = s1[:len(s1) - 1]

	fmt.Println(s1) // [1 3]
}