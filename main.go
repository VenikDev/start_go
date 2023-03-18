package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	GET  = "GET"
	POST = "POST"
)

const (
	URL_JSON_PLACEHOLDER = "https://jsonplaceholder.typicode.com/posts/"
)

type Article struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	for iter := 0; iter < 10; iter++ {
		go func(iter int) {
			req, err := http.NewRequest(GET, fmt.Sprint(URL_JSON_PLACEHOLDER, iter), nil)
			if err != nil {
				log.Fatal(err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			req = req.WithContext(ctx)

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
		}(iter)
	}
}
