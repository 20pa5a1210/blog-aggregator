package main

import (
	"context"
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


func (apiconfig *APIConfig) handleGetPosts(w http.ResponseWriter, r *http.Request,user User) {
    posts,err:=GetPostsForUser(apiconfig,r.Context(),GetPostsForUserParams{
        UserID:user.ID,
        Limit:10,
    })
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    respondWithJSON(w, 200, posts)
}


const getPostsForUser = `-- name: GetPostsForUser :many

SELECT posts.id, posts.created_at, posts.updated_at, posts.title, posts.url, posts.description, posts.published_at, posts.feed_id FROM posts
JOIN feed_follows ON feed_follows.feed_id = posts.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2
`

type GetPostsForUserParams struct {
	UserID uuid.UUID
	Limit  int32
}

func  GetPostsForUser(apiconfig *APIConfig,ctx context.Context, arg GetPostsForUserParams) ([]Post, error) {
	rows, err := apiconfig.DB.QueryContext(ctx, getPostsForUser, arg.UserID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Title,
			&i.Url,
			&i.Description,
			&i.PublishedAt,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
