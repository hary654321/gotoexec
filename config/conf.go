/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-06-19 10:28:12
 * @LastEditors: lhl
 * @LastEditTime: 2024-06-19 10:29:43
 */
package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	// CoreConf crocodile conf
	CoreConf *coreConf
)

// Init Config
func Init(conf string) {
	_, err := toml.DecodeFile(conf, &CoreConf)
	if err != nil {
		fmt.Printf("Err %v", err)
		os.Exit(1)
	}
}

type coreConf struct {
	ApiPort     int    `json:"apiPort"`
	ListenPort  int    `json:"listenPort"`
	HttpsServer bool   `json:"https"`
	BasicAuth   string `json:"basicAuth"`
	Version     string
	PortCont    int
}
