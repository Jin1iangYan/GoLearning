# 函数声明和调用
## 函数声明
让我们来看一个函数声明：
```go
func SquaresOfSumAndDiff(a int64, b int64) (s int64, d int64) {
	x, y := a+b, a-b
	s = x * x
	d = y * y
	return // <=> return s, d
}
```
- 输出结果声明列表在输入参数列表后面。在Go中，一个函数可以有多个返回值。
- 当一个函数的输出结果声明列表为空或者只包含一个匿名结果声明时，此列表可以不用一对小括号括起来；否则，小括号是必需的。
- 输出结果声明列表中的所有声明中的结果名称可以（而且必须）同时出现或者同时省略。 如果一个返回结果声明中的结果名称没有省略，则这个返回结果称为具名返回结果。否则称为匿名返回结果。

不允许匿名返回结果和具名返回结果混合使用，下面编译错误：
```go
func test() (a string, string) () {
	return "nihao", "nihao"
}
// syntax error: mixed named and unnamed parameters
```

这样的写法？？`string` 关键字被当成了变量，要避免这种写法！
```go
func test() (string, b int) {
	string = 100
	return string, b
}

func main() {
	fmt.Println(test())
}
```
下面是更多函数声明的例子：
```go
func CompareLower4bits(m, n uint32) (r bool) {
	// 下面这两行等价于：return m&0xFF > n&0xff
	r = m&0xF > n&0xf
	return
}

// 此函数没有输入参数。它的结果声明列表只包含一个
// 匿名结果声明，因此它不必用()括起来。
func VersionString() string {
	return "go1.0"
}

// 此函数没有返回结果。它的所有输入参数都是匿名的。
// 它的结果声明列表为空，因此可以被省略掉。
func doNothing(string, int) {
}
```

## 函数调用
略

## 函数调用的退出阶段
在Go中，当一个函数调用返回后（比如执行了一个return语句或者函数中的最后一条语句执行完毕）， 此调用可能并未立即退出。一个函数调用从返回开始到最终退出的阶段称为此函数调用的退出阶段（exiting phase）。 函数调用的退出阶段的意义将在讲解延迟函数的时候体现出来。

## 匿名函数
一个匿名函数的定义中不包含函数名称部分。 注意匿名函数定义不是一个函数声明。一个匿名函数在定义后可以被立即调用。

```go
package main

func main() {
	// 这个匿名函数没有输入参数，但有两个返回结果。
	x, y := func() (int, int) {
		println("This function has no parameters.")
		return 3, 4
	}() // 一对小括号表示立即调用此函数。不需传递实参。

	// 下面这些匿名函数没有返回结果。

	func(a, b int) {
		println("a*a + b*b =", a*a+b*b) // a*a + b*b = 25
	}(x, y) // 立即调用并传递两个实参。

	func(x int) {
		// 形参x遮挡了外层声明的变量x。
		println("x*x + y*y =", x*x+y*y) // x*x + y*y = 32
	}(y) // 将实参y传递给形参x。

	func() {
		println("x*x + y*y =", x*x+y*y) // x*x + y*y = 25
	}() // 不需传递实参。
}
```

Go中的所有的自定义函数（包括声明
的函数和匿名函数）都可以被视为闭包。 

## 内置函数
我们可以不引入任何库包（见下一篇文章）而调用一个内置函数。
