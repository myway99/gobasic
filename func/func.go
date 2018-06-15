package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//return a / b
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}

//返回多个值：13/3=4...1
//常用于返回错误
func div(a, b int) (q, r int) {
	//q = a / b
	//r = a % b
	//return
	return a / b, a % b
}

//函数op接收两个参数，返回一个参数;
//用收到的a和b两个参数去apply函数op，然后将函数op的结果作为函数apply的返回结果int
func apply(op func(int, int) int, a, b int) int {
	//通过reflect包来进行反射获取函数op的真正的指针
	p := reflect.ValueOf(op).Pointer()
	//通过runtime的FuncForPC将指针传进去，从而获得函数op的名字
	opName := runtime.FuncForPC(p).Name()
	//打印函数op的名字
	fmt.Printf("Calling function %s with args"+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

//求a的b次方
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//使用可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

//无法交换位置
func swap1(a, b int) {
	b, a = a, b

}

//交换位置：值传递
func swap2(a, b *int) {
	//b,a = a,b
	*b, *a = *a, *b
}

//交换位置：
func swap3(a, b int) (int, int) {
	return b, a
}

func main() {

	//fmt.Println(eval(131,32,"x"))
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	q, r := div(131, 32)
	fmt.Println(q, r)

	//fmt.Println(apply(pow,3,4))  //调用自定义函数

	//使用匿名函数
	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(
				float64(a), float64(b)))
		}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))

	a, b := 3, 4

	c, d := swap3(a, b)
	fmt.Println(c, d)

	swap1(a, b)
	fmt.Println(a, b)

	//swap(a,b)
	swap2(&a, &b)
	fmt.Println(a, b)

}
