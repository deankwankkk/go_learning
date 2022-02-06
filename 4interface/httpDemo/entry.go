package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.imooc.com/")

	if err != nil {
		panic("error")
	}

	defer res.Body.Close()

	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("%s/n", bytes)
}
