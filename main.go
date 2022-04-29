package main

import (
	"GinProject/dependency"
	"GinProject/global"
	"GinProject/initialize"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	global.Vp = dependency.Viper()
	global.Log = dependency.Zap()
	zap.ReplaceGlobals(global.Log)
	global.Log.Info("zap日志初始化成功")
	global.Db = initialize.Gorm()
	router := gin.Default()
	// Listen and Server in 0.0.0.0:8080
	port := fmt.Sprintf(":%d", global.SystemConfig.System.Port)
	err := router.Run(port)
	if err != nil {
		global.Log.Error("启动异常", zap.Error(err))
	} else {
		global.Log.Info("启动端口", zap.String("port", string(rune(global.SystemConfig.System.Port))))
	}
}
