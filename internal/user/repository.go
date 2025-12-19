package user

import (
	"sync"

	"github.com/google/uuid"
)

type UserRepository interface {
	GetAll() []User
	GetByID(id uuid.UUID) (User, bool)
	Add(user User) (User, bool)
	UpdateUser(user User) (User, bool)
	Delete(id uuid.UUID) bool
}

type UserRepositoryImpl struct {
	users []User
	once  sync.Once
}

func (u *UserRepositoryImpl) GetAll() []User {
	return u.users
}

func (u *UserRepositoryImpl) GetByID(id uuid.UUID) (User, bool) {
	for _, u := range u.users {
		if u.Id == id {
			return u, true
		}
	}
	return User{}, false
}

func (u *UserRepositoryImpl) Add(user User) (User, bool) {
	u.users = append(u.users, user)
	return user, true
}

func (u *UserRepositoryImpl) UpdateUser(user User) (User, bool) {
	for i, _user := range u.users {
		if _user.Id == user.Id {
			u.users[i].Username = user.Username
			return u.users[i], true
		}
	}
	return User{}, false
}

func (u *UserRepositoryImpl) Delete(id uuid.UUID) bool {
	for i, _user := range u.users {
		if _user.Id == id {
			u.users = append(u.users[:i], u.users[i+1:]...)
			return true
		}
	}
	return false
}
