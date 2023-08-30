package handlers

import (
	"encoding/json"
	"net/http"

	"abc.com/db"
	"abc.com/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("USER_NOT_FOUND"))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	//if exists return error

	db.DB.Where(models.User{Email: user.Email}).First(&user)

	if user.ID != 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("USER_ALREADY_EXISTS"))
		return
	}

	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("USER_NOT_FOUND"))
		return
	}

	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
}
