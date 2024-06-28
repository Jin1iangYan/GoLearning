# 关键字和标识符
## 关键字
截至目前（Go 1.22），Go中共有25个关键字。
```go
1| break default func interface select
2| case defer go map struct
3| chan else goto package switch
4| const fallthrough if range type
5| continue for import return var
```
这些关键字可以分为四组：
- const、func、import、package、type和var用来声明各种代码元素。
- chan、interface、map和struct用做 一些组合类型的字面表示中。
- break、case、continue、default、 else、fallthrough、for、 goto、if、range、 return、select和switch用在流程控制语句中。 详见基本流程控制语法（第12章）。
- defer和go也可以看作是流程控制关键字， 但它们有一些特殊的作用。详见协程和延迟函数调用（第13章）。

## 标识符

`_` 是一个特殊的标识符，空标识符

由大写字母开头的标识符称为 __导出标识符__（即public）。其他为 __非导出标识符__（private）
