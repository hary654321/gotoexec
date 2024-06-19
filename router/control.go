/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-06-19 10:34:05
 * @LastEditors: lhl
 * @LastEditTime: 2024-06-19 10:43:46
 */
package router

import (
	"gotoexec/middlewares"

	"github.com/gin-gonic/gin"
)

func InitPortRouter(Router *gin.RouterGroup) {
	p := Router.Group("port").Use(middlewares.CostTime()).Use(middlewares.BasicAuth())
	{
		p.POST("/start", port.Start)
		p.POST("/progress", port.Progress)
		p.POST("/result", port.Res)
		p.POST("/stop", port.Stop)
	}
}
