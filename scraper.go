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
){
    log.Printf("Starting %d scrapers", concurrency)
    ticker := time.NewTicker(timeBetweenFetches)
    for ; ; <-ticker.C {
        feeds,err:=db.DB.QueryContext(
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

        wg:= sync.WaitGroup{}
        for _,feed := range items{
            wg.Add(1)

            go scrapeFeed(&wg)

        }
    }
}

func scrapeFeed(wg *sync.WaitGroup){
    defer wg.Done()
}
