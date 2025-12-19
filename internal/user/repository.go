package user

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/google/uuid"
)

var (
	users []User
	once  sync.Once
)

func LoadUsers() {
	once.Do(func() {
		file, err := os.ReadFile("resources/users.json")
		if err != nil {
			log.Fatal("cannot read users.json:", err)
		}

		if err := json.Unmarshal(file, &users); err != nil {
			log.Fatal("cannot unmarshal users.json:", err)
		}

		log.Printf("users loaded: %d records\n", len(users))
	})
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

func Add(u User) (User, bool) {
	users = append(users, u)
	return u, true
}

func UpdateUser(u User) (User, bool) {
	for i, user := range users {
		if user.Id == u.Id {
			users[i].Username = u.Username
			return users[i], true
		}
	}
	return User{}, false
}

func Delete(id uuid.UUID) bool {
	for i, u := range users {
		if u.Id == id {
			users = append(users[:i], users[i+1:]...)
			return true
		}
	}
	return false
}
