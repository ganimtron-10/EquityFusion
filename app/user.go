package app

import "github.com/google/uuid"

type User struct {
	ID string `json:"id"`
}

func RegisterUser() User {
	return User{ID: uuid.NewString()}
}
