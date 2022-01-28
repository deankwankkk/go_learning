// 数组是值类型
package main

import "fmt"

func main() {
	// 声明方式
	var array1 [3]int // [0 0 0]
	array2 := [3]int{1, 2, 3}
	array3 := [...]int{11, 22, 33, 44}
	// 多维数组
	var array4 [4][5]int

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array4)

	// 数组遍历 range
	// 使用range的好处：意义明确，可以直接获取到index和value
	for _, v := range array3 {
		fmt.Printf("%d ", v)
	}
}
