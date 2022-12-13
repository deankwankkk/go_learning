package main

import "fmt"

type User1 struct {
	id   int
	name string
}

func main() {
	// 创建结构体指针
	//var u User
	//user := &u // 通过取址符& 得到指针

	// 直接创建结构体指针
	user := &User1{
		id:   3,
		name: "kkk",
	}
	fmt.Printf("user pointer is %p\n", user)

	// new()函数实体化一个结构体，并返回其指针
	user = new(User1)
	fmt.Printf("new user pointer is %p\n", user)
}
