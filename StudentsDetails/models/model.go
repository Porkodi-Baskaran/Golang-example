package models

type StudentDetails struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Class   string `json:"class"`
	Address string `json:"address"`
}

type LoginUser struct {
	UID      uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var Users = []LoginUser{}
