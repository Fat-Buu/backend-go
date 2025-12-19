package user

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func NewUserRepositoryFromFile(path string) (*UserRepository, error) {
	var users []User
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read users file: %w", err)
	}

	if err := json.Unmarshal(file, &users); err != nil {
		return nil, fmt.Errorf("unmarshal users: %w", err)
	}

	log.Printf("users loaded: %d records\n", len(users))
	return &UserRepository{
		users: users,
	}, nil
}
