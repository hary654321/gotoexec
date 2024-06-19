/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-06-13 10:25:11
 * @LastEditors: lhl
 * @LastEditTime: 2024-06-19 11:49:33
 */
package main

import (
	"fmt"
	"gotoexec/config"
	"gotoexec/initialize"
	"gotoexec/middlewares"
	"gotoexec/util"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	util.Banner()
	var (
		err    error
		Router *gin.Engine
	)

	//加载配置
	config.Init("conf.toml")

	//加载路由
	Router = initialize.Routers()
	if config.CoreConf.HttpsServer {
		Router.Use(middlewares.TlsHandler())
		if err = Router.RunTLS(fmt.Sprintf(":%d", config.CoreConf.ApiPort), "pem/.cert.pem", "pem/.key.pem"); err != nil {
			log.Fatalln("Router.RunTLS" + err.Error())
		}
	} else {
		if err = Router.Run(fmt.Sprintf(":%d", config.CoreConf.ApiPort)); err != nil {
			log.Fatalln("Router.Run" + err.Error())
		}
	}

}
