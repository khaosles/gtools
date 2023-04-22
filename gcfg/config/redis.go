package config

/*
   @File: redis.go
   @Author: khaosles
   @Time: 2023/4/11 21:11
   @Desc:
*/

type Redis struct {
	Addr     string `mapstructure:"addr" default:"" yaml:"addr" json:"addr"`
	Password string `mapstructure:"password" default:"" yaml:"password" json:"password"`
	DB       string `mapstructure:"db" default:"" yaml:"db" json:"db"`
}
