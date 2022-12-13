package main

import "fmt"

func main() {
	color := "yellow"

	// 基本使用
	switch color {
	case "green":
		fmt.Println("go")
	case "red":
		fmt.Println("stop")
	default:
		fmt.Println("waiting")
	}

	// fallthrough，后面有fallthrough，即使满足，也会继续往下判断
	switch color {
	case "yellow":
		fmt.Printf("color is %s\n", color)
		fallthrough
	case "blue":
		fmt.Printf("color is %s\n", color)
	default:
		fmt.Printf("color is nil")
	}

	// switch type，判断数据类型
	var num interface{} = 8.8

	switch num.(type) {
	case int:
		fmt.Println("num is int")
	case float64:
		fmt.Println("num is float64")
	case byte, string, bool:
		fmt.Println("num is byte or string or boolean")
	default:
		fmt.Println("num是其他类型")
	}
}
