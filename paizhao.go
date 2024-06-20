package main

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/disintegration/gift"
)

func main() {
	// 打开摄像头
	f, err := hub.Open("0")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 创建一个GIFT操作，这里我们不做任何处理，直接捕获图像
	g := gift.New(gift.Brightness(0))

	// 创建一个图像 mosaic 用于显示视频流
	mosaic := image.NewNRGBA(image.Rect(0, 0, 100, 100))

	for {
		// 从摄像头读取一帧
		frame, err := f.Frame()
		if err != nil {
			panic(err)
		}

		// 将帧转换为gift.Image
		gimg := gift.Image(frame)

		// 应用GIFT操作
		g.Draw(gimg, gimg.Bounds())

		// 将处理后的帧放入mosaic图像中
		mosaic.SubImage(mosaic.Rect).(*image.NRGBA).Image = gimg

		// 保存帧到jpeg文件
		jpeg.Encode(os.Stdout, mosaic, &jpeg.Options{Quality: 95})

		// 如果需要截图，可以在这里保存图像到文件
		// out, _ := os.Create("output.jpg")
		// jpeg.Encode(out, mosaic, &jpeg.Options{Quality: 95})
		// out.Close()
	}
}
