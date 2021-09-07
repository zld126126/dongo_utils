package dongo_utils

import (
	"fmt"
)

// 打印命令
func PrintCommand(command string) {
	fmt.Println("========================")
	fmt.Println(command)
	fmt.Println("========================")
}

// 参考CMD语法 等待用户随机指令关闭程序
func Pause() {
	var str string
	fmt.Println("")
	fmt.Print("请按任意键继续...")
	fmt.Scanln(&str)
	fmt.Print("程序退出...")
}
