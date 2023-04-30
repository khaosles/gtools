package gcaptcha

import (
	"image/color"
	"math/rand"

	"github.com/mojocn/base64Captcha"
)

/*
   @File: captcha.go
   @Author: khaosles
   @Time: 2023/3/9 22:29
   @Desc:
*/

var baseStore = base64Captcha.DefaultMemStore

func GetCaptcha() (string, string) {
	var driver base64Captcha.Driver
	driver = stringConfig()
	c := base64Captcha.NewCaptcha(driver, baseStore)
	id, b64s, err := c.Generate()
	if err != nil {
		return GetCaptcha()
	}
	return id, b64s
}

func Verify(id, value string) bool {
	return baseStore.Verify(id, value, true)
}

func GetAnswer(id string) string {
	return baseStore.Get(id, true)
}

func digitConfig() *base64Captcha.DriverDigit {
	digitType := &base64Captcha.DriverDigit{
		Height:   40,
		Width:    90,
		Length:   4,
		MaxSkew:  0.45,
		DotCount: 80,
	}
	return digitType
}

func stringConfig() *base64Captcha.DriverString {
	stringType := &base64Captcha.DriverString{
		Height:          40,
		Width:           90,
		NoiseCount:      3,
		ShowLineOptions: 0,
		Length:          4,
		Source:          base64Captcha.TxtAlphabet + base64Captcha.TxtNumbers,
		BgColor: &color.RGBA{
			R: uint8(rand.Intn(40) + 180),
			G: uint8(rand.Intn(20) + 200),
			B: uint8(rand.Intn(40) + 180),
			A: uint8(rand.Intn(40) + 180),
		},
		Fonts: nil,
	}
	return stringType
}
