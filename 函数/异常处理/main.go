package main

import (
	"errors"
	"fmt"
	"time"
)

// PathError 自定义error
type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

func NewPathError(path, op, message string) PathError {
	return PathError{
		path:       path,
		op:         op,
		createTime: time.Now().Format("2006-01-02"),
		message:    message,
	}
}

// 实现Error()
func (e PathError) Error() string {
	return e.createTime + ": " + e.op + " " + e.path + " " + e.message
}

func deletePath(path string) error {
	return NewPathError(path, "delete", "path is not exist")
}

/*
panic发生场景：
1. 运行时错误会导致panic，比如数组越界
2. 程序主动调用panic()

panic会执行什么:
1. 逆序执行当前goroutine的defer链
2. 打印错误信息和调用堆栈
3. 调用exit(2)结束整个进程
*/
// recover截断panic
func recoverPanic() {
	defer func() {
		// recover只能写在defer中
		// recover会截断panic，输出自定义错误信息
		if err := recover(); err != nil {
			fmt.Printf("recoverPanic函数发生了panic：%s\n", err)
		}
	}()
	panic(errors.New("发生了错误"))
}
func main() {
	//err := deletePath("abc")
	//fmt.Println(err)
	recoverPanic()
	fmt.Println("程序没有退出，仍然执行")
}
