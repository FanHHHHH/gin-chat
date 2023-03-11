package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// println(viper.Get("app"))
	fmt.Println("app", viper.Get("app"))
	fmt.Println("mysql", viper.Get("mysql"))
}

func InitMysql() {
	logger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Info,
		Colorful:      true,
	})

	var err error

	DB, err = gorm.Open(mysql.Open(
		viper.GetString("mysql.username")+":"+
			viper.GetString("mysql.password")+"@tcp("+
			viper.GetString("mysql.host")+":"+
			viper.GetString("mysql.port")+")/"+
			viper.GetString("mysql.dbname")+"?"+
			viper.GetString("mysql.charset")+"&"+
			viper.GetString("mysql.parseTime")+"&"+
			viper.GetString("mysql.loc"),
	), &gorm.Config{Logger: logger})
	if err != nil {
		panic("failed to connect database")
	}
}
