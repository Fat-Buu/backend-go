package user

import "github.com/google/uuid"

type UserRequest struct {
	Username string `json:"username"`
}

type UserResponse struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}
