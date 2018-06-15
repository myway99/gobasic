package main

import (
	"fmt"
)

//[]int：表示数组的切片
//[5]int：表示长度为5的数组
func printArray(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}

}

func main() {
	//
	var arr1 [5]int
	//
	arr2 := [3]int{1, 3, 5}
	//让编译器帮忙计数
	arr3 := [...]int{2, 4, 6, 8, 10}

	//二维数组
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	//遍历数组1：普通方法
	//for i := 0; i < len(arr3); i++ {
	//	fmt.Println(arr3[i])

	//遍历数组2: 可通过range获取array的下标
	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	//遍历数组3
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	//遍历数组4: 不要下标，只要值
	for _, v := range arr3 {
		fmt.Println(v)
	}

	//函数内改变
	fmt.Println("printArray(arr1)")
	printArray(&arr1)

	//函数内改变
	fmt.Println("printArray(arr3)")
	printArray(&arr3)

	//函数外没变
	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)

}
