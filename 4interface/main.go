package main

import (
	"fmt"
	"learnGo/4interface/mock"
	"learnGo/4interface/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"This is fake contents"}
	r = real.Retriever{}
	//fmt.Println(download(r))

	// 判断获取的值类型
	// 通过xx.(type)进行类型判断  Type assertion
	//if mockRetriever, ok := r.(mock.Retriever); ok {
	//	fmt.Println(mockRetriever.Contents)
	//} else {
	//	fmt.Println("not a mock retriever")
	//}

	inspect(r)
}

func inspect(r Retriever) {
	// 通过type进行类型判断
	switch r.(type) {
	case mock.Retriever:
		fmt.Println("is mock retriever")
	case real.Retriever:
		fmt.Println("is real retriever")
	}
}
