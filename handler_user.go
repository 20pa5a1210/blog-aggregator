package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (apiconfig *APIConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := &parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	const createUser = `
    INSERT INTO users (id, created_at, updated_at, name, api_key)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        encode(sha256(random()::text::bytea), 'hex')
    )
    RETURNING id, created_at, updated_at, name, api_key
    `
	var res User
	err = apiconfig.DB.QueryRow(createUser,
		uuid.New(),
		time.Now().UTC(),
		time.Now().UTC(),
		params.Name,
	).Scan(&res.ID, &res.CreatedAt, &res.UpdatedAt, &res.Name, &res.ApiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, 201, res)
}

func (apiconfig *APIConfig) handleGetUser(w http.ResponseWriter, r *http.Request,user User) {
    respondWithJSON(w, 200, databaseUserToUser(user))
}
