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
    var res User
    err = apiconfig.DB.QueryRow("INSERT INTO users (id,created_at,updated_at,name) VALUES ($1,$2,$3,$4) RETURNING *",
                                     uuid.New(),
                                     time.Now().UTC(),
                                     time.Now().UTC(),
                                    params.Name,
                                ).Scan(&res.ID,&res.CreatedAt,&res.UpdatedAt,&res.Name)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    respondWithJSON(w, http.StatusOK, res)
}
