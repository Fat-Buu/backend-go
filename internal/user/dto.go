package user

import "github.com/google/uuid"

type UserRequest struct {
	Username string `json:"username"`
}

type UserResponse struct {
	Id           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	ProfielImage string    `json:"profileImage"`
}
