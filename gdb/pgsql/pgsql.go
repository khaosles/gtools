package pgsql

import (
	"gorm.io/driver/postgres"

	"github.com/khaosles/gtools/gcfg"
	"github.com/khaosles/gtools/gdb/internal"
	"gorm.io/gorm"
)

/*
   @File: pgsql.go
   @Author: khaosles
   @Time: 2023/4/22 23:39
   @Desc:
*/

var DB *gorm.DB

func init() {
	var err error
	cfg := gcfg.GCfg.Pgsql
	pgsqlConfig := postgres.Config{
		DSN:                  cfg.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	if DB, err = gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config(cfg.Prefix, cfg.Singular, cfg.LogMode, cfg.LogZap)); err != nil {
		return
	} else {
		sqlDB, _ := DB.DB()
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	}
}
