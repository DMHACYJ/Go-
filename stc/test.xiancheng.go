package main

import (
	"fmt"
)
func test_goroute(a int, b int) {
	sum := a + b//这里 := 自动把sum初始化成int（自动加上类型）
	fmt.Println(sum)
}

func test_print(a int) {
	fmt.Println(a)
}