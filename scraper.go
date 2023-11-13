package main

import (
	"context"
	"log"
	"sync"
	"time"
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
        log.Println("Found post", item.Title)
    }
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
