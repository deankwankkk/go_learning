package main

import "fmt"

func main() {
	// 基本不用，一般用于无限循环， Label不仅可以配合goto使用，也可以配合break和continue
	SIZE := 5
	for i := 0; i < SIZE; i++ {
		if i == 2 {
			// 执行label为L1的逻辑，并起到break作用
			goto L1
		}
		fmt.Println(i)
	}

L1:
	fmt.Println("goto")
}
