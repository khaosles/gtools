package mysql

import (
	"github.com/khaosles/gtools/gcfg"
	"github.com/khaosles/gtools/gdb/internal"
	"github.com/khaosles/gtools/glog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
   @File: mysql.go
   @Author: khaosles
   @Time: 2023/4/15 22:32
   @Desc:
*/

var DB *gorm.DB

// GormMysql 初始化Mysql数据库
func init() {
	var err error
	cfg := gcfg.GCfg.Mysql
	mysqlConfig := mysql.Config{
		DSN:                       cfg.Dsn(), // DSN data source name
		DefaultStringSize:         256,       // string 类型字段的默认长度
		SkipInitializeWithVersion: true,      // 根据版本自动配置
	}
	if DB, err = gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(cfg.Prefix, cfg.Singular, cfg.LogMode, cfg.LogZap)); err != nil {
		glog.Error("Database connection failure=> ", cfg.Dsn())
		return
	} else {
		DB.InstanceSet("gorm:table_options", "ENGINE="+cfg.Engine)
		sqlDB, _ := DB.DB()
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
		return
	}
}
