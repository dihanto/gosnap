package response

import (
	"time"

	"github.com/google/uuid"
)

type UserRegister struct {
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAT time.Time `json:"createdAt"`
}

type UserUpdate struct {
	Id             uuid.UUID `json:"id"`
	Email          string    `json:"email"`
	Username       string    `json:"username"`
	ProfilePicture string    `json:"profilePicture"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type FindUser struct {
	Username       string `json:"username"`
	Name           string `json:"name"`
	ProfilePicture string `json:"profilePicture"`
}

type FindAllUser struct {
	Username       string `json:"username"`
	ProfilePicture string `json:"profilePicture"`
}
