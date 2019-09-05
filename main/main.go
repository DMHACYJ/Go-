package main

import(
	"go_dev/day1/package_example/calc"
	"fmt"
)

func main() {
	sum := calc.ADD(100, 300)
	sub := calc.SUB(100, 300)

	fmt.Println("sum=", sum)
	fmt.Println("sub=", sub)
}