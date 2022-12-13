package main

import (
	"fmt"
	"time"
)

/*
可见性：
  1. 可见性统一规则：大写字母开头，跨package也可以访问，否则只能本package内部访问
  2. 结构体名称以大写开头时，package外部可见，在此前提下，结构体中以大写开头的成员变量或成员函数，在package外部也可见
*/

// User 定义结构体
type User struct {
	id         int
	Score      float32
	name, addr string
	enrollment time.Time
}

// Man 结构体中的匿名成员
type Man struct {
	id     int
	string // 匿名字段
	float32
}

func main() {
	// 匿名结构体，通常使用场景，只使用一次
	var student struct {
		Name string
		Addr string
	}
	student.Name = "dean"
	student.Addr = "China"

	//var u User
	u := User{id: 0, name: "kwan", addr: "shanghai", Score: 100}
	u.enrollment = time.Now()
	fmt.Printf("user's enrollment is %v\n", u.enrollment)

	// 匿名成员的使用
	m := Man{id: 1, string: "nothing", float32: 666}
	fmt.Printf("id is %d, second ele is %s, third ele is %g\n", m.id, m.string, m.float32)

	u.toHello()
}

// 成员函数，这里可以理解为User结构体中的函数
func (user User) toHello() {
	fmt.Printf("hello, %v\n", user.name)
}
