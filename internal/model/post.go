package model

type Post struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	DateOfCreation string `json:"date"`
	Author         string `json:"author"`
}
