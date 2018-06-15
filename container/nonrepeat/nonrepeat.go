package main

import "fmt"

var lastOccurred = make([]int, 0xffff) //65535

func lengthOfNonRepeatingSubStr(s string) int {
	//使用make生成map：lastOccurred
	//lastOccurred := make(map[byte]int)
	//lastOccurred := make(map[rune]int)

	//lastOccurred[0x65] = 1
	//lastOccurred[0x8BFE] = 6
	for i := range lastOccurred {
		lastOccurred[i] = -1
	}
	start := 0
	maxLength := 0

	//使用range遍历待查字符串s
	//for i, ch := range []byte(s) {
	for i, ch := range []rune(s) {
		//检查是否已有字符，并比较判断
		// 若在start后（大于等于start），则修改start
		if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		//若不存在或在start前（小于start），则修改maxLength
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		//更新lastOccurred[ch]
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("acccbacb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("f"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdefghijklmnopqrstuvwxyz12378"))
	fmt.Println(lengthOfNonRepeatingSubStr("这里是地球这里是中华人民共和国这里是北京这里是myway"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二三四五六五四三二一"))

}
