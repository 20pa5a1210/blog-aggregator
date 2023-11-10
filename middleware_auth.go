package main

import (
	"net/http"

	"github.com/20pa5a1210/blog-aggregator/internal/auth"
)

type authedHandler func(http.ResponseWriter, *http.Request, User)

func (apiconfig *APIConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
			return
		}

        const getUser = `
        SELECT id, created_at, updated_at, name, api_key
        FROM users
        WHERE api_key = $1
        `
		var res User
		err = apiconfig.DB.QueryRow(getUser, apiKey).Scan(&res.ID, &res.CreatedAt, &res.UpdatedAt, &res.Name, &res.ApiKey)
		if err != nil {
			respondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
			return
		}
        handler(w, r, res)
	}
}
