package web

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func addRoutes(ech *echo.Echo) {
	stu := e.Group("/stu")
	stu.GET("/id", getStudentById)
	stu.GET("/sid", getStudentByStudentId)
	stu.GET("/name", getAllStudentsByName)
	stu.POST("/create", createStudent)
	stu.POST("/update", updateStudent)
	stu.DELETE("/delete", deleteStudent)
}

var e *echo.Echo

func InitWebFramework() {
	e = echo.New()
	e.HideBanner = true
	addRoutes(e)
	logrus.Info("Echo Framework initialized")
}

func StartServer() {
	e.Logger.Fatal(e.Start(":1323"))
}
