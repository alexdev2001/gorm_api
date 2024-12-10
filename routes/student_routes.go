package routes

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"gorm_api/handlers"
	"gorm_api/services"
	"net/http"
)

func SetupStudentRoutes(db *gorm.DB) (r *mux.Router) {
	router := mux.NewRouter()

	// initialize the service and handler
	studentService := services.NewStudentService(db)
	studentHandler := handlers.NewStudentHandler(studentService)

	// Define routes
	router.HandleFunc("/students", studentHandler.GetAllStudents).Methods(http.MethodGet)
	router.HandleFunc("/students/{id}", studentHandler.GetStudentById).Methods(http.MethodGet)
	router.HandleFunc("/students", studentHandler.AddStudent).Methods(http.MethodPost)
	router.HandleFunc("/students/{id}", studentHandler.UpdateStudentById).Methods(http.MethodPut)
	router.HandleFunc("/students/{id}", studentHandler.DeleteStudentById).Methods(http.MethodDelete)

	return router
}
