package Handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"site-api/models"
	. "site-api/repository"
	"site-api/utils"
)

var userRepository User

var GetAllUsers = func(w http.ResponseWriter, r *http.Request) {
	users, err := userRepository.GetAll()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Info(users)
	utils.RespondWithJson(w, http.StatusOK, users)
}

var GetUserById = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented!")
}

var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	log.Info(user)
	insertResult, err := userRepository.Insert(user)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	user.ID = insertResult.(primitive.ObjectID)
	utils.RespondWithJson(w, http.StatusCreated, user)
}

var UpdateUser = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented")
}

var DeleteUser = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
	if err := userRepository.Delete(id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
