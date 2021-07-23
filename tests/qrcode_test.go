package tests_test

import (
	"testing"

	"github.com/zld126126/dongo_utils/dongo_utils"
)

func TestQrcode(t *testing.T) {
	fileName := "qr.png"
	content := "https://www.baidu.com"
	dongo_utils.CreateQrCode(fileName, content)
	dongo_utils.LoadQrCode(fileName, content)
}
