package global

import (
	"GinProject/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	SystemConfig config.Server
	Vp           *viper.Viper
	Log          *zap.Logger
	Db           *gorm.DB
)
