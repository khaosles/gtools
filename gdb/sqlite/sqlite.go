package sqlite

import (
	"github.com/khaosles/gtools/gcfg"
	"github.com/khaosles/gtools/gdb/internal"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
   @File: sqlite.go
   @Author: khaosles
   @Time: 2023/4/23 01:02
   @Desc:
*/

var DB *gorm.DB

func init() {
	var err error
	cfg := gcfg.GCfg.Sqlite
	if DB, err = gorm.Open(sqlite.Open(cfg.Dsn()), internal.Gorm.Config(
		cfg.Prefix, cfg.Singular, cfg.LogMode, cfg.LogZap),
	); err != nil {
		return
	} else {
		sqlDB, _ := DB.DB()
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	}
}
