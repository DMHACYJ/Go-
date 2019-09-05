package main

import (
	"fmt"
)

func calc(a int, b int) (int, int) {//两个以上的参数返回值一定要用括号括起来
	c := a + b
	d := (a + b) / 2
	return c, d//返回值用逗号隔开
}

func main() {
	sum, avg := calc(100, 200)//如果要湖绿掉一个返回值  sum, _ := calc(100, 200)
	fmt.Println("sum=", sum, "avg=", avg)
}