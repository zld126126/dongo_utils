package dongo_utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

// 获取上一级文件夹路径
func GetParentDir(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

// 获取当前文件夹路径
func GetCurrentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

//判断文件或文件夹是否存在
func CheckExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

// 删除无用文件
func DelFile(s string) {
	pwd, _ := os.Getwd()
	m := make(map[string]string)
	filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
		fileName := path + info.Name()
		if strings.Contains(fileName, s) {
			m[info.Name()] = path
		}
		return nil
	})

	for k, v := range m {
		err := os.Remove(v)
		if err != nil {
			fmt.Println(k + "删除失败")
		} else {
			fmt.Println(k + "删除成功")
		}
	}
}

// 复制文件
func CopyFile(source, target string) bool {
	src, err := os.Open(source)
	if err != nil {
		fmt.Println("复制文件失败")
		return false
	}
	defer src.Close()

	dst, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("复制文件失败")
		return false
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		fmt.Println("复制文件失败")
		return false
	}
	return true
}
