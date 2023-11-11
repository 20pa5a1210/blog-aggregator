package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiconfig *APIConfig) handleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user User) {
	type parameter struct {
		FeedID string `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := &parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	const createFeedFollow = `
        insert into feed_follows (id, created_at, updated_at, feed_id, user_id)
        values ($1, $2, $3, $4, $5) returning id, created_at, updated_at, feed_id, user_id
    `
	var res FeedFollow
	err = apiconfig.DB.QueryRow(createFeedFollow,
		uuid.New(),
		time.Now().UTC(),
		time.Now().UTC(),
		params.FeedID,
		user.ID,
	).Scan(&res.ID, &res.CreatedAt, &res.UpdatedAt, &res.FeedID, &res.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(res))
}

func (apiconfig *APIConfig) handleGetFeedFollow(w http.ResponseWriter, r *http.Request, user User) {
	const getFeedFollow = `select * from feed_follows where user_id = $1`
	rows, err := apiconfig.DB.Query(getFeedFollow, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var res []FeedFollow
	for rows.Next() {
		var feedFollow FeedFollow
		err := rows.Scan(&feedFollow.ID, &feedFollow.CreatedAt, &feedFollow.UpdatedAt, &feedFollow.FeedID, &feedFollow.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res = append(res, databaseFeedFollowToFeedFollow(feedFollow))
	}
	respondWithJSON(w, 200, res)
}

func (apiconfig *APIConfig) handleUnFollowFeed(w http.ResponseWriter, r *http.Request, user User) {
	paramsID_EN:= chi.URLParam(r, "feedid")
    paramsID,err:= uuid.Parse(paramsID_EN)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid feed ID")
        return
    }
	const deleteFeedFollow = `delete from feed_follows where id = $1 and user_id = $2;`
    _, err = apiconfig.DB.Exec(deleteFeedFollow, paramsID, user.ID)
	if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid feed ID")
		return
	}
	respondWithJSON(w, 200, map[string]string{"status": "success"})
}
