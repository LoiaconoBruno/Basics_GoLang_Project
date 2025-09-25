package main // O el paquete correspondiente

import (
	"encoding/json"
	"fmt"
	"net/http"

	"coursegolang/internal/database"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if params.Name == "" || params.Email == "" {
		respondWithError(w, http.StatusBadRequest, "Name and email are required")
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:    uuid.New(),
		Name:  params.Name,
		Email: params.Email,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code {
			case "23505": // Unique violation
				if pqErr.Constraint == "users_email_key" {
					respondWithError(w, http.StatusConflict, "Email already exists")
				} else {
					respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Database error: %v", err))
				}
			default:
				respondWithError(
					w,
					http.StatusInternalServerError,
					fmt.Sprintf("Database error: %v", err),
				)
			}
			return
		}
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprintf("Error creating user: %v", err),
		)
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
