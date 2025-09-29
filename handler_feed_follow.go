package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"coursegolang/internal/database"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (apiCfg *apiConfig) handlerFeedFollow(
	w http.ResponseWriter,
	r *http.Request,
	user database.User,
) {
	// 📌 Estructura esperada en el body
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	var params parameters
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid JSON: %v", err))
		return
	}

	// 📌 Validación de datos
	if params.FeedID == uuid.Nil {
		respondWithError(w, http.StatusBadRequest, "feed_id is required")
		return
	}

	// 📌 Insertar en DB
	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		IDFeedsFollow: uuid.New(),
		UserID:        user.ID,
		FeedID:        params.FeedID,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code {
			case "23505": // unique violation
				respondWithError(w, http.StatusConflict, "You are already following this feed")
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
			fmt.Sprintf("Error following feed: %v", err),
		)
		return
	}

	// 📌 Éxito
	respondWithJSON(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feedFollow))
}
