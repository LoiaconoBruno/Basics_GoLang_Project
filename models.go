package main // O el paquete correspondiente

import (
	"time"

	"coursegolang/internal/database"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"user-id"`
	Name      string    `json:"user-name"`
	Email     string    `json:"user-email"`
	CreatedAt time.Time `json:"user-created-at"`
}

func databaseUserToUser(dbUser database.User) User {
	createdAt := time.Time{}
	if dbUser.CreatedAt.Valid {
		createdAt = dbUser.CreatedAt.Time
	}
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		Email:     dbUser.Email,
		CreatedAt: createdAt,
	}
}
