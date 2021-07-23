package tests_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ahmetb/go-linq/v3"
)

func TestLinq(t *testing.T) {
	var ids []int
	for i := 1; i < 150000; i++ {
		ids = append(ids, i)
	}
	defaultFor(ids)
	linqFor(ids)
}

func defaultFor(ids []int) {
	startTime := time.Now().UnixNano() / 1e6
	for _, id := range ids {
		fmt.Println(id)
	}
	endTime := time.Now().UnixNano() / 1e6
	fmt.Println("defaultFor总耗时：", endTime-startTime)
}

func linqFor(ids []int) {
	startTime := time.Now().UnixNano() / 1e6
	linq.From(ids).ForEachT(func(id int) {
		fmt.Println(id)
	})
	endTime := time.Now().UnixNano() / 1e6
	fmt.Println("linqFor总耗时：", endTime-startTime)
}
