package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, erro := auth.GetUserId(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var post models.Post
	if erro = json.Unmarshal(requestBody, &post); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	post.AuthorID = userID

	if erro = post.Prepare(); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	post.ID, erro = repository.Create(post)

	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, post)
}
func GetPersonalPosts(w http.ResponseWriter, r *http.Request) {
	userID, erro := auth.GetUserId(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	posts, erro := repository.GetPersonalPosts(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}
func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, erro := strconv.ParseUint(params["post_id"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	post, erro := repository.GetById(postID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, post)
}
func UpdatePost(w http.ResponseWriter, r *http.Request) {}
func DeletePost(w http.ResponseWriter, r *http.Request) {}
