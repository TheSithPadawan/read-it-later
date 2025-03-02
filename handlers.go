package main

import (
	"encoding/json"
	"net/http"
)

// Save articles to collection
func SaveArticleHandler(w http.ResponseWriter, r *http.Request) {
	var article Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	Articles = append(Articles, article)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Article saved to collection!"})
}

// Retrieve all articles
func GetArticlesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Articles)
}
