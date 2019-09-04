package main

import(
	"fmt"
)

func test_pipe() {
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2//管道放数据
	pipe <- 3

	var t1 int
	t1 =<- pipe//把1取出，取出方法类似于队列，先进先出

	pipe <- 4

	fmt.Println(t1)
	fmt.Println(len(pipe))
}
func main() {
	test_pipe()
}