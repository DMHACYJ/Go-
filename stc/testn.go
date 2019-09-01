package main

import (
	"fmt"
	"time"
)

func add(a int, b int) int {
	var sum int 
	sum = a + b
	return sum
}

func main() {
	//var c int
	//c = add(100, 200)
	//go test_gotoute(300, 300)
	//fmt.Println("add(100,200)=", c)

	for i := 0; i < 100; i++ { //起100个线程
		go test_print(i)
	}
	time.Sleep(time.second)
}