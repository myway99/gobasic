package main

import (
	"../../functional/fib"
	"bufio"
	"errors"
	"fmt"
	"os"
)

func tryDefer() {

	for i := 0; i < 100; i++ {

		// 倒过来输出
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}

	// 栈：先进后出
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(3)
	//panic("error occured")
	//return
	//fmt.Println(4)

}

func writeFile(filename string) {
	file, err := os.OpenFile(
		filename, os.O_EXCL|os.O_CREATE, 0666)

	err = errors.New("this is a custom error")
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		//fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	//由内存写入文件
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	writeFile("fib.txt")

}
