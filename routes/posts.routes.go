package routes

import (
	"encoding/json"
	"github.com/armandosalazar/go-gorm-rest-api/db"
	"github.com/armandosalazar/go-gorm-rest-api/models"
	"net/http"
)

func GetPostsHandle(w http.ResponseWriter, r *http.Request) {
	connection := db.GetConnection()
	var posts []models.Post
	connection.Find(&posts)
	w.Header().Set("Content-Type", "applications/json")
	json.NewEncoder(w).Encode(&posts)
}
