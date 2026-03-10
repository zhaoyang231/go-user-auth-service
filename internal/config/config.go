package config

import (
	"fmt"
	"go-user-auth-service/internal/model"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AppConfig struct {
	Name string
	Env  string
	Port int
}

var App AppConfig
var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: &s", err))
	}
	App = AppConfig{
		Name: viper.GetString("app.name"),
		Env:  viper.GetString("app.env"),
		Port: viper.GetInt("app.port"),
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.dbname"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("connect mysql failed: &s", err))
	}
	DB.AutoMigrate(&model.User{})
}
