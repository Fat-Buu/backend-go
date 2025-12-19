package user

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type IUserRepository interface {
	GetAll() []User
	GetByID(id uuid.UUID) (User, error)
	Add(user User) (User, error)
	Update(user User) (User, error)
	Delete(id uuid.UUID) error
}

type UserRepository struct {
	users []User
	once  sync.Once
}

func (u *UserRepository) GetAll() []User {
	copied := make([]User, len(u.users))
	copy(copied, u.users)
	return copied
}

func (u *UserRepository) GetByID(id uuid.UUID) (User, error) {
	for _, u := range u.users {
		if u.Id == id {
			return u, nil
		}
	}
	return User{}, errors.New("User " + id.String() + " not found")
}

func (u *UserRepository) Add(user User) (User, error) {
	if user.Id == uuid.Nil {
		user.Id = uuid.New()
	}
	u.users = append(u.users, user)
	return user, nil
}

func (u *UserRepository) Update(user User) (User, error) {
	for i, _user := range u.users {
		if _user.Id == user.Id {
			u.users[i] = user
			return u.users[i], nil
		}
	}
	return User{}, errors.New("User " + user.Id.String() + " not found")
}

func (u *UserRepository) Delete(id uuid.UUID) error {
	for i, _user := range u.users {
		if _user.Id == id {
			u.users = append(u.users[:i], u.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user " + id.String() + " not found")
}
