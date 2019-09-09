package main

import(
	a "go_dev/day2/example2/add"//给add包取别名a
	"fmt"
)

func main() {
	a.Test()
	fmt.Println("Name=", a.Name)
	fmt.Println("Age=", a.Age)
}

