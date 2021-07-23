package tests_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/zld126126/dongo_utils/dongo_utils"
)

func TestMemory(t *testing.T) {
	t1 := newMemory("dongtech")
	t1.Put("haha")
	t1.Put("lala")
	t1.Put("haha")
	t1.Get()
	t1.Length()
	time.Sleep(time.Second * 4)
	fmt.Println(time.Now().Unix())
	if nil == t1.Get() {
		t1.Put("haha")
	}
	t1.Length()
}

func newMemory(name string) *dongo_utils.Memory {
	t := &dongo_utils.Memory{
		Name: name,
	}
	go func() {
		for {
			t.Repair(time.Now().Unix())
			time.Sleep(time.Second)
		}
	}()
	return t
}
