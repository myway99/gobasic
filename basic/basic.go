package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//定义包内变量（Go语言没全局变量的说法）

//函数外每一行开始都须有关键字
//函数外不能用简短声明：如 bb := true
//var aa = 3
//var ss "kkk"
//var bb = true

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	//type  name
	var a int
	var s string
	//打印,后跟变量名
	fmt.Println(a, s)
	//按格式打印
	fmt.Printf("%d %s\n", a, s)
	//%q：打印"号
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, s, b)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, s, c, b)
}

func variableShorter() {
	//函数内可用简短声明
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, s, c, b)
}

func myeuler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}

func euler() {
	//fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

//Go须强制类型转换
func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

//const filename  = "abc.txt"
func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)

}

//使用常数定义枚举类型
func enums() {
	const (
		//使用iota自增值
		cpp = iota
		//java
		_
		python
		golang
		javascript
	)
	fmt.Println(cpp, javascript, python, golang)

	//b,kb,mb,gb,tb,pb
	const (
		//依次左移10位
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world!")

	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	println(aa, ss, bb)

	myeuler()
	euler()
	triangle()

	consts()
	enums()
}
