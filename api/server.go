package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var userE string
var passE string

func ParseJson(body io.ReadCloser, v any) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(&v)
}

func WriteJson(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if v != nil {
		data, err := json.Marshal(v)
		if err != nil {
			http.Error(w, "Error parsing JSON", http.StatusInternalServerError)
		}

		w.Write(data)
	}
}

func DebugRequest(r *http.Request) {
	content, _ := io.ReadAll(r.Body)

	r.Body = io.NopCloser(bytes.NewBuffer(content))

	log.Println("Request log")
	log.Println("\tMethod: " + r.Method)
	log.Println("\tAuth Header: " + r.Header.Get("Authorization"))
	log.Println("\tContent: " + string(content))
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		DebugRequest(r)
		if strings.HasPrefix(r.RequestURI, "/post") && r.Method == http.MethodGet {
			next.ServeHTTP(w, r)
			return
		}

		if strings.HasPrefix(r.RequestURI, "/projects") && r.Method == http.MethodGet {
			next.ServeHTTP(w, r)
			return
		}

		token := r.Header.Get("Authorization")

		if token == "" || !strings.HasPrefix(token, "Basic ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		token = strings.TrimPrefix(token, "Basic ")

		raw, err := base64.StdEncoding.DecodeString(token)

		if err != nil {
			log.Println("Authorization failed ", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		args := strings.Split(string(raw), ":")

		if len(args) < 2 {
			log.Println("Authorization failed malformed token", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if args[0] != userE {
			log.Println("Authorization failed malformed token user expected="+userE+" received="+args[0], err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if args[1] != passE {
			log.Println("Authorization failed malformed token user expected="+userE+" received="+args[0], err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Load() {
	userE = os.Getenv("API_USER")
	passE = os.Getenv("API_PASSWORD")
	log.Println("API ENVs loaded")

	mux := http.NewServeMux()
	mux.Handle("/post", AuthMiddleware(http.HandlerFunc(PostHandler)))
	mux.Handle("/post/", AuthMiddleware(http.HandlerFunc(PostHandler)))

	mux.Handle("/projects", AuthMiddleware(http.HandlerFunc(ProjectHandler)))
	mux.Handle("/projects/", AuthMiddleware(http.HandlerFunc(ProjectHandler)))

	log.Println("Server listenning at 8080")
	http.ListenAndServe(":8080", mux)
}
