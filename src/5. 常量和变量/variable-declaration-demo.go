package main

import "fmt"

var lang, website string = "Go", "https://golang.org"
var compiled, dynamic bool = true, false
var announceYear int = 2009

// 省略变量类型
// 变量lang和dynamic的类型将被推断为内置类型string和bool。
// var lang, dynamic = "Go", false
// var compiled, announceYear = true, 2009
// var website = "https://golang.org"

// 省略初始值
// var lang, website string      // 两者都被初始化为空字符串。
// var interpreted, dynamic bool // 两者都被初始化为false。
// var n int                     // 被初始化为0。

// 和常量声明一样，多个变量可以用一对小括号组团在一起被声明。
// var (
//     lang, bornYear, compiled = "Go", 2007, true
//     announceAt, releaseAt int = 2009, 2012
//     createdBy, website string
// )

func main() {
	fmt.Printf("lang: %v\n", lang)
	fmt.Printf("website: %v\n", website)
	fmt.Printf("compiled: %v\n", compiled)
	fmt.Printf("dynamic: %v\n", dynamic)
	fmt.Printf("announceYear: %v\n", announceYear)
}
