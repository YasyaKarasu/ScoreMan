package web

import (
	"fmt"
	"net/http"
	"scoreman/model"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func getStudentById(c echo.Context) error {
	var id uint
	err := echo.QueryParamsBinder(c).MustUint("id", &id).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	stu, stuerr := model.QueryStudentById(id)
	if stuerr != nil {
		return c.JSON(http.StatusInternalServerError, stuerr)
	}

	return c.JSON(http.StatusOK, stu)
}

func getStudentByStudentId(c echo.Context) error {
	var sid string
	err := echo.QueryParamsBinder(c).MustString("sid", &sid).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	stu, stuerr := model.QueryStudentByStudentId(sid)
	if stuerr != nil {
		return c.JSON(http.StatusInternalServerError, stuerr)
	}

	return c.JSON(http.StatusOK, stu)
}

func getAllStudentsByName(c echo.Context) error {
	var name string
	err := echo.QueryParamsBinder(c).MustString("name", &name).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	stu, stuerr := model.QueryAllStudentsByName(name)
	if stuerr != nil {
		return c.JSON(http.StatusInternalServerError, stuerr)
	}

	return c.JSON(http.StatusOK, stu)
}

func createStudent(c echo.Context) error {
	id, err := model.CreateStudent(&model.Student{StudentID: c.FormValue("sid"), Name: c.FormValue("name")})
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, "create student failed")
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("created success. id: %d", id))
}

func updateStudent(c echo.Context) error {
	id, uinterr := strconv.ParseUint(c.FormValue("id"), 10, 32)
	if uinterr != nil {
		return c.JSON(http.StatusBadRequest, uinterr)
	}

	stu, stuerr := model.QueryStudentById(uint(id))
	if stuerr != nil {
		return c.JSON(http.StatusInternalServerError, stuerr)
	}

	sid := c.FormValue("sid")
	name := c.FormValue("name")
	if sid != "" {
		stu.StudentID = sid
	}
	if name != "" {
		stu.Name = name
	}

	upderr := model.UpdateStudent(stu, uint(id))
	if upderr != nil {
		logrus.Error(upderr)
		return c.JSON(http.StatusInternalServerError, upderr)
	}
	return c.JSON(http.StatusOK, stu)
}

func deleteStudent(c echo.Context) error {
	var id uint
	err := echo.QueryParamsBinder(c).MustUint("id", &id).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	delerr := model.DeleteStudentById(id)
	if delerr != nil {
		logrus.Error(delerr)
		return c.JSON(http.StatusInternalServerError, delerr)
	}
	return c.JSON(http.StatusOK, "deleted success")
}
