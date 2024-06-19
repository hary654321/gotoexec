/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-06-19 11:14:31
 * @LastEditors: lhl
 * @LastEditTime: 2024-06-19 11:34:51
 */
package middlewares

import (
	"gotoexec/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		token := c.GetHeader("Authorization")
		//slog.Printf(slog.INFO, "Authorization %s  %s", token, global.ServerSetting.BasicAuth)
		if token != config.CoreConf.BasicAuth {
			code = 401
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "验签不通过",
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
