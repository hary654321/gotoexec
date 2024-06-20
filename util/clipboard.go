/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-06-13 10:25:11
 * @LastEditors: lhl
 * @LastEditTime: 2024-06-20 17:22:29
 */
package util

import (
	"log"

	"github.com/atotto/clipboard"
)

func Clipboard() string {
	text, err := clipboard.ReadAll()
	if err != nil {
		//剪切板是文件的时候会报错
		log.Println("获取剪切板失败", err, text)
	}

	return text
}
