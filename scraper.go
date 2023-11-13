package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

func startScrap(
	db *APIConfig,
	concurrency int,
	timeBetweenFetches time.Duration,
) {
	log.Printf("Starting %d scrapers", concurrency)
	ticker := time.NewTicker(timeBetweenFetches)
	for ; ; <-ticker.C {
		feeds, err := db.DB.QueryContext(
			context.Background(),
			`select * from feeds order BY last_fetched_at desc nulls first limit $1;`,
			int32(concurrency),
		)
		if err != nil {
			log.Println(err)
			continue
		}
		defer feeds.Close()

		var items []Feed
		for feeds.Next() {
			feed := Feed{}
			err := feeds.Scan(
				&feed.ID,
				&feed.CreatedAt,
				&feed.UpdatedAt,
				&feed.Name,
				&feed.Url,
				&feed.UserID,
				&feed.LastFetchedAt,
			)
			if err != nil {
				log.Println(err)
				continue
			}
			items = append(items, feed)
		}

		wg := sync.WaitGroup{}
		for _, feed := range items {
			wg.Add(1)

			go scrapeFeed(db, &wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *APIConfig, wg *sync.WaitGroup, feed Feed) {
	defer wg.Done()
	sqlQuery := `UPDATE feeds
    SET last_fetched_at = NOW(),
    updated_at = NOW()
    WHERE id = $1
    RETURNING id, created_at, updated_at, name, url, user_id, last_fetched_at`
	row := db.DB.QueryRowContext(context.Background(), sqlQuery, feed.ID)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.LastFetchedAt,
	)
    if err != nil {
        log.Println(err)
        return
    }
    rssFeed,err:=urlToFeed(feed.Url)
    if err != nil {
        log.Println(err)
        return
    }
    for _,item:=range rssFeed.Channel.Item{
        description:=sql.NullString{}
        if item.Description != "" {
            description.String = item.Description
            description.Valid = true
        }
        t,err:=time.Parse(time.RFC1123Z,item.PubDate)
        if err != nil {
            log.Println(err)
            continue
        }
        _,err=CreatePost(db,context.Background(),CreatePostParams{
            ID: uuid.New(),
            CreatedAt: time.Now().UTC(),
            UpdatedAt: time.Now().UTC(),
            Title: item.Title,
            Url: item.Link,
            Description: description,
            PublishedAt:t,
            FeedID:feed.ID,
        })
        if err != nil {
            if strings.Contains(err.Error(),"duplicate key value violates unique constraint") {
                continue
            }
            log.Println(err)
        }

    }
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}

const createPost = `-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, created_at, updated_at, title, url, description, published_at, feed_id
`

type CreatePostParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Url         string
	Description sql.NullString
	PublishedAt time.Time
	FeedID      uuid.UUID
}

func CreatePost(db *APIConfig,  ctx context.Context, arg CreatePostParams) (Post, error) {
	row := db.DB.QueryRowContext(ctx, createPost,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Title,
		arg.Url,
		arg.Description,
		arg.PublishedAt,
		arg.FeedID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Url,
		&i.Description,
		&i.PublishedAt,
		&i.FeedID,
	)
	return i, err
}
