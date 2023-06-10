package web

import "time"

type UserRegister struct {
	Email     string    `validate:"required,email,email_uniq" json:"email"`
	Username  string    `validate:"required,username_uniq" json:"username"`
	Password  string    `validate:"required,min=6" json:"password"`
	Age       int       `validate:"required,gt=8" json:"age"`
	CreatedAT time.Time `json:"createdAt"`
}

type UserUpdate struct {
	Id        int       `json:"id"`
	Email     string    `validate:"required,email" json:"email"`
	Username  string    `validate:"required" json:"username"`
	Age       int       `validate:"required,gt=8" json:"age"`
	UpdatedAt time.Time `json:"updatedAt"`
}
