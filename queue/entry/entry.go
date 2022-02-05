package main

import (
	"fmt"
	"learnGo/queue"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	fmt.Println(q)
	popNumber := q.Pop()
	fmt.Println(popNumber)
	isEmpty := q.IsEmpty()
	fmt.Println(isEmpty)
}
