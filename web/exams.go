package web

import (
	"fmt"
	"net/http"
	"scoreman/model"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func getExamById(c echo.Context) error {
	var eid uint
	err := echo.QueryParamsBinder(c).MustUint("eid", &eid).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	exam, examerr := model.QueryExamById(eid)
	if examerr != nil {
		return c.JSON(http.StatusInternalServerError, examerr)
	}

	return c.JSON(http.StatusOK, exam)
}

func getExamByName(c echo.Context) error {
	var ename string
	err := echo.QueryParamsBinder(c).MustString("ename", &ename).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	exam, examerr := model.QueryExamByName(ename)
	if examerr != nil {
		return c.JSON(http.StatusInternalServerError, examerr)
	}

	return c.JSON(http.StatusOK, exam)
}

func createExam(c echo.Context) error {
	eid, err := model.CreateExam(&model.Exam{ExamName: c.FormValue("ename")})
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("created success. eid: %d", eid))
}

func updateExam(c echo.Context) error {
	eid, uinterr := strconv.ParseUint(c.FormValue("eid"), 10, 32)
	if uinterr != nil {
		return c.JSON(http.StatusBadRequest, uinterr)
	}

	exam, examerr := model.QueryExamById(uint(eid))
	if examerr != nil {
		return c.JSON(http.StatusInternalServerError, examerr)
	}

	ename := c.FormValue("ename")

	exam.ExamName = ename

	upderr := model.UpdateExam(exam)
	if upderr != nil {
		logrus.Error(upderr)
		return c.JSON(http.StatusInternalServerError, upderr)
	}
	return c.JSON(http.StatusOK, exam)
}

func deleteExam(c echo.Context) error {
	var eid uint
	err := echo.QueryParamsBinder(c).MustUint("eid", &eid).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	delerr := model.DeleteExamById(eid)
	if delerr != nil {
		logrus.Error(delerr)
		return c.JSON(http.StatusInternalServerError, delerr)
	}
	return c.JSON(http.StatusOK, "deleted success")
}

func getScoreById(c echo.Context) error {
	var scrid uint
	err := echo.QueryParamsBinder(c).MustUint("scrid", &scrid).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	score, scoreerr := model.QueryScoreById(scrid)
	if scoreerr != nil {
		return c.JSON(http.StatusInternalServerError, scoreerr)
	}
	return c.JSON(http.StatusOK, score)
}

func getScoreByExamId(c echo.Context) error {
	var eid uint
	err := echo.QueryParamsBinder(c).MustUint("eid", &eid).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	scores, scoreserr := model.QueryAllScoreByExamId(eid)
	if scoreserr != nil {
		return c.JSON(http.StatusInternalServerError, scoreserr)
	}
	return c.JSON(http.StatusOK, scores)
}

func getScoreByExamName(c echo.Context) error {
	var ename string
	err := echo.QueryParamsBinder(c).MustString("ename", &ename).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	scores, scoreserr := model.QueryAllScoreByExamName(ename)
	if scoreserr != nil {
		return c.JSON(http.StatusInternalServerError, scoreserr)
	}
	return c.JSON(http.StatusOK, scores)
}

func getScoreByStudentId(c echo.Context) error {
	var sid string
	err := echo.QueryParamsBinder(c).MustString("sid", &sid).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	scores, scoreserr := model.QueryAllScoreByStudentId(sid)
	if scoreserr != nil {
		return c.JSON(http.StatusInternalServerError, scoreserr)
	}
	return c.JSON(http.StatusOK, scores)
}

func getScoreByStudentName(c echo.Context) error {
	var name string
	err := echo.QueryParamsBinder(c).MustString("name", &name).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, name)
	}

	scores, scoreserr := model.QueryAllScoreByStudentName(name)
	if scoreserr != nil {
		return c.JSON(http.StatusInternalServerError, scoreserr)
	}
	return c.JSON(http.StatusOK, scores)
}

func createScore(c echo.Context) error {
	eid, uinterr := strconv.ParseUint(c.FormValue("eid"), 10, 32)
	if uinterr != nil {
		return c.JSON(http.StatusBadRequest, uinterr)
	}

	scr, interr := strconv.Atoi(c.FormValue("score"))
	if interr != nil {
		return c.JSON(http.StatusBadRequest, interr)
	}

	sid, uinterr := strconv.ParseUint(c.FormValue("sid"), 10, 32)
	if uinterr != nil {
		return c.JSON(http.StatusBadRequest, uinterr)
	}

	scrid, err := model.CreateScore(&model.Score{ExamID: uint(eid), SId: uint(sid), Score: scr})
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("create success. scoreid : %d", scrid))
}

func updateScore(c echo.Context) error {
	scrid, uinterr := strconv.ParseUint(c.FormValue("scrid"), 10, 32)
	if uinterr != nil {
		return c.JSON(http.StatusBadRequest, uinterr)
	}

	score, scoreerr := model.QueryScoreById(uint(scrid))
	if scoreerr != nil {
		return c.JSON(http.StatusInternalServerError, scoreerr)
	}

	scr, interr := strconv.Atoi(c.FormValue("score"))
	if interr != nil {
		return c.JSON(http.StatusBadRequest, interr)
	}

	score.Score = scr
	upderr := model.UpdateScore(score)
	if upderr != nil {
		return c.JSON(http.StatusInternalServerError, upderr)
	}
	return c.JSON(http.StatusOK, score)
}

func deleteScore(c echo.Context) error {
	var scrid uint

	err := echo.QueryParamsBinder(c).MustUint("scrid", &scrid).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	delerr := model.DeleteScoreById(scrid)
	if delerr != nil {
		logrus.Error(delerr)
		return c.JSON(http.StatusInternalServerError, delerr)
	}
	return c.JSON(http.StatusOK, "delete success")
}
