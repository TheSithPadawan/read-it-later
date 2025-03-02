package main

import "time"

type Article struct {
	ID          int       `json:"id"`
	URL         string    `json:"url"`
	Title       string    `json:"title"`
	Description string    `json:description"`
	Tags        []string  `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
}
