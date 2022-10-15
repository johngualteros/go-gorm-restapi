package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/johngualteros/go-gorm-restapi/db"
	"github.com/johngualteros/go-gorm-restapi/models"
	"github.com/johngualteros/go-gorm-restapi/routes"
)

func main() {

	db.DBConnection()
	
	db.DB.AutoMigrate(&models.User{}, &models.Task{})

	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler).Methods("GET")
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	router.HandleFunc("/tasks", routes.PostTaskHandler).Methods("POST")
	router.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
