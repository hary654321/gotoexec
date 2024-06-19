/*
 * @Description:InitControlRouter
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-06-19 11:14:31
 * @LastEditors: lhl
 * @LastEditTime: 2024-06-19 11:40:55
 */
package initialize

import (
	"gotoexec/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	ApiGroup := Router.Group("/v1")
	router.InitControlRouter(ApiGroup)
	router.InitHealthRouter(ApiGroup)
	return Router
}
