package dongo_utils

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
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

//获取int64 获取毫秒
func Tick(t ...time.Time) int64 {
	if len(t) == 0 {
		return time.Now().UnixNano() / 1e6
	} else {
		return t[0].UnixNano() / 1e6
	}
}

// 转换json
func ToJson(i interface{}) string {
	data, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	return string(data)
}
