package utils

import (
	"fmt"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
	"image"
	"image/draw"
	"image/png"
	"os"
)

func CreateCompanyQR(url string, qrId string) string {
	bgImg, _ := generaQrCode(url)
	logoFile, _ := os.Open("./public/logo.png")
	logoImg, _ := png.Decode(logoFile)
	logoImg = ImageResize(logoImg, 80, 80)
	b := bgImg.Bounds()
	// 设置为居中
	offset := image.Pt((b.Max.X-logoImg.Bounds().Max.X)/2, (b.Max.Y-logoImg.Bounds().Max.Y)/2)
	m := image.NewRGBA(b)
	draw.Draw(m, b, bgImg, image.Point{X: 0, Y: 0}, draw.Src)
	draw.Draw(m, logoImg.Bounds().Add(offset), logoImg, image.Point{X: 0, Y: 0}, draw.Over)

	QRFileName := fmt.Sprintf("./QrCode/%v.png", qrId)
	// 上传至oss时这段要改
	i, _ := os.Create(QRFileName)
	_ = png.Encode(i, m)
	defer i.Close()
	return QRFileName
}

func generaQrCode(url string) (img image.Image, err error) {
	qrCode, _ := qrcode.New(url, qrcode.Medium)
	qrCode.DisableBorder = true
	img = qrCode.Image(300)
	return img, nil
}

func CreateQrCode(url string, qrId string) string {
	QRFileName := fmt.Sprintf("./QrCode/%v.png", qrId)
	_ = qrcode.WriteFile(url, qrcode.Medium, 256, QRFileName)
	return QRFileName
}

func ImageResize(src image.Image, w, h int) image.Image {
	return resize.Resize(uint(w), uint(h), src, resize.Lanczos3)
}
