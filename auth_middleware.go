package main

import (
	"net/http"

	"github.com/rayhan889/rss-aggr/auth"
	"github.com/rayhan889/rss-aggr/handle_json"
	"github.com/rayhan889/rss-aggr/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apf *ApiConfig) authMiddleware(handler authHandler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			handle_json.RespondWithError(w, 403, err)
			return
		}

		user, err := apf.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			handle_json.RespondWithError(w, 403, err)
			return
		}

		handler(w, r, user)
	}
}