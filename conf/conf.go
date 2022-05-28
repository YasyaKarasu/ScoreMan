package conf

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("conf")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic(err)
	}
	logrus.Info("Configuration file loaded")

	var confItem = []string{"user", "password", "host", "port", "db_name"}
	for i := range confItem {
		if !viper.IsSet(confItem[i]) {
			logrus.WithField(confItem[i], nil).Fatal("The following item in conf.yaml isn't set properly : ")
		}
	}
}

func GetDatabaseLoginInfo() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.Get("user"),
		viper.Get("password"),
		viper.Get("host"),
		viper.Get("port"),
		viper.Get("db_name"))
}
