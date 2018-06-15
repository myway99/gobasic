package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string {
	result := ""
	//省略起始条件
	for ; n > 0; n /= 2 {
		//将取模作为最低位lsb
		lsb := n % 2
		//转换为字符串后再加
		result = strconv.Itoa(lsb) + result
	}
	return result

}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	//省略递增条件
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//死循环，goroutine经常会用到
func forever() {
	for {
		fmt.Println("abc")
	}

}

func main() {
	fmt.Println(
		convertToBin(5),  //101
		convertToBin(13), //先对2取模：1011 --> 再转换位置1101
		convertToBin(34567913))
	fmt.Printf("%q\n", convertToBin(0))

	printFile("basic/branch/abc.txt")
	s := `abc"d"
	kkkk
	3456

	p`
	printFileContents(strings.NewReader(s))

	//forever()

}
