package web

import (
	"fmt"
	"net/http"
	"scoreman/model"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func getStudentById(c echo.Context) error {
	var id uint
	err := echo.QueryParamsBinder(c).MustUint("id", &id).BindError()
	if err != nil {
		return err
	}

	stu, stuerr := model.QueryStudentById(id)
	if stuerr != nil {
		return stuerr
	}

	return c.JSON(http.StatusOK, stu)
}

func getStudentByStudentId(c echo.Context) error {
	var sid string
	err := echo.QueryParamsBinder(c).MustString("sid", &sid).BindError()
	if err != nil {
		return err
	}

	stu, stuerr := model.QueryStudentByStudentId(sid)
	if stuerr != nil {
		return stuerr
	}

	return c.JSON(http.StatusOK, stu)
}

func getAllStudentsByName(c echo.Context) error {
	var name string
	err := echo.QueryParamsBinder(c).MustString("name", &name).BindError()
	if err != nil {
		return err
	}

	stu, stuerr := model.QueryAllStudentsByName(name)
	if stuerr != nil {
		return stuerr
	}

	return c.JSON(http.StatusOK, stu)
}

func createStudent(c echo.Context) error {
	id, err := model.CreateStudent(&model.Student{StudentID: c.QueryParam("sid"), Name: c.QueryParam("name")})
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, "create student failed")
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("created success. id: %d", id))
}

func updateStudent(c echo.Context) error {
	var id uint
	err := echo.QueryParamsBinder(c).MustUint("id", &id).BindError()
	if err != nil {
		return err
	}

	stu, stuerr := model.QueryStudentById(id)
	if stuerr != nil {
		return stuerr
	}

	sid := c.QueryParam("sid")
	name := c.QueryParam("name")
	if sid != "" {
		stu.StudentID = sid
	}
	if name != "" {
		stu.Name = name
	}

	upderr := model.UpdateStudent(stu, id)
	if upderr != nil {
		return upderr
	}
	return c.JSON(http.StatusOK, stu)
}

func deleteStudent(c echo.Context) error {
	var id uint
	err := echo.QueryParamsBinder(c).MustUint("id", &id).BindError()
	if err != nil {
		return err
	}

	delerr := model.DeleteStudentById(id)
	if delerr != nil {
		return delerr
	}
	return c.JSON(http.StatusOK, "deleted success")
}
