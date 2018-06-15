package main

import (
	"time"
	"fmt"
)

func main() {
	//var a [10]int
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				//a[i]++
				//// 交出控制权
				//runtime.Gosched()
				fmt.Printf("Hello from" +
					"goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
	//fmt.Println(a)
}
