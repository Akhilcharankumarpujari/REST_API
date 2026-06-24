package models

type Book struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	RollNumber int    `json:"roll_number"`
}
