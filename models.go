package main // O el paquete correspondiente

import (
	"time"

	"coursegolang/internal/database"

	"github.com/google/uuid"
)

type User struct {
	ID      uuid.UUID `json:"user-id"`
	Name    string    `json:"user-name"`
	Email   string    `json:"user-email"`
	Created time.Time `json:"user-created-at"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:      dbUser.ID,
		Name:    dbUser.Name,
		Email:   dbUser.Email,
		Created: dbUser.Created,
	}
}
