package api

import (
	"encoding/json"
	"estebandev_api/db"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func ResourceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ResourceGetHandler(w, r)
	case http.MethodPost, http.MethodPut:
		ResourceUploadHandler(w, r)
	case http.MethodDelete:
		ResourceDeleteHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func ResourceUploadHandler(w http.ResponseWriter, r *http.Request) {
	var resource db.Resource
	if err := ParseJson(r.Body, &resource); err != nil {
		log.Println("Parsing json", err)
		http.Error(w, "Invalid or malformed JSON", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost {
		resource.ID = 0
	}

	if err := resource.Save(); err != nil {
		log.Println("Error saving resource", err)
		http.Error(w, "Error saving resource", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ResourceDeleteHandler(w http.ResponseWriter, r *http.Request) {
	endpoint := strings.SplitAfter(r.RequestURI, "/resources/")
	if len(endpoint) < 2 {
		http.Error(w, "missing ID", http.StatusBadRequest)
		return
	}

	id, convErr := strconv.Atoi(endpoint[1])
	if convErr != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	resource := db.Resource{ID: id}
	err := resource.Delete()
	if err != nil {
		http.Error(w, "Can't delete this resource", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ResourceGetHandler(w http.ResponseWriter, r *http.Request) {
	endpoint := strings.SplitAfter(r.RequestURI, "/resources/")
	all := len(endpoint) < 2

	if all {
		WriteJson(w, http.StatusOK, db.FindAllResources())
		return
	}

	if id, err := strconv.Atoi(endpoint[1]); err == nil {
		resource, dbErr := db.FindResource(id)
		if dbErr != nil {
			WriteJson(w, http.StatusNotFound, nil)
			return
		}

		data, err := json.Marshal(resource)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}
}
