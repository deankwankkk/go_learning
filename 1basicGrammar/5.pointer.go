// Go语言只有值传递
package main

import "fmt"

func main() {
	// var a int = 2
	// var pa *int = &a
	// *pa = 3
	// fmt.Println(a)
	// fmt.Println(swap(3, 4))
	// enhanceNumber(a)
	// whoIsReal(a)
	// fmt.Println(a) // 2

	a, b := 666, 888
	swap(&a, &b)
	fmt.Println(a, b)
}

// func swap(a, b int) (int, int) {
// 	a, b = b, a
// 	return a, b
// }

// 指针写法
func swap(a, b *int) {
	*a, *b = *b, *a
}

func enhanceNumber(num int) {
	enhanceA := &num
	*enhanceA = 3
	fmt.Println(*enhanceA) // 3
}

func whoIsReal(num int) {
	fmt.Println(num) // 2
	realA := &num
	*realA = 4
	fmt.Println(*realA) // 4
}
