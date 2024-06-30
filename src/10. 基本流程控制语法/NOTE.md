# 基本流程控制语法

## Go中的流程控制语句简单介绍
Go语言中有三种基本的流程控制代码块：
- if-else条件分支代码块；
- for循环代码块；
- switch-case多条件分支代码块。

Go中另外还有几种和特定种类的类型相关的流程控制代码块：
- 用来遍历整数、各种容器和通道的for-range循环代码块。
- 接口相关的type-switch多条件分支代码块。
- 通道相关的select-case多分支代码块。

和很多其它流行语言一样，Go也支持 `break`、`continue` 和 `goto` 等跳转语句。 另外，Go还支持一个特有的 `fallthrough` 跳转语句。

## if-else 条件分支控制代码块
一个if-else条件分支控制代码块的完整形式如下：
```go
package main

func main() {

	if InitSimpleStatement; Condition {
		// do something
	} else {
		// do something
	}
}
```
- `InitSimpleStatement` 部分是可选的，如果它没被省略掉，则它必须为一条简单语句。 如果它被省略掉，它可以被视为一条空语句（简单语句的一种）。 在实际编程中，`InitSimpleStatement` 常常为一条变量短声明语句。
- `Condition` 必须为一个结果为布尔值的表达式（它被称为条件表达式）。`Condition` 部分可以用一对小括号括起来，但大多数情况下不需要。

## for循环代码块
for循环代码块的完整形式如下：
```go
for InitSimpleStatement; Condition; PostSimpleStatement { 
   // do something 
} 
```
- `InitSimpleStatement`、`Condition` 和 `PostSimpleStatement` 都是可选的。
- `InitSimpleStatement`（初始化语句）和 `PostSimpleStatement`（步尾语句）两个部分必须均为简单语句，并且 `PostSimpleStatement` 不能为一个变量短声明语句。
- `Condition` 必须为一个结果为布尔值的表达式（它被称为条件表达式）。

下面是一个使用for循环流程控制的例子。此程序将逐行打印出0到9十个数字：
```go
for i := 0; i < 10; i++ { 
   fmt.Println(i) 
} 
```

> - 在Go 1.22之前，每一个声明的循环变量在整个循环的执行过程中只会被实例化一次。此唯一的实例将被所有循环步共享。
> - 从Go 1.22开始，每一个声明的循环变量将会在每个循环步被实例化一次。每个实例只作用于当前循环步。

## for-range流程控制代码块用来遍历整数
for-range流程控制代码块可以用来遍历整数、各种容器和通道等。 本文只介绍如何使用for-range流程控制代码块来遍历整数。
> 注意：使用for-range流程控制代码块来遍历整数是从Go 1.22才开始支持的。

例子：
```go
package main

import (
	"fmt"
)

func main() {
	for i := range 10 {
		fmt.Printf("i: %v\n", i)
	}
}
```

## switch-case流程控制代码块
一个switch-case流程控制代码块的完整形式为：
```go
package main

func main() {
	switch InitSimpleStatement; CompareOperand0 {
	case CompareOperandList1:
		// do something
	case CompareOperandList2:
		// do something
	case CompareOperandListN:
		// do something
	default:
		// do something
	}
}
```
- `InitSimpleStatement` 语句和 `CompareOperand0` 表达式都是可选的。 如果 `CompareOperand0` 表达式被省略，则它被认为类型为bool类型的 `true` 值。 如果 `InitSimpleStatement` 语句被省略，其后的分号也可一并被省略。
- `InitSimpleStatement` 部分必须为一条简单语句，它是可选的。
- `CompareOperand0` 部分必须为一个表达式（如果它没被省略的话，见下）。 此表达式的估值结果总是被视为一个类型确定值。如果它是一个类型不确定值，则它被视为类型为它的默认类型的类型确定值。 因为这个原因，此表达式不能为类型不确定的nil值。`CompareOperand0` 常被称为 `switch` 表达式。
- 每个 `CompareOperandListX` 部分（X表示1到N）必须为一个用（英文）逗号分隔开来的表达式列表。 其中每个表达式都必须能和 `CompareOperand0` 表达式进行比较。 每个这样的表达式常被称为 `case` 表达式。 如果其中 `case` 表达式是一个类型不确定值，则它必须能够自动隐式转化为对应的 `switch` 表达式的类型，否则编译将失败。

switch-case代码块属于可跳出流程控制。 `break` 可以使用在一个switch-case流程控制的任何分支代码块之中以提前跳出此switch-case流程控制。

和很多其它语言不一样，每个分支代码块的结尾不需要一条 `break` 语句就可以自动跳出当前的switch-case流程控制。 那么如何让执行从一个case分支代码块的结尾跳入下一个分支代码块？Go提供了一个 `fallthrough` 关键字来完成这个任务。
- 一条fallthrough语句必须为一个分支代码块中的最后一条语句。
- 一条fallthrough语句不能出现在一个switch-case流程控制中的最后一个分支代码块中。

Go中另外一个和其它语言的显著不同点是 `default` 分支不必一定为最后一个分支。

## 包含跳转标签的break和continue语句
解决 break 只能作用当前控制块的问题：
```go
package main

import "fmt"

func FindSmallestPrimeLargerThan(n int) int {
Outer:
	for n++; ; n++ {
		for i := 2; ; i++ {
			switch {
			case i*i > n:
				break Outer
			case n%i == 0:
				continue Outer
			}
		}
	}
	return n
}

func main() {
	for i := 90; i < 100; i++ {
		n := FindSmallestPrimeLargerThan(i)
		fmt.Print("最小的比", i, "大的素数为", n)
		fmt.Println()
	}
}

```
