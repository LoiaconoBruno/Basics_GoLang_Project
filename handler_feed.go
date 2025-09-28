package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"coursegolang/internal/database"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (apiCfg *apiConfig) handlerCreateFeed(
	w http.ResponseWriter,
	r *http.Request,
	user database.User,
) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if params.Name == "" || params.Url == "" {
		respondWithError(w, http.StatusBadRequest, "Name and url are required")
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		IDFeeds: uuid.New(),
		Name:    params.Name,
		Url:     "",
		UserID:  user.ID,
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
