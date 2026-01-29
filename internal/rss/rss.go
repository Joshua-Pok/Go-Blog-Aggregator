package rss

import (
	"context"
	"encoding/xml"
	"io"
	"log"
	"net/http"
)

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	}
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, feedURL, nil)
	request.Header.Set("User-Agent", "gator") //identify our program to the server

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("err performing get request: %v", err)

	}

	content := &RSSItem{}
	bodyBytes, err := io.ReadAll(response.Body)

	err = xml.Unmarshal(bodyBytes, content)
	if err != nil {
		log.Fatalf("Unable to unmarshal content body: %v", err)
	}

	feed := &RSSFeed{}

	feed.Channel.Item = append(feed.Channel.Item, *content)

	return feed, nil

}
