package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// if
func ifBranch() {
	if content, err := ioutil.ReadFile("1.basicGrammar/abc.txt"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s\n", content)
	}
}

// switch 后面可以没有表达式，也不需要break
func switchBranch(score int) string {
	res := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("invaild score: %d", score))
	case score < 60:
		res = "F"
	case score < 85:
		res = "B"
	case score <= 100:
		res = "A"
	}
	return res
}

func main() {
	// fmt.Println(switchBranch(59))
	ifBranch()
}