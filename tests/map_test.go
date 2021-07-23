package tests_test

import (
	"testing"

	"github.com/zld126126/dongo_utils/dongo_utils"
)

func TestMap(t *testing.T) {
	dongo_utils.Put("name", "zhangld")
	dongo_utils.Put("age", "123")
	dongo_utils.Put("name", "zhangld")
	dongo_utils.Delete("name")
	t.Log(dongo_utils.Get("age"))
	t.Log(dongo_utils.Get("name"))
}
