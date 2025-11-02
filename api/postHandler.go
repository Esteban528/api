package api

import (
	"encoding/json"
	"estebandev_api/db"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		PostGetHandler(w, r)
	case http.MethodPost, http.MethodPut:
		PostUploadHandler(w, r)
	case http.MethodDelete:
		PostDeleteHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func PostUploadHandler(w http.ResponseWriter, r *http.Request) {
	var post db.Post
	println(r.Body)
	if err := ParseJson(r.Body, &post); err != nil {
		log.Println("Parsing json", err)
		http.Error(w, "Invalid or malformed JSON", http.StatusBadRequest)
		return
	}

	println("struct")
	fmt.Println(post)

	if r.Method == http.MethodPost {
		post.ID = 0
	}

	if err := post.Save(); err != nil {
		log.Println("Error saving post", err)
		http.Error(w, "Error saving post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func PostDeleteHandler(w http.ResponseWriter, r *http.Request) {
	endpoint := strings.SplitAfter(r.RequestURI, "/post/")
	if len(endpoint) < 2 {
		http.Error(w, "missing ID", http.StatusBadRequest)
		return
	}

	id, convErr := strconv.Atoi(endpoint[1])
	if convErr != nil {
		http.Error(w, "invalid ID ", http.StatusBadRequest)
		return
	}

	post := db.Post{ID: id}
	err := post.Delete()
	if err != nil {
		http.Error(w, "Can't delete this post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func PostGetHandler(w http.ResponseWriter, r *http.Request) {
	endpoint := strings.SplitAfter(r.RequestURI, "/post/")
	all := len(endpoint) < 2

	if all {
		WriteJson(w, http.StatusOK, db.FindAllPost())
		return
	}

	if id, err := strconv.Atoi(endpoint[1]); err == nil {
		post, dbErr := db.FindPost(id)
		if dbErr != nil {
			WriteJson(w, http.StatusNotFound, nil)
			return
		}

		data, err := json.Marshal(post)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}
}
