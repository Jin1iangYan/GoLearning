# 基本类型和它们的字面量表示
## 基本内置类型
Go支持如下内置基本类型：
- 一种内置布尔类型：bool。
- 11种内置整数类型：int8、uint8、int16、uint16、int32、uint32、int64、uint64、int、uint和uintptr。
- 两种内置浮点数类型：float32和float64。
- 两种内置复数类型：complex64和complex128。
- 一种内置字符串类型：string。

> byte是uint8的内置别名。 我们可以将byte和uint8看作是同一个类型。
>
> rune是int32的内置别名。 我们可以将rune和int32看作是同一个类型。

### type 关键字
Go代码中可以出现很多布尔和字符串类型（数值类型也同样）
```go
// 一些类型定义声明
type status bool // status和bool是两个不同的类型
type MyString string // MyString和string是两个不同的类型
type Id uint64 // Id和uint64是两个不同的类型
type real float32 // real和float32是两个不同的类型

// 一些类型别名声明
type boolean = bool // boolean和bool表示同一个类型
type Text = string // Text和string表示同一个类型
type U8 = uint8 // U8、uint8和 byte表示同一个类型
type char = rune // char、rune和int32表示同一个类型
```
