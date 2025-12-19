package user

import (
	"testing"

	"github.com/google/uuid"
)

func setupService() *UserService {
	userRepository := &UserRepository{
		users: []User{
			{Id: uuid.New(), Username: "john.go", Password: "password", FirstName: "John", LastName: "Go"},
			{Id: uuid.New(), Username: "jane.go", Password: "123456", FirstName: "Jane", LastName: "Go"},
		},
	}
	return NewUserService(userRepository)
}
func TestGetAllUser(t *testing.T) {
	service := setupService()
	var actual []UserResponse = service.GetAllUser()
	if len(actual) != 2 {
		t.Errorf("expected 2 users, got %d", len(actual))
	}
}

func TestGetUserByID(t *testing.T) {
	service := setupService()
	var users []UserResponse = service.GetAllUser()
	actual, err := service.GetUserByID(users[0].Id)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if users[0].Id != actual.Id {
		t.Errorf("expected user ID %v, got %v", users[0].Id, actual.Id)
	}
}

func TestCreateUser(t *testing.T) {
	service := setupService()
	var user = UserRequest{Username: "Test.Create"}
	actual, err := service.CreateUser(user)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if actual.Id == uuid.Nil {
		t.Errorf("expected user ID to be assigned, got nil")
	}
	users := service.GetAllUser()
	if len(users) != 3 {
		t.Errorf("expected 3 users after adding, got %d", len(users))
	}

	if user.Username != actual.Username {
		t.Errorf("expected username %v, got %v", user.Username, actual.Username)
	}
}
