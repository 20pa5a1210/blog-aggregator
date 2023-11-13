
package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (apiconfig *APIConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request,user User) {
	type parameter struct {
		Name string `json:"name"`
        URL string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := &parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	const createFeed = `
    INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
    VALUES ($1, $2,$3,$4,$5,$6) Returning *;
    `
	var res Feed
	err = apiconfig.DB.QueryRow(createFeed,
		uuid.New(),
		time.Now().UTC(),
		time.Now().UTC(),
		params.Name,
        params.URL,
        user.ID,
	).Scan(&res.ID, &res.CreatedAt, &res.UpdatedAt, &res.Name, &res.Url, &res.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, 201, databaseFeedToFeed(res))
}

func (apiconfig *APIConfig) handleGetFeed(w http.ResponseWriter, r *http.Request) {
    const getFeeds = `select * from feeds`

    rows, err := apiconfig.DB.Query(getFeeds)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var feeds []Feed

    for rows.Next() {
        var feed Feed
        err := rows.Scan(&feed.ID, &feed.CreatedAt, &feed.UpdatedAt, &feed.Name, &feed.Url, &feed.UserID)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        feeds = append(feeds, feed)
    }
    respondWithJSON(w, 200, feeds)
}
