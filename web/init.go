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

	exam := e.Group("/exam")
	exam.GET("/eid", getExamById)
	exam.GET("/ename", getExamByName)
	exam.POST("/create", createExam)
	exam.POST("/update", updateExam)
	exam.DELETE("/delete", deleteExam)

	score := e.Group("/score")
	score.GET("/scrid", getScoreById)
	score.GET("/eid", getScoreByExamId)
	score.GET("/ename", getScoreByExamName)
	score.GET("/sid", getScoreByStudentId)
	score.GET("/name", getScoreByStudentName)
	score.POST("/create", createScore)
	score.POST("/update", updateScore)
	score.DELETE("/delete", deleteScore)
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
