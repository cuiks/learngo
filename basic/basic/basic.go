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
	var c = calcTriangle(a, b)
	fmt.Println(c)
}

func calcTriangle(a, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}

func consts() {
	// 定义常量 go的常量不会全部大写。
	// const数值可作为各种类型使用

	//const filename = "abc.txt"
	//const a, b = 3, 4
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	// 枚举类型
	//const (
	//	cpp    = 0
	//	java   = 1
	//	python = 2
	//	golang = 3
	//)
	//const (
	//	cpp = iota
	//	java
	//	python
	//	golang
	//)
	const (
		cpp = iota
		_
		python
		golang
		JavaScript
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, JavaScript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)

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
	consts()
	enums()
}
