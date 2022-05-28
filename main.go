package main

import (
	"scoreman/conf"
	"scoreman/model"
	"scoreman/web"

	"gorm.io/driver/mysql"
)

func main() {
	conf.Init()

	model.Connect(mysql.Open(conf.GetDatabaseLoginInfo()))
	model.CreateTables()

	web.InitWebFramework()
	web.StartServer()
}
