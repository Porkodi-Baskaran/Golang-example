package models

type StudentDetails struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Class   string `json:"class"`
	Address string `json:"address"`
}
