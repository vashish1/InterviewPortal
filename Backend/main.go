package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/vashish1/InterviewPortal/Backend/api"
)

var port = os.Getenv("PORT")
var r *mux.Router

func setuproutes() {
	r.HandleFunc("/admin/create", api.AddInterview).Methods("POST")
	r.HandleFunc("/admin/edit", api.EditInterview).Methods("PUT")
	r.HandleFunc("/admin/list", api.GetInterviewList).Methods("GET")
	r.HandleFunc("/add/user", api.AddUser).Methods("POST")
	r.HandleFunc("/list/users", api.GetUsers).Methods("GET")
}

func main() {
	if port == "" {
		port = "8000"
	}
	r = mux.NewRouter()
	setuproutes()
	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}
