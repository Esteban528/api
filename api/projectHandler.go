package api

import (
	"encoding/json"
	"estebandev_api/db"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func ProjectHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ProjectGetHandler(w, r)
	case http.MethodPost, http.MethodPut:
		ProjectUploadHandler(w, r)
	case http.MethodDelete:
		ProjectDeleteHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func ProjectUploadHandler(w http.ResponseWriter, r *http.Request) {
	var project db.Project
	if err := ParseJson(r.Body, &project); err != nil {
		log.Println("Parsing json", err)
		http.Error(w, "Invalid or malformed JSON", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost {
		project.ID = 0
	}

	if err := project.Save(); err != nil {
		log.Println("Error saving project", err)
		http.Error(w, "Error saving project", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ProjectDeleteHandler(w http.ResponseWriter, r *http.Request) {
	endpoint := strings.SplitAfter(r.RequestURI, "/projects/")
	if len(endpoint) < 2 {
		http.Error(w, "missing ID", http.StatusBadRequest)
		return
	}

	id, convErr := strconv.Atoi(endpoint[1])
	if convErr != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	project := db.Project{ID: id}
	err := project.Delete()
	if err != nil {
		http.Error(w, "Can't delete this project", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ProjectGetHandler(w http.ResponseWriter, r *http.Request) {
	endpoint := strings.SplitAfter(r.RequestURI, "/projects/")
	all := len(endpoint) < 2

	if all {
		WriteJson(w, http.StatusOK, db.FindAllProject())
		return
	}

	if id, err := strconv.Atoi(endpoint[1]); err == nil {
		project, dbErr := db.FindProject(id)
		if dbErr != nil {
			WriteJson(w, http.StatusNotFound, nil)
			return
		}

		data, err := json.Marshal(project)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}
}
