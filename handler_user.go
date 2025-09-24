package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"coursegolang/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name  string `json:name`
		Email string `json:email`
	}

	decode := json.NewDecoder(r.Body)
	params := parameters{}

	err := decode.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		Name:  params.Name,
		Email: params.Email,
		CreatedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Creating the user: %v", err))
		return
	}

	respondWithJSON(w, 200, user)
}
