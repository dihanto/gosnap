package response

import (
	"time"

	"github.com/google/uuid"
)

type UserRegister struct {
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	CreatedAT time.Time `json:"createdAt"`
}

type UserUpdate struct {
	Id        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updatedAt"`
}
