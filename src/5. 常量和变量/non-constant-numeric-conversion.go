package main

const a = -1.23

// 变量b的类型被推断为内置类型float64。
var b = a

// error: 常量1.23不能被截断舍入到一个整数。(非常量可以)
var x = int32(a)

// error: float64类型值不能被隐式转换到int32。
var y int32 = b

// ok: z == -1，变量z的类型被推断为int32。
// z的小数部分将被舍弃。
var z = int32(b)

const k int16 = 255

var n = k            // 变量n的类型将被推断为int16。
var f = uint8(k + 1) // error: 常量256溢出了uint8。
var g uint8 = n + 1  // error: int16值不能隐式转换为uint8。
var h = uint8(n + 1) // ok: h == 0，变量h的类型为uint8。
// (n+1)溢出uint8，所以只有低8位
// bits（都为0）被保留。
