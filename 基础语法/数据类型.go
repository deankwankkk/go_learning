/**
 * 1. 变量类型写在变量名之后 e.g: var age int
 * 2. 编译器可推测变量类型  e.g: name := "kwan" // name is string
 * 3. 没有char，只有rune // rune是32位
 * 4. 原生支持复数类型
 * 5. 内置变量类型：
 * 		5-1. bool, string
 *    5-2. (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr // u代表无符号， uintptr: ptr是指针的意思，int类型字节数是依据系统的位数
 *    5-3. byte, rune // rune是字符型，32位，等同于char，由于国际化语言出现
 *    5-4. float32, float64, complex64, complex128 // complex是复数
 */
package main

import (
	"fmt"
	"math"
)

// 变量定义具有作用域范围，并且全局变量不能使用简写声明
var (
	name = "kwan"
	age = 22
)

func variableValue() {
	var a int = 1
	var b string = "kwan"
	fmt.Printf("%d, %q\n", a, b)
	fmt.Println(name, age)
}

func variableShorter() {
	a, b, c, d := 0, "kwan", false, 2.11
	b = "dean"
	fmt.Println(a, b, c, d)
}

// 强制类型转换
func variableTypeTransform() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 常量
func consts() {
	// const 数值可以作为各种类型使用
	const (
		filename = "deankwan"
		a, b = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

// 枚举
func enums() {
	// 1. 普通枚举类型
	// const (
	// 	cpp = 0
	// 	java = 1
	// 	python = 2
	// 	golang = 3
	// )

	// 2. 自增值枚举类型
	// 如果枚举是递增的，第一个枚举值可以使用iota，后面的枚举不需要赋值
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)
	fmt.Println(cpp, javascript, python, golang)

	// b, kb, mb, gb, tb, pb
	// iota也可以参与运算
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	// variableValue()
	// variableShorter()
	// variableTypeTransform()
	// consts()
	enums()
}