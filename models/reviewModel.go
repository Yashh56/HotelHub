package models

type ReviewModel struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}
