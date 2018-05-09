package main

import "fmt"

/**
 * @description 
 * @time 2018/5/10 0:02
 * @version 
 */
func main() {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 20
	persons["李四"] = 23
	persons["王五"] = 26

	var w welcome = "大家好"

	//HandlerFunc(w.selfInfo)不是方法的调用，而是转型，
	//因为selfInfo和HandlerFunc是同一种类型，所以可以强制转型
	Each(persons, HandlerFunc(w.selfInfo))

	//EachFunc(persons, w.selfInfo)

	EachFunc(persons, selfInfo)

}

type welcome string

type HandlerFunc func(k, v interface{})

//func (w welcome) Do(k, v interface{}) {
//	fmt.Printf("%s,我叫%s,今年%d岁\n", w, k, v)
//}

func selfInfo(k, v interface{}) {
	fmt.Printf("我叫%s,今年%d岁\n", k, v)
}

func (w welcome) selfInfo(k, v interface{}) {
	fmt.Printf("%s,我叫%s,今年%d岁\n", w, k, v)
}

func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}

//新增了一个EachFunc函数，帮助调用者强制转型，调用者就不用自己做了
func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	Each(m,HandlerFunc(f))
}
