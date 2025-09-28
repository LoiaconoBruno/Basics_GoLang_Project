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

	// Decodificar body
	var params parameters
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// Validar campos
	if params.Name == "" || params.Url == "" {
		respondWithError(w, http.StatusBadRequest, "Name and URL are required")
		return
	}

	// Crear feed en DB
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		IDFeeds: uuid.New(),
		Name:    params.Name,
		Url:     params.Url,
		UserID:  user.ID,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code {
			case "23505": // unique violation
				if pqErr.Constraint == "feeds_url_key" { // ✅ constraint correcta
					respondWithError(w, http.StatusConflict, "Feed URL already exists")
				} else {
					respondWithError(w, http.StatusConflict, "Feed already exists")
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
			fmt.Sprintf("Error creating feed: %v", err), // ✅ mensaje corregido
		)
		return
	}

	// Respuesta con feed creado
	respondWithJSON(w, http.StatusCreated, databaseFeedToFeed(feed)) // ✅ 201 Created
}

func (apiCfg *apiConfig) handlerGetFeed(
	w http.ResponseWriter,
	r *http.Request,
) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprintf("an Error has occured, %v", err),
		)
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}
