package main

import (
	"github.com/skip2/go-qrcode"
	"image/color"
	"log"
)

func main() {
	qrCodeForSite()
	customQrCode()
}

func qrCodeForSite() {
	qrcode.WriteFile("http://c.biancheng.net/", qrcode.Medium, 256, "7_package/site_qrcode.png")
}

func customQrCode() {
	qr,err:=qrcode.New("http://c.biancheng.net/", qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	} else {
		qr.BackgroundColor = color.RGBA{50,205,50,255}
		qr.ForegroundColor = color.White
		qr.WriteFile(256, "7_package/custom_qrcode.png")
	}
}