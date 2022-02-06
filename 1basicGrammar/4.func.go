/**
 *  1. 返回值类型写在最后面
 *  2. 可返回多个值
 *  3. 函数作为参数
 *  4. 没有默认参数，可选参数
 */
package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func main() {
	// fmt.Println(calcNumber("@", 1, 2))
	// fmt.Println(multiReturnValue(3, 2))
	// fmt.Println(argumentIsFunc(pow, 2, 3))
	fmt.Println(sumNum(1, 2, 3, 4, 5))
}

// 常规函数
// 包含error信息
func calcNumber(op string, a, b int) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s" + op)
	}
}

// 返回多值, 返回除数结果及余数
// 注意点： 返回多值时，可以为值命名，仅用于非常简单的函数
func multiReturnValue(a, b int) (int, int) {
	return a/b, a%b
}

// 参数可以是函数
func argumentIsFunc(fn func(int, int) int, a, b int) int {
	p := reflect.ValueOf(fn).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args " + "(%d %d)\n", opName, a, b)
	return fn(a, b)
}

// 重写pow方法
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表
func sumNum(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum
}