package user

import (
	"github.com/google/uuid"
)

// Business logic

type IUserService interface {
	GetAllUser() []UserResponse
	GetUserByID(id uuid.UUID) (UserResponse, error)
	CreateUser(u UserRequest) (UserResponse, error)
	UpdateUser(id uuid.UUID, u UserRequest) (UserResponse, error)
	DeleteUser(id uuid.UUID) error
}

// UserService contains business logic
type UserService struct {
	userRepository IUserRepository
}

func NewUserService(userRepository IUserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetAllUser() []UserResponse {
	return ToUserResponseList(s.userRepository.GetAll())
}

func (s *UserService) GetUserByID(id uuid.UUID) (UserResponse, error) {
	user, err := s.userRepository.GetByID(id)
	if err != nil {
		return UserResponse{}, err
	}
	return ToUserResponse(user), nil
}

func (s *UserService) CreateUser(u UserRequest) (UserResponse, error) {
	var newUser = User{Id: uuid.New(), Username: u.Username}
	user, err := s.userRepository.Add(newUser)
	if err != nil {
		return UserResponse{}, err
	}
	return ToUserResponse(user), nil
}

func (s *UserService) UpdateUser(id uuid.UUID, u UserRequest) (UserResponse, error) {
	var updateUser = User{Id: id, Username: u.Username}
	user, err := s.userRepository.Update(updateUser)
	if err != nil {
		return UserResponse{}, err
	}
	return ToUserResponse(user), nil
}

func (s *UserService) DeleteUser(id uuid.UUID) error {
	return s.userRepository.Delete(id)
}
