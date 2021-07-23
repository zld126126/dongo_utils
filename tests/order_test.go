package tests_test

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestAlgorithm(t *testing.T) {
	//100万次算法运行
	for a := 0; a < 1000000; a++ {
		t.Log(Algorithm(a))
	}

	//抽奖算法2
	prize := newLottery()
	for i := 0; i < 10; i++ {
		item := Lottery(prize)
		fmt.Println("抽奖结果:", item)
	}
}

// 抽奖算法1
func Algorithm(a int) string {
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

// 抽奖算法2
type PrizeCount int

const (
	PrizeCountWithout  PrizeCount = iota // 没有奖品数量
	PrizeCountLeft                       // 还有奖品数量
	PrizeCountInfinite                   // 无限奖品数量
)

type PrizeType int

const (
	PrizeTypeEmpty  PrizeType = iota // 没有奖品
	PrizeTypeCoupon                  // 优惠券奖品
	PrizeTypeGoods                   // 实物奖品
	PrizeTypeCoin                    // 金币奖品
	PrizeTypeMoney                   // 金钱奖品
)

type Prize struct {
	PrizeType  PrizeType  `json:"prize_type"`  // 奖品类型
	PrizeCount PrizeCount `json:"prize_count"` // 奖品数量
}

func newLottery() []*Prize {
	empty := &Prize{
		PrizeType:  PrizeTypeEmpty,
		PrizeCount: PrizeCountInfinite,
	}

	coupon := &Prize{
		PrizeType:  PrizeTypeCoupon,
		PrizeCount: PrizeCountInfinite,
	}

	goods := &Prize{
		PrizeType:  PrizeTypeGoods,
		PrizeCount: PrizeCountLeft,
	}

	coin := &Prize{
		PrizeType:  PrizeTypeCoin,
		PrizeCount: PrizeCountInfinite,
	}

	money := &Prize{
		PrizeType:  PrizeTypeMoney,
		PrizeCount: PrizeCountWithout,
	}

	prize := []*Prize{empty, coupon, goods, coin, money}
	var result []*Prize
	for _, item := range prize {
		if item.PrizeCount == PrizeCountWithout {
			continue
		}
		result = append(result, item)
	}
	return result
}

func Lottery(prize []*Prize) *Prize {
	prizeLen := len(prize)
	if prizeLen == 0 {
		return nil
	}
	randPrize := rand.Intn(prizeLen - 1)
	fmt.Println("随机数:", randPrize)
	return prize[randPrize]
}
