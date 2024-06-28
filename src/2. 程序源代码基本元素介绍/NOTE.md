# 一个简单的Go示例程序
一个简单的示例程序
```go
package main // 指定当前源文件所在的包名

import "math/rand" // 引入一个标准库包

const MaxRand = 16 // 声明一个具名整形常量

// 一个函数声明
/*
	StatRandomNumbers 生成一些不大于MaxRand的非负
	随机整数，并统计和返回小于和大于MaxRand/2的随机数
	个数。输入参数numRands指定了要生成的随机数的总数。
*/
func StatRandomNumbers(numRands int) (int, int) {
	// 声明了两个变量（类型都为int 初始值都为0）
	var a, b int
	// for 循环
	for i := 0; i < numRands; i++ {
		// 一个if-else条件控制代码块
		if rand.Intn(MaxRand) < MaxRand/2 {
			a += 1
		} else {
			b++
		}
	}
	return a, b // 函数返回两个结果
}

// main 函数，或主函数，是一个程序的入口函数
func main() {
	var num = 200
	// 调用上面的 StatRandomNumbers 函数,
	// 并将结果赋值给使用短声明语句声明的两个变量
	x, y := StatRandomNumbers(num)
	// 调用两个内置函数（print 和 println）。
	print("Result: ", x, " + ", y, " = ", num, " ? ")
	println(x+y == num)
}

```
> 编译器会自动推导出变量 `i`、`num`、`x`和
`y`的类型均为int类型，因为它们的初始值都是整型字面量表示的

> 注意，一般 `print` 和 `println` 这两个内置函数并不推荐使用。 在正式的项目中，我们应该尽
量使用 `fmt` 标准库包中声明的相应函数。 

一个Go程序的主函数（main函数）必须被声明在一个名称为main的包中

## 关于代码断行
很多左大括号 `{` 不能被放到下一行
1. 它们使得Go程序编译得非常快。
2. 它们使得不同的Go程序员编写的代码风格类似，从而一个Go程序员写的代码很容易被另一个
程序员看懂。