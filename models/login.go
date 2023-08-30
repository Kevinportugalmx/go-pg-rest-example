package models

type Login struct {
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}
