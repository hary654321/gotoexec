/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-06-13 10:25:11
 * @LastEditors: lhl
 * @LastEditTime: 2024-06-20 18:19:11
 */
package util

import (
	"bytes"
	"fmt"
	"image"
	"image/png"

	"github.com/kbinani/screenshot"
)

func Screenshot() []*image.RGBA {
	var images []*image.RGBA
	//获取当前活动屏幕数量
	i := screenshot.NumActiveDisplays()
	if i == 0 {
		//linux 服务器可能没有屏幕
	}
	for j := 0; j <= i-1; j++ {
		image, _ := screenshot.CaptureDisplay(j)
		images = append(images, image)
	}
	return images
}

func ImageToByte(image *image.RGBA) []byte {

	if image == nil {
		return nil // Or handle the nil case as appropriate for your application
	}

	buf := new(bytes.Buffer)
	err := png.Encode(buf, image)

	if err != nil {
		// Handle the error, e.g., log it
		fmt.Printf("Error encoding PNG: %v\n", err)
		return nil // Return nil or an appropriate error indicator
	}

	return buf.Bytes()
}
