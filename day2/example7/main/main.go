package main
//获取当前运行的操作系统名称和PATH环境变量的值，并打印在中端。
import(
	"fmt"
	"os"
)

func main() {
	var goos string = os.Getenv("GOOS")
	fmt.Printf("The operating system is: %s\n",goos)
	p := os.Getenv("PATH")
	fmt.Printf("Path is %s\n",p)
}