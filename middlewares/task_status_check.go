/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-06-19 11:14:31
 * @LastEditors: lhl
 * @LastEditTime: 2024-06-19 11:36:17
 */
package middlewares

import (
	"time"

	"log"

	"github.com/gin-gonic/gin"
)

func CostTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		//请求前获取当前时间
		nowTime := time.Now()

		//请求处理
		c.Next()

		//处理后获取消耗时间
		costTime := time.Since(nowTime)
		//todo 当前url获取方法可能有用
		url := c.Request.URL.String()
		log.Printf("the request URL %s cost %v\n", url, costTime)
	}
}
