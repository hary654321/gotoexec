/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-06-19 11:14:31
 * @LastEditors: lhl
 * @LastEditTime: 2024-08-05 15:08:45
 */

package router

import (
	"context"
	"encoding/base64"
	"gotoexec/global"
	"gotoexec/grpcapi"
	"gotoexec/middlewares"
	"gotoexec/server/control"
	"gotoexec/util"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func InitControlRouter(Router *gin.RouterGroup) {
	p := Router.Group("control").Use(middlewares.CostTime()).Use(middlewares.BasicAuth())
	{
		p.POST("", func(rctx *gin.Context) {

			cmd := rctx.PostForm("cmd")
			ip := rctx.PostForm("ip")
			gcmd := new(grpcapi.Command)
			gcmd.In, _ = util.EncryptByAes([]byte(cmd))
			gcmd.Ip = ip

			log.Println("gcmd", gcmd)

			baseCtx := context.Background()

			// 设置超时时间为5秒
			ctx, cancel := context.WithTimeout(baseCtx, 8*time.Second)
			defer cancel()

			cmdout, err := control.ControlInstance.RunCommandCtx(ctx, gcmd)

			if err != nil {
				log.Panicln(err)
			}

			if cmdout.Out == "off" {
				rctx.JSON(http.StatusOK, gin.H{
					"code": http.StatusBadRequest,
					"data": "不在线",
					"msg":  "off",
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

			log.Println("out:" + outstring)
			rctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"data": outstring,
				"msg":  "ok",
			})

			return
		})

		p.GET("", func(context *gin.Context) {

			context.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"data": global.FixedSizeStackInstance.Get(),
				"msg":  "health",
			})

			return
		})

	}
}
