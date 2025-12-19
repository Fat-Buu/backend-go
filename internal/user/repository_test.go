package user

import (
	"testing"

	"github.com/google/uuid"
)

func setupRepo() *UserRepositoryImpl {
	return &UserRepositoryImpl{
		users: []User{
			{Id: uuid.New(), Username: "john.go", Password: "password", FirstName: "John", LastName: "Go"},
			{Id: uuid.New(), Username: "jane.go", Password: "123456", FirstName: "Jane", LastName: "Go"},
		},
	}
}
func TestGetAllUser(t *testing.T) {
	repo := setupRepo()
	users := repo.GetAll()

	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}
}

func TestGetByID(t *testing.T) {
	repo := setupRepo()
	users := repo.GetAll()
	user, err := repo.GetByID(users[0].Id)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if users[0].Id != user.Id {
		t.Errorf("expected user ID %v, got %v", users[0].Id, user.Id)
	}
}

func TestAddUser(t *testing.T) {
	repo := setupRepo()
	user := User{Id: uuid.New(), Username: "Test.create", Password: "123456", FirstName: "Test", LastName: "Go"}
	userCreated, err := repo.Add(user)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if userCreated.Id == uuid.Nil {
		t.Errorf("expected user ID to be assigned, got nil")
	}

	users := repo.GetAll()
	if len(users) != 3 {
		t.Errorf("expected 3 users after adding, got %d", len(users))
	}

	if userCreated.Username != user.Username {
		t.Errorf("expected username %v, got %v", user.Username, userCreated.Username)
	}
}

func TestUpdate(t *testing.T) {
	repo := setupRepo()
	users := repo.GetAll()
	user := users[0]
	user.Username = "New.name"
	userUpdated, err := repo.UpdateUser(user)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if userUpdated.Id != user.Id {
		t.Errorf("expected user ID %v, got %v", user.Id, userUpdated.Id)
	}
	if userUpdated.Username != user.Username {
		t.Errorf("expected username %v, got %v", user.Username, userUpdated.Username)
	}
}

func TestDelete(t *testing.T) {
	repo := setupRepo()
	users := repo.GetAll()
	error := repo.Delete(users[0].Id)
	if error != nil {
		t.Fatalf("expected delete to succeed, but it failed")
	}
	users = repo.GetAll()
	if len(users) != 1 {
		t.Errorf("expected 1 user after deleting, got %d", len(users))
	}
}
