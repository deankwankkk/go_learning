package main

import (
	"fmt"
)

// channel是一个环形队列，先进先出
func main() {
	// 基本用法
	var cha chan int                    // 默认值nil
	fmt.Printf("len is %d\n", len(cha)) // 0

	// 初始化长度为10的channel
	cha = make(chan int, 10)
	fmt.Printf("cap is %d\n", cap(cha)) // 10

	// 写入 && 读取
	for i := 0; i < 10; i++ {
		cha <- 8
	}
	value := <-cha
	fmt.Printf("value is %d\n", value)
	// 如果channel已经满了，会被阻塞
	cha <- 3
	// 如果不是直接进行<-取出操作，需要先close channel
	//close(cha)
	//for ele := range cha {
	//	fmt.Println(ele)
	//}
	// 对channel进行遍历，相当于<-cha（取出）操作
	//fmt.Printf("cha len is %d\n", len(cha))

	LENGTH := len(cha)
	for i := 0; i < LENGTH; i++ {
		ele := <-cha
		fmt.Printf("ele is ---> %d\n", ele)
	}
}
