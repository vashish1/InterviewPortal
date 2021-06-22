package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/vashish1/InterviewPortal/api"
)

var port = os.Getenv("PORT")
var r *mux.Router

func setuproutes() {
	//creating schedule
	r.HandleFunc("/admin/create", api.AddInterview).Methods("POST")

	r.HandleFunc("/admin/edit", api.EditInterview).Methods("PUT")

	r.HandleFunc("/admin/list", api.GetInterviewList).Methods("GET")

	r.HandleFunc("/admin/list/{email}",api.GetInterviewDetails).Methods("GET")

	r.HandleFunc("/add/user", api.AddUser).Methods("POST")

	r.HandleFunc("/list/users", api.GetUsers).Methods("GET")
}

func main() {
	if port == "" {
		port = "8000"
	}
	r = mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	setuproutes()
	http.Handle("/", handlers.CORS(headers, methods, origins)(r))
	http.ListenAndServe(":"+port, nil)
}
