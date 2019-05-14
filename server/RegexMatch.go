package main

import (
	"fmt"
	"regexp"
)

func main() {

	//字符串出现了数字，字母，下划线
	//matched, _ := regexp.MatchString(`\w`, "abc123")
	//字符串中没有出现数字，字母，下划线
	//matched, _ := regexp.MatchString(`\w`, "[]',/.;")
	//字符串中出现了数字就能匹配
	//matched, _ := regexp.MatchString(`\d`, "1[]',/.;")
	//字符串中出现了非数字就匹配
	//matched, _ := regexp.MatchString(`\D`, "1[]',/.;")
	//字符串开始第一个字符，为非数字才能匹配
	//matched, _ := regexp.MatchString(`^\D`, "a[]',/.;")
	//从开始到结束就一个字符，所以会验证整个字符串的个数长度为1才能匹配
	//matched, _ := regexp.MatchString(`^\D$`, "a123123s")
	//一个正则字符代表一个字符 才能匹配上
	//matched, _ := regexp.MatchString(`^\D\d$`, "a1")
	//前面多个字母，后面多个数字
	//matched, _ := regexp.MatchString(`^\D+\d+$`, "aaaa1111")
	//fmt.Println(matched)

	//整个字符串中是否有(包含)这样的字符串 前面是数字，后面是小写字母或者数字
	compile := regexp.MustCompile(`\d[a-z0-9]`)

	//整个字符串中发现有2个片段，满足要求
	fmt.Println(compile.MatchString("aaa1g"))

	//n 代表取出符合提交的前n个 -1 就是全部
	fmt.Println(compile.FindAllString("aaa1ggg111", -1))

	//n 代表 n > 0 返回最多n个子字符串 n < 0 返回所有字符串
	fmt.Println(compile.Split("aaa1ggg111", -1))

	fmt.Println(compile.ReplaceAllString("aaa1ggg111", "替换"))
}
