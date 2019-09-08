package main

import(
	"time"
)

func main() {

	for i := 0; i < 100; i++ {
		go test_goroute(i)//起一个goroute和主线程并发执行
	}

	time.Sleep(time.Second)
}