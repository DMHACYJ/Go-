package add

//var Name string = "hello world"//声明并初始化
//var Age int = 10

var Name string
var Age int 

func Test() {
	Name = "hello world"//
	Age = 10//这类执行语句必须放在函数内,
	//如果基于上面的main函数，主函数没有声明的话，默认输出：字符串初始值为空，整形初始值为0
}