package main

import (
	"github.com/armandosalazar/go-gorm-rest-api/db"
	"github.com/armandosalazar/go-gorm-rest-api/models"
	"github.com/armandosalazar/go-gorm-rest-api/routes"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	// Database Connection
	conn := db.GetConnection()
	conn.AutoMigrate(models.User{}, models.Post{})

	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/posts", routes.PostUserHandler).Methods("GET")

	http.ListenAndServe(":3000", router)
}
