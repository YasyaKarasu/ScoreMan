package model

import "errors"

type Exam struct {
	ExamID   uint   `gorm:"not null;autoIncrement;primaryKey"`
	ExamName string `gorm:"not null"`
}

type Score struct {
	ScoreId uint `gorm:"not null;autoIncrement;primaryKey"`
	ExamID  uint `gorm:"not null"`
	SId     uint `gorm:"not null"` // id of student
	Score   int  `gorm:"not null"`
}

func CreateExam(exam *Exam) (uint, error) {
	result := db.Create(exam)
	return exam.ExamID, result.Error
}

func UpdateExam(exam *Exam) error {
	result := db.Model(&Exam{ExamID: exam.ExamID}).Updates(exam)
	return result.Error
}

func QueryExamById(eid uint) (*Exam, error) {
	var exam Exam
	result := db.Model(&Exam{}).First(&exam, eid)
	return &exam, result.Error
}

func QueryExamByName(ename string) (*Exam, error) {
	var exam Exam
	result := db.Model(&Exam{ExamName: ename}).First(&exam)
	return &exam, result.Error
}

func DeleteExamById(eid uint) error {
	var exam Exam

	result := db.Model(&Exam{}).First(&exam, eid)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}

	result = db.Delete(&exam)
	return result.Error
}

func CreateScore(score *Score) (uint, error) {
	result := db.Create(score)
	return score.ScoreId, result.Error
}

func UpdateScore(score *Score) error {
	result := db.Model(&Score{}).Where("score_id = ?", score.ScoreId).Updates(score)
	return result.Error
}

func QueryScoreById(scrid uint) (*Score, error) {
	var score Score
	result := db.Model(&Score{}).First(&score, scrid)
	return &score, result.Error
}

func QueryAllScoreByExamId(eid uint) (*[]Score, error) {
	var scr []Score
	if findScoreError := db.Model(&Score{}).Where("exam_id = ?", eid).First(&Score{}).Error; findScoreError != nil {
		return nil, findScoreError
	}

	result := db.Model(&Score{}).Where("exam_id = ?", eid).Find(&scr)
	return &scr, result.Error
}

func QueryAllScoreByExamName(ename string) (*[]Score, error) {
	exam, err := QueryExamByName(ename)
	if err != nil {
		return nil, err
	}

	return QueryAllScoreByExamId(exam.ExamID)
}

func QueryAllScoreByStudentId(sid string) (*[]Score, error) {
	stu, stuerr := QueryStudentByStudentId(sid)
	if stuerr != nil {
		return nil, stuerr
	}

	var scr []Score
	if findScoreError := db.Model(&Score{}).Where("s_id = ?", stu.ID).First(&Score{}).Error; findScoreError != nil {
		return nil, findScoreError
	}

	result := db.Model(&Score{}).Where("s_id = ?", stu.ID).Find(&scr)
	return &scr, result.Error
}

func QueryAllScoreByStudentName(name string) (*[]Score, error) {
	stu, stuerr := QueryAllStudentsByName(name)
	if stuerr != nil {
		return nil, stuerr
	}

	var scr []Score
	for _, v := range *stu {
		currentScr, err := QueryAllScoreByStudentId(v.StudentID)
		if err != nil {
			return nil, err
		}
		scr = append(scr, *currentScr...)
	}
	return &scr, nil
}

func DeleteScoreById(scrid uint) error {
	var score Score

	result := db.Model(&Score{}).Where("score_id = ?", scrid).First(&score)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}

	result = db.Delete(&score)
	return result.Error
}
