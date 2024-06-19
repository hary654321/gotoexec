/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-06-19 10:34:05
 * @LastEditors: lhl
 * @LastEditTime: 2024-06-19 10:36:13
 */
package router

import (
	"gotoexec/middlewares"
	"gotoexec/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitHealthRouter(Router *gin.RouterGroup) {
	c := Router.Group("check").Use(middlewares.CostTime()).Use(middlewares.BasicAuth())
	{

		c.GET("/heartbeat", func(c *gin.Context) {
			data := make(map[string]interface{})
			data["time"] = util.GetTime()

			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "",
				"data": data,
			})
		})

	}
}
