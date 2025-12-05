package user

import "github.com/google/uuid"

// Business logic

// UserService contains business logic
type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) getAllUser() []UserResponse {
	return ToUserResponseList(GetAll())
}

func (s *UserService) GetUserByID(id uuid.UUID) (UserResponse, bool) {
	user, found := GetByID(id)
	if !found {
		return UserResponse{}, false
	}
	return ToUserResponse(user), true
}
