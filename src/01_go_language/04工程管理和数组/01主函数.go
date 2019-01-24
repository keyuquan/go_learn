package main

import "01_go_language/04工程管理和数组/userinfo"

//全局变量
var num int = 123

//在同级别目录报名要相同
func main() {
	add(10, 20)
	//包名.函数名
	userinfo.Login()
	userinfo.DeleteUser()
	userinfo.SelectUser()

}
