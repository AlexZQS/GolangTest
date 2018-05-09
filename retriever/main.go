package main

import (
	"TestImmoc/retriever/mock"
	"fmt"
)

/**
 * @description 
 * @time 2018/4/16 22:42
 * @version 
 */


type Retriever interface {
	Get(url string) string
}

func main() {
	var r Retriever
	retriever := &mock.Retriever{"i am alex"}
	r = retriever
	fmt.Println(r.Get("\nthis is a test"))
}