package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/akhil/go-Studentstore/pkg/models"
	"github.com/akhil/go-Studentstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewStudent models.Student

func GetStudent(w http.ResponseWriter, r *http.Request) {
	newStudents := models.GetAllStudents()
	res, _ := json.Marshal(newStudents)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetStudentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	StudentId := vars["StudentId"]
	Id, err := strconv.ParseInt(StudentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	StudentDetails, _ := models.GetStudentById(Id)
	res, _ := json.Marshal(StudentDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	CreateStudent := &models.Student{}
	utils.ParseBody(r, CreateStudent)
	b := CreateStudent.CreateStudent()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	StudentId := vars["StudentId"]
	ID, err := strconv.ParseInt(StudentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	Student := models.DeleteStudent(ID)
	res, _ := json.Marshal(Student)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var updateStudent = &models.Student{}
	utils.ParseBody(r, updateStudent)
	vars := mux.Vars(r)
	StudentId := vars["StudentId"]
	ID, err := strconv.ParseInt(StudentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	StudentDetails, db := models.GetStudentById(ID)
	if updateStudent.Name != "" {
		StudentDetails.Name = updateStudent.Name
	}
	if updateStudent.Class != "" {
		StudentDetails.Class = updateStudent.Class
	}
	if updateStudent.Marks != "" {
		StudentDetails.Marks = updateStudent.Marks
	}
	db.Save(&StudentDetails)
	res, _ := json.Marshal(StudentDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
