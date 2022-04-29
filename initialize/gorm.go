package initialize

import (
	"GinProject/global"
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {
	switch global.SystemConfig.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}
