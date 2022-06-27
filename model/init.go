package model

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect(dialector gorm.Dialector) {
	var err error
	db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}
	if db == nil {
		logrus.Fatal("Database is nil")
	}
	logrus.Info("Database connected")
}

func CreateTables() {
	if db == nil {
		logrus.Fatal("Database is nil")
	}

	err := db.AutoMigrate(&Student{}, &Exam{}, &Score{})

	if err != nil {
		logrus.Fatal(err)
	}
}
