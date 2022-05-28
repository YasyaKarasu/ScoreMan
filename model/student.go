package model

import "errors"

type Student struct {
	ID        uint   `gorm:"not null;autoIncrement;primaryKey"`
	StudentID string `gorm:"not null"`
	Name      string `gorm:"not null"`
}

func CreateStudent(student *Student) (uint, error) {
	result := db.Create(student)
	return student.ID, result.Error
}

func UpdateStudent(student *Student, id uint) error {
	result := db.Model(&Student{ID: id}).Updates(student)
	return result.Error
}

func QueryStudentById(id uint) (*Student, error) {
	var stu Student
	result := db.Model(&Student{}).First(&stu, id)
	return &stu, result.Error
}

func QueryStudentByStudentId(sid string) (*Student, error) {
	var stu Student
	result := db.Where(&Student{StudentID: sid}).First(&stu)
	return &stu, result.Error
}

func QueryAllStudentsByName(name string) (*[]Student, error) {
	var stu []Student

	if findNameError := db.Where("name = ?", name).First(&Student{}).Error; findNameError != nil {
		return nil, findNameError
	}

	result := db.Model(&Student{}).Where(&Student{Name: name}).Find(&stu)
	return &stu, result.Error
}

func DeleteStudentById(id uint) error {
	var stu Student

	result := db.Model(&Student{}).First(&stu, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}

	result = db.Delete(&stu)
	return result.Error
}
