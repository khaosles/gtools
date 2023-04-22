package gcfg

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/khaosles/gtools/gcfg/config"
	"github.com/spf13/viper"
)

/*
   @File: Viper.go
   @Author: khaosles
   @Time: 2023/2/19 11:04
   @Desc:
*/

var GCfg config.Config // 配置类
var Viper *viper.Viper // 配置源

func init() {
	var cfg string
	// 命令行获取配置文件
	flag.StringVar(&cfg, "c", "", "choose config file.")
	flag.Parse()
	// 命令行无输入
	if cfg == "" {
		cfg = "config.yml"
	}
	rootPath, _ := os.Getwd()
	cfg = filepath.Join(rootPath, "resource", cfg)

	// 配置文件不存在
	if !exist(cfg) {
		log.Fatal("Configure not exists.")
	}

	// 创建viper
	Viper = viper.New()
	Viper.SetConfigFile(cfg)
	Viper.SetConfigType("yml")

	// 读取配置文件
	err := Viper.ReadInConfig()
	if err != nil {
		log.Fatal("Configure reading error.")
	}
	Viper.WatchConfig()
	if err = Viper.Unmarshal(&GCfg); err != nil {
		log.Fatal("Configure parse error.")
	}
	// 识别 default 标签
	setDefaults(reflect.ValueOf(&GCfg).Elem())
	// 解析配置文件
	Viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err = Viper.Unmarshal(&GCfg); err != nil {
			log.Fatal("Configure parse error.")
		}
		// 识别 default 标签
		setDefaults(reflect.ValueOf(&GCfg).Elem())
	})

}

// setDefaults 设置结构体默认值
func setDefaults(v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		defaultTag := v.Type().Field(i).Tag.Get("default")
		if defaultTag != "" && field.Interface() == reflect.Zero(field.Type()).Interface() {
			var defaultValue any
			switch field.Type().Kind() {
			case reflect.Int:
				defaultValue, _ = strconv.Atoi(defaultTag)
			case reflect.Int32:
				defaultValue, _ = strconv.ParseInt(defaultTag, 8, 32)
			case reflect.Int64:
				defaultValue, _ = time.ParseDuration(defaultTag)
			case reflect.Float32:
				defaultValue, _ = strconv.ParseFloat(defaultTag, 32)
			case reflect.Float64:
				defaultValue, _ = strconv.ParseFloat(defaultTag, 64)
			case reflect.Bool:
				defaultValue, _ = strconv.ParseBool(defaultTag)
			default:
				defaultValue = defaultTag
			}
			field.Set(reflect.ValueOf(defaultValue).Convert(field.Type()))
		}
		if field.Kind() == reflect.Struct {
			setDefaults(field)
		}
	}
}

func exist(path string) bool {
	// path stat
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}
