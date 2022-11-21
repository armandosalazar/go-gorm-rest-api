package routes

import (
	"encoding/json"
	"github.com/armandosalazar/go-gorm-rest-api/db"
	"github.com/armandosalazar/go-gorm-rest-api/models"
	"github.com/gorilla/mux"
	"net/http"
)

func GetUsersHandler(writer http.ResponseWriter, _ *http.Request) {
	connection := db.GetConnection()
	var users []models.User
	connection.Find(&users)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	connection := db.GetConnection()
	var user models.User
	params := mux.Vars(r)
	connection.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&user)
}

func PostUserHandler(writer http.ResponseWriter, request *http.Request) {
	connection := db.GetConnection()
	var user models.User
	json.NewDecoder(request.Body).Decode(&user)
	createUser := connection.Create(&user)
	if createUser.Error != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(createUser.Error.Error()))
	}
	json.NewEncoder(writer).Encode(user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	connection := db.GetConnection()
	var user models.User
	params := mux.Vars(r)
	connection.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	}
	connection.Delete(&user)
	w.WriteHeader(http.StatusNoContent)
}
