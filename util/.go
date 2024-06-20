/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-06-13 10:25:11
 * @LastEditors: lhl
 * @LastEditTime: 2024-06-20 17:05:28
 */
package util

import (
	"bytes"
	"image"
	"image/png"
	"log"

	"github.com/atotto/clipboard"
	"github.com/kbinani/screenshot"
)

func Screenshot() []*image.RGBA {
	var images []*image.RGBA
	//获取当前活动屏幕数量
	i := screenshot.NumActiveDisplays()
	if i == 0 {

	}
	for j := 0; j <= i-1; j++ {
		image, _ := screenshot.CaptureDisplay(j)
		images = append(images, image)
	}
	return images
}

func ImageToByte(image *image.RGBA) []byte {
	buf := new(bytes.Buffer)
	png.Encode(buf, image)
	b := buf.Bytes()
	return b
}

func Clipboard() string {
	text, err := clipboard.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Clipboard content:", text)

	return text
}
