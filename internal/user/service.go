package user

import (
	"github.com/google/uuid"
)

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

func (s *UserService) CreateUser(u UserRequest) (UserResponse, bool) {
	var newUser = User{Id: uuid.New(), Username: u.Username}
	user, err := Add(newUser)
	if !err {
		return UserResponse{}, false
	}
	return ToUserResponse(user), true
}

func (s *UserService) UpdateUser(id uuid.UUID, u UserRequest) (UserResponse, bool) {
	var updateUser = User{Id: id, Username: u.Username}
	user, err := UpdateUser(updateUser)
	if !err {
		return UserResponse{}, false
	}
	return ToUserResponse(user), true
}

func (s *UserService) DeleteUser(id uuid.UUID) bool {
	return Delete(id)
}
