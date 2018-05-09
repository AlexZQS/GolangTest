package main

import "fmt"

/**
 * @description 
 * @time 2018/5/9 1:03
 * @version 
 */
func main() {
}

func sum(a, b int) int {
	return a + b
}

func function(a, b int, sm func(int,int) int) {
	fmt.Println(sm(a,b))
}
