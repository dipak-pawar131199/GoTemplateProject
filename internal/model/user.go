package model

type User struct {
	UserId    int    `json:"Id"`
	UserName  string `json:"name"`
	UserEmail string `json:"email"`
	Deleted   int    `json:"deleted"`
}
