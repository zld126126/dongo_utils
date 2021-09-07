package dongo_utils

// 键值对
var PariMap = make(map[string]interface{})

func ParisMap_Get(key string) interface{} {
	v, exist := PariMap[key]
	if exist {
		return v
	} else {
		return nil
	}
}

func ParisMap_Put(key string, value interface{}) {
	PariMap[key] = value
}

func ParisMap_Del(key string) {
	PariMap[key] = nil
}

func ParisMap_DelAll() {
	PariMap = make(map[string]interface{})
}
