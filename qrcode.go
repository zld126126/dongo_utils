package dongo_utils

import (
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/tuotoo/qrcode"
)

// 创建二维码
func CreateQrCode(fileName, content string) {
	qrCode, _ := qr.Encode(content, qr.M, qr.Auto)

	qrCode, _ = barcode.Scale(qrCode, 256, 256)

	file, _ := os.Create(fileName)
	defer file.Close()

	png.Encode(file, qrCode)
}

// 加载二维码
func LoadQrCode(fileName, content string) {
	fi, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fi.Close()
	matrix, err := qrcode.Decode(fi)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("contrast result:", matrix.Content == content)
	fmt.Println(matrix.Content)
}
