package main

import (
	"database/sql"

	"coursegolang/internal/database"
)

type User struct {
	ID        int32        `json:"user-id"`
	Name      string       `json:"user-name"`
	Email     string       `json:"user-email"`
	CreatedAt sql.NullTime `json:"user-created-at"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:    dbUser.ID,
		Name:  dbUser.Name,
		Email: dbUser.Email,
		CreatedAt: sql.NullTime{
			Time:  dbUser.CreatedAt.Time,
			Valid: dbUser.CreatedAt.Valid,
		},
	}
}
