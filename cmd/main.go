package main

import (
	"bluebell/config"
	"bluebell/internal/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "project config file path.")
)

func main() {

	// 初始化配置文件
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// 初始化数据库
	config.DB.InitDb()
	defer config.DB.Close()

	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	router.Load(
		g,
	)

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}
