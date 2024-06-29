package main

// 声明了两个单独的具名常量。（是的，
// 非ASCII字符可以用做标识符。）
const π = 3.1416
const Pi = π // 等价于：const Pi = 3.1416

// 声明了一组具名常量。
const (
	No         = !Yes
	Yes        = true
	MaxDegrees = 360
	Unit       = "弧度"
)

// 常量声明中的自动补全
const (
	X float32 = 3.14
	Y         // 这里必须只有一个标识符
	Z         // 这里必须只有一个标识符

	A, B = "Go", "language"
	C, _
	// 上一行中的空标识符是必需的（如果
	// 上一行是一个不完整的常量描述的话）。
)

func main() {
	// 声明了三个局部具名常量。
	const DoublePi, HalfPi, Unit2 = π * 2, π * 0.5, "度"
}
