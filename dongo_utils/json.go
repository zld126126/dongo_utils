package dongo_utils

import (
	"encoding/json"
	"fmt"
)

// 打印json
func PrintJson(any ...interface{}) {
	for _, obj := range any {
		b, err := json.Marshal(obj)
		fmt.Println(err, string(b))
	}
}

// 转换json
func ToJsonString(v interface{}) (string, error) {
	bs, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
