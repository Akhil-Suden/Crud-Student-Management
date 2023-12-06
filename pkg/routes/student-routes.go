// package routes
package routes

import (
	"github.com/akhil/go-Studentstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterStudentStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/Student/", controllers.CreateStudent).Methods("POST")
	router.HandleFunc("/Student/", controllers.GetStudent).Methods("GET")
	router.HandleFunc("/Student/{StudentId}", controllers.GetStudentById).Methods("GET")
	router.HandleFunc("/Student/{StudentId}", controllers.UpdateStudent).Methods("PUT")
	router.HandleFunc("/Student/{StudentId}", controllers.DeleteStudent).Methods("DELETE")
}
