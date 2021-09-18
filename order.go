package dongo_utils

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成订单号 日期20191025时间戳1571987125435+3位随机数
func GenerateCode() string {
	date := ParseTimeToYearMonthDay(time.Now())
	r := rand.Intn(1000)
	code := fmt.Sprintf("%s%d%03d", date, Tick64(), r)
	fmt.Println(code + " rand ID generate successed!\n")
	return code
}

// 抽奖
func Lottery(a int) string {
	start := 0
	var end int
	//probabilities，一共几个概率事件，另外各自概率是多少 必须相加=10000
	probabilities := []int{300, 1000, 2400, 1600, 3000, 1200, 500}
	rand := rand.Intn(10000) //0-10000的随机数共10000个数
	for _, probability := range probabilities {
		end += probability
		if start <= rand && end > rand {
			return fmt.Sprintf("第%d次,开始:%d,随机数:%d,结束:%d", a, start, rand, end)
		}
		start = end
	}
	return "概率错误"
}
