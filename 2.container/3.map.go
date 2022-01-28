package main

import "fmt"

func main() {
	// start()
	//getMostLengthOfChar()
	res := noRepeating()
	fmt.Println(res)
}

func start() {
	// 声明方式
	// 1. var mmap map[string]string
	// 2. mmap := map[string]string {}
	// 3. make(map[string]string)

	// map[key_type]value_type
	newMap := map[string]string{
		"name":    "dean",
		"address": "shanghai",
		"sex":     "male",
	}

	// map是无序的，是hashMap，可以使用range进行遍历
	// 如果要有序的，需要将key值放入slice中，slice会自动排序，然后再range
	// for key, value := range newMap {
	// 	fmt.Println(key, value)
	// }

	// 删除元素
	delete(newMap, "name")

	// 获取某key值，返回值有value 以及 是否存在的状态
	// 如果key不存在，则返回value类型的初始值
	if name, ok := newMap["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("key does not exist")
	}
}

// 一个字符串中不重复的字符长度
func getMostLengthOfChar() {
	template := "abcabcbb"
	// lastOccurred
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for k, v := range []byte(template) {
		if lastI, ok := lastOccurred[v]; ok && lastI >= start {
			start = lastI + 1
			// 3
		}

		if k-start+1 > maxLength {
			maxLength = k - start + 1 // 3
		}
		lastOccurred[v] = k // { 97: 3, 98: 4, 99: 5 }
	}
	fmt.Println(start, maxLength)
}

// 上面例题，支持中文语言
func noRepeating() int {
	s := "中国人爱中国"
	// rune相当于go的char
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for k, v := range []rune(s) {
		if lastI, ok := lastOccurred[v]; ok && lastI >= start {
			start = lastI + 1
		}

		if k-start+1 > maxLength {
			maxLength = k - start + 1
		}
		lastOccurred[v] = k
	}
	return maxLength
}
