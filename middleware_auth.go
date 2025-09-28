package main

import (
	"fmt"
	"net/http"

	"coursegolang/internal/auth"
	"coursegolang/internal/database"

	"github.com/lib/pq"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth Erorr: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
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

		handler(w, r, user)
	}
}
