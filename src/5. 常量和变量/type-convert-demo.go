package main

import "fmt"

func main() {
	// 结果为complex128类型的1.0+0.0i。虚部被舍入了。
	fmt.Printf("complex128(1 + -1e-1000i): %v\n", complex128(1+-1e-1000i))
	// 结果为float32类型的0.5。这里也舍入了。
	fmt.Printf("float32(0.49999999): %v\n", float32(0.49999999))
	// 只要目标类型不是整数类型，舍入都是允许的。
	fmt.Printf("float32(17000000000000000): %v\n", float32(17000000000000000))
	fmt.Printf("float32(123): %v\n", float32(123))
	fmt.Printf("uint(1.0): %v\n", uint(1.0))
	fmt.Printf("int8(-123): %v\n", int8(-123))
	fmt.Printf("int16(6 + 0i): %v\n", int16(6+0i))
	fmt.Printf("complex128(789): %v\n", complex128(789))
	fmt.Printf("string(65): %v\n", string(65))                   // "A"
	fmt.Printf("string('A'): %v\n", string('A'))                 // "A"
	fmt.Printf("string('\\u68ee'): %v\n", string('\u68ee'))      // "森"
	fmt.Printf("string(-1): %v\n", string(-1))                   // "\uFFFD"
	fmt.Printf("string(0xFFFD): %v\n", string(0xFFFD))           // "\uFFFD"
	fmt.Printf("string(0x2FFFFFFFF): %v\n", string(0x2FFFFFFFF)) // "\uFFFD"

	// 下面是非法转换
	// int(1.23) // 1.23不能被表示为int类型值。
	// uint8(-1) // -1不能被表示为uint8类型值。
	// float64(1+2i) // 1+2i不能被表示为float64类型值。
	// // -1e+1000不能被表示为float64类型值。不允许溢出。
	// float64(-1e1000)
	// // 0x10000000000000000做为int值将溢出。
	// int(0x10000000000000000)
	// // 字面量65.0的默认类型是float64（不是一个整数类型）。
	// string(65.0)
	// // 66+0i的默认类型是complex128（不是一个整数类型）。
	// string(66+0i)
}
