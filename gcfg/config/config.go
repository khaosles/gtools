package config

/*
   @File: config.go
   @Author: khaosles
   @Time: 2023/4/10 23:27
   @Desc:
*/

type Config struct {
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Server  Server  `mapstructure:"server" json:"server" yaml:"server"`
	Pgsql   Pgsql   `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Sqlite  Sqlite  `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	Logging Logging `mapstructure:"logging" json:"logging" yaml:"logging"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
}
