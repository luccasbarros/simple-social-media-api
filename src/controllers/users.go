package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Create user create an user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User

	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("signup"); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepositories(db)

	user.ID, erro = repository.Create(user)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, user)

}

// GetUsers get all users from database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepositories(db)

	users, erro := repository.List(nameOrNick)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, users)

}

// GetUser get an user from database
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, erro := strconv.ParseUint(params["id"], 10, 64)

	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repository := repositories.NewUsersRepositories(db)
	user, erro := repository.GetById(userID)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if user.ID == 0 {
		responses.Erro(w, http.StatusNotFound, errors.New("User not found"))
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// UpdateUser updates an user in the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["id"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	tokenUserId, erro := auth.GetUserId(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userId != tokenUserId {
		responses.Erro(w, http.StatusForbidden, errors.New("This is not your user"))
		return
	}

	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro := json.Unmarshal(requestBody, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("update"); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepositories(db)
	if erro := repository.Update(userId, user); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteUseer delete user info from database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["id"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	tokenUserId, erro := auth.GetUserId(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
	}

	if userId != tokenUserId {
		responses.Erro(w, http.StatusForbidden, errors.New("It's not possible delete an user that's not yours"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepositories(db)
	if erro = repository.Delete(userId); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

//FollowUser
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, erro := auth.GetUserId(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["id"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if followerId == userId {
		responses.Erro(w, http.StatusForbidden, errors.New("It's not possible follow yourself"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepositories(db)
	if erro = repository.Follow(userId, followerId); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, erro := auth.GetUserId(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["id"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if followerId == userId {
		responses.Erro(w, http.StatusForbidden, errors.New("It's not possible unfollow yourself"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepositories(db)
	if erro = repository.Unfollow(userId, followerId); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}
