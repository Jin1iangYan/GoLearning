package main

const y = 70

var x int = 123 // 包级变量

func main() {
	// 此x变量遮挡了包级变量x。
	var x = true

	// 一个内嵌代码块。
	{
		x, y := x, y-10 // 这里，左边的x和y均为新声明
		// 的变量。右边的x为外层声明的
		// bool变量。右边的y为包级变量。

		// 在此内层代码块中，从此开始，
		// 刚声明的x和y将遮挡外层声明x和y。

		x, z := !x, y/10 // z是一个新声明的变量。
		// x和y是上一句中声明的变量。
		println(x, y, z) // false 60 6
	}
	println(x) // true
	println(y) // 70 （包级变量y从未修改）
	/*
	   println(z) // error: z未定义。
	   // z的作用域仅限于上面的最内层代码块。
	*/
}
