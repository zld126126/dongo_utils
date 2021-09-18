package dongo_utils

import mapset "github.com/deckarep/golang-set"

//数组取出不同元素 放入结果
//sourceList中的元素不在sourceList2中 则取到result中
func GetDifferentStringArray(sourceList, sourceList2 []string) (result []string) {
	for _, src := range sourceList {
		var find bool
		for _, target := range sourceList2 {
			if src == target {
				find = true
				continue
			}
		}
		if !find {
			result = append(result, src)
		}
	}
	return
}

//合并两个字符串数组并去重
func MergeDuplicateStringArray(slice []string, elems []string) []string {
	listPId := append(slice, elems...)
	t := mapset.NewSet()
	for _, i := range listPId {
		t.Add(i)
	}
	var result []string
	for i := range t.Iterator().C {
		result = append(result, i.(string))
	}
	return result
}

//字符串数组去重
func DuplicateStringArray(m []string) []string {
	s := make([]string, 0)
	smap := make(map[string]string)
	for _, value := range m {
		//计算map长度
		length := len(smap)
		smap[value] = value
		//比较map长度, 如果map长度不相等， 说明key不存在
		if len(smap) != length {
			s = append(s, value)
		}
	}
	return s
}

//字符串数组存在某个字符串
func ExistStringArray(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
