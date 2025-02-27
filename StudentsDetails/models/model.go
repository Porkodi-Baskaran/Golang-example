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

type StudMarks struct {
	ID            int    `json:"markid"`
	StudID        int    `json:"studid"`
	Class         int    `json:"class"`
	Maths         int    `json:"maths"`
	Science       int    `json:"science"`
	English       int    `json:"english"`
	Tamil         int    `json:"tamil"`
	Socialscience int    `json:"socialSci"`
	PassFail      string `json:"passfail"`
}
