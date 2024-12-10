package main

import (
	"fmt"
	"gorm_api/db"
	"gorm_api/routes"
	"log"
	"net/http"
)

func main() {
	dbConn, err := db.DatabaseConnection()
	if err != nil {
		log.Fatal("Failed to connect to database: %v", err)
	}

	// set up routes
	router := routes.SetupStudentRoutes(dbConn)

	fmt.Println("Starting server on port 8080...")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Failed to start server: %v", err)
	}
}
