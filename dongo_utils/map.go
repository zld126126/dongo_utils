package dongo_utils

import (
	"fmt"
	"sync"
)

// 键值对
var m map[interface{}]interface{}
var mu sync.RWMutex

func init() {
	m = make(map[interface{}]interface{})
}

func Put(key, value string) {
	mu.Lock()
	m[key] = value
	mu.Unlock()
	fmt.Println(fmt.Sprintf("key:%v value:%v option:%v", key, value, "create"))
}

func Delete(key string) {
	mu.Lock()
	delete(m, key)
	mu.Unlock()
	fmt.Println(fmt.Sprintf("key:%v: option:%v", key, "delete"))
}

func Get(key string) interface{} {
	mu.Lock()
	val, ok := m[key]
	mu.Unlock()
	if ok {
		return fmt.Sprintf("key:%v: value:%v", key, val)
	} else {
		return fmt.Sprintf("key:%v: option:%v", key, "not exist")
	}
}
