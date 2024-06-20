/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-06-19 11:14:31
 * @LastEditors: lhl
 * @LastEditTime: 2024-06-19 18:23:23
 */

package router

import (
	"gotoexec/global"
	"gotoexec/grpcapi"
	"gotoexec/middlewares"
	"gotoexec/server/control"
	"gotoexec/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitControlRouter(Router *gin.RouterGroup) {
	p := Router.Group("control").Use(middlewares.CostTime()).Use(middlewares.BasicAuth())
	{
		p.POST("", func(context *gin.Context) {

			cmd := context.PostForm("cmd")
			ip := context.PostForm("ip")
			gcmd := new(grpcapi.Command)
			gcmd.In, _ = util.EncryptByAes([]byte(cmd))
			gcmd.Ip = ip

			cmdout, err := control.ControlInstance.RunCommand(gcmd)

			if err != nil {
				log.Panicln(err)
			}

			out, err := util.DecryptByAes(cmdout.Out)

			if err != nil {
				log.Panicln(err)
			}

			context.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"data": string(out),
				"msg":  "health",
			})
		})

		p.GET("", func(context *gin.Context) {

			context.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"data": global.FixedSizeStackInstance.Get(),
				"msg":  "health",
			})
		})

	}
}
