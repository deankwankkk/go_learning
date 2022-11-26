// 循环体只有for，没有while
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// for 循环简单应用
func sumTo100() {
	res := 0
	// for的条件里不需要括号
	for i := 1; i <= 100; i++ {
		res += i
	}
	fmt.Println(res)
}

// for循环，求整数的二进制
func convertToBin(num int) string {
	result := ""
	if num == 0 {
		result = strconv.Itoa(num) + result
		return result
	}

	for ; num > 0; num /= 2 {
		lsb := num % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// 循环读取某一文件
func readFileContent() {
	content, err := os.Open("1.basicGrammar/abc.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 死循环
func forever() {
	for {
		fmt.Println("死循环")
	}
}

func main() {
	// sumTo100()
	// fmt.Println(convertToBin(0))
	readFileContent()
}
