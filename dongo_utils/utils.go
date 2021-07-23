package dongo_utils

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
)

// 检查错误
func CheckError(e error, msg string) {
	if e != nil {
		fmt.Println(msg + "," + e.Error())
		log.Fatal(e)
	}
}

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

// 转换json
func ToJson(i interface{}) string {
	data, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	return string(data)
}

// 正则检查
func RegexpString(token string) bool {
	matched, err := regexp.MatchString(`^[A-Za-z0-9]{64}$`, token)
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		fmt.Println(token, " result :", matched)
		return matched
	}
}
