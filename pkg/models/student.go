// package models
package models

import (
	"github.com/akhil/go-Studentstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Student struct {
	gorm.Model
	Name  string `gorm:"" json:"name"`
	Class string `json:"class"`
	Marks string `json:"marks"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate((&Student{}))
}

func (b *Student) CreateStudent() *Student {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllStudents() []Student {
	var Students []Student
	db.Find(&Students)
	return Students
}

func GetStudentById(Id int64) (*Student, *gorm.DB) {
	var getStudent Student
	db := db.Where("ID=?", Id).Find(&getStudent)
	return &getStudent, db
}

func DeleteStudent(ID int64) Student {
	var Student Student
	db.Where("ID= ?", ID).Delete(Student)
	return Student
}
