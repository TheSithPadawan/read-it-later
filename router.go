package main

import "github.com/gorilla/mux"

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/articles", SaveArticleHandler).Methods("POST")
	r.HandleFunc("/articles", GetArticlesHandler).Methods("GET")
	return r
}
