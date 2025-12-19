package user

import (
	"github.com/google/uuid"
)

// Business logic

// UserService contains business logic
type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) getAllUser() []UserResponse {
	return ToUserResponseList(s.userRepository.GetAll())
}

func (s *UserService) GetUserByID(id uuid.UUID) (UserResponse, bool) {
	user, found := s.userRepository.GetByID(id)
	if !found {
		return UserResponse{}, false
	}
	return ToUserResponse(user), true
}

func (s *UserService) CreateUser(u UserRequest) (UserResponse, bool) {
	var newUser = User{Id: uuid.New(), Username: u.Username}
	user, err := s.userRepository.Add(newUser)
	if !err {
		return UserResponse{}, false
	}
	return ToUserResponse(user), true
}

func (s *UserService) UpdateUser(id uuid.UUID, u UserRequest) (UserResponse, bool) {
	var updateUser = User{Id: id, Username: u.Username}
	user, err := s.userRepository.UpdateUser(updateUser)
	if !err {
		return UserResponse{}, false
	}
	return ToUserResponse(user), true
}

func (s *UserService) DeleteUser(id uuid.UUID) bool {
	return s.userRepository.Delete(id)
}
