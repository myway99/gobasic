package main

import (
	"regexp"
	"fmt"
)

const text  = `
My email is ccmouse@gmail.com
email1 is abc@def.io
email2 is myway@hero.org
email3 is        xwang@qq.com.cn
`

func main() {

	re := regexp.MustCompile(
		`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	// for all: -1
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}









}
