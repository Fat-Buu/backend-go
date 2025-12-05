package user

import "github.com/google/uuid"

// Data layer / mock DB
var users = []User{
	{Id: uuid.New(), Username: "John Go"},
	{Id: uuid.New(), Username: "Jane Fiber"},
}

// Get all users
func GetAll() []User {
	return users
}

func GetByID(id uuid.UUID) (User, bool) {
	for _, u := range users {
		if u.Id == id {
			return u, true
		}
	}
	return User{}, false
}

func Add(u User) {
	users = append(users, u)
}
