package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 函数外部定义变量必须使用var关键词，不可以使用:=。
// 函数外部定义变量作用域为包内部
//var aa = 3
//var ss = "kkk"
//var bb = true

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	// 创建不赋初值
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	// 创建赋初值
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	// 自动推断变量类型
	var a, b, c, d = 3, 4, true, "def"
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	// 第一次定义变量可以使用 := 更简单
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	// 欧拉公式
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))
	//e := cmplx.Pow(math.E, 1i*math.Pi) + 1
	e := cmplx.Exp(1i*math.Pi) + 1
	fmt.Printf("%.3f\n", e)
}

func triangle() {
	// 强制类型转换
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
	euler()
	triangle()
}
