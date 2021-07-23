package tests_test

import (
	"testing"
	"time"

	"github.com/zld126126/dongo_utils/dongo_utils"
)

func TestGetParentDir(t *testing.T) {
	s := dongo_utils.GetCurrentDir()
	t.Log(s)
}

func TestGenerateCode(t *testing.T) {
	s := dongo_utils.GenerateCode()
	t.Log(s)
}

func TestSome(t *testing.T) {
	a := []int{1, 2, 1, 3, 4}
	b := []int{4, 3, 5, 6}

	//多个数组合并去重
	c := dongo_utils.MergeDuplicateIntArray(a, b)
	t.Log(c)

	//单个数组去重
	d := dongo_utils.DuplicateIntArray(a)
	t.Log(d)

	//取不同
	e := dongo_utils.GetDifferentIntArray(a, b)
	t.Log(e)

	//判断数组有某个元素
	flag := dongo_utils.ExistIntArray(a, 1)
	t.Log(a, "存在", 1, flag)
	str1 := []string{"a", "b", "c", "a", "c"}
	str2 := []string{"b", "a", "d", "e", "f"}

	//多字符串数组去重合并
	str3 := dongo_utils.MergeDuplicateStringArray(str1, str2)
	t.Log(str3)

	//单个字符串数组去重
	str4 := dongo_utils.DuplicateStringArray(str1)
	t.Log(str4)

	//取不同
	str5 := dongo_utils.GetDifferentStringArray(str1, str2)
	t.Log(str5)

	//字符串数组判断存在某个字符串
	flag2 := dongo_utils.ExistStringArray(str1, "a")
	t.Log(str1, "存在", "a", flag2)

	//时间转换
	timeStr := "2019-01-02 15:04:05"
	parseTimeReason := "时间转换原因"
	t1, err := dongo_utils.ParseTimeByTimeStr(timeStr, parseTimeReason)
	if err != nil {
		t.Log(err.Error())
		dongo_utils.Catch(err)
	}
	t.Log(t1)

	//获取int64的时间戳
	time1 := dongo_utils.Tick()
	t.Log("当前时间:", time1)
	time2 := dongo_utils.Tick(t1)
	t.Log("2019-01-02 15:04:05转换时间:", time2)

	now := time.Now()
	//获取最近的周一
	time3 := dongo_utils.ParseCurrentMonday(now)
	t.Log(time3)
	//返回某一天的当地时区0点
	time4 := dongo_utils.ParseMorningTime(now)
	t.Log(time4)
	//当月第一天0点
	time5 := dongo_utils.ParseFirstDayOfMonthMorning(now)
	t.Log(time5)
	//获取传入时间前一天的时间，不传默认是昨天
	time6 := dongo_utils.ParseYesterdayTime(now)
	t.Log(time6)
	//把int64转换成1993-12-26 10:30:00
	time7 := dongo_utils.Tick(time6)
	timeStr2 := dongo_utils.ParseTimeToString(time7)
	t.Log(timeStr2)
}
