package request

import "github.com/google/uuid"

type UserRegister struct {
	Email    string `validate:"required,email,email_uniq" json:"email"`
	Username string `validate:"required,username_uniq" json:"username"`
	Name     string `validate:"required" json:"name"`
	Password string `validate:"required,min=6" json:"password"`
}
type UserLogin struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required,min=6" json:"password"`
}

type UserUpdate struct {
	Id             uuid.UUID `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	ProfilePicture string    `json:"profilePicture"`
}
