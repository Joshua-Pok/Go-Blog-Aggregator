package rss

import (
	"context"
	"encoding/xml"
	"github.com/Joshua-Pok/BlogAggregator/internal/config"
	"github.com/Joshua-Pok/BlogAggregator/internal/database"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"time"
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

func Addfeed(s *config.State, name string, url string) (database.Feed, error) {
	username := s.Cfg.Current_user_name

	user, err := s.Db.GetUser(context.Background(), username)
	if err != nil {
		return database.Feed{}, err
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.Db.CreateFeed(context.Background(), params)
	if err != nil {
		return database.Feed{}, err
	}

	return feed, nil
}

func ListFeeds(s *config.State) ([]database.Feed, error) {

	feeds, err := s.Db.GetFeeds()
	if err != nil {
		return nil, err
	}

	return feeds, nil

}

func follow(s *config.State, url string) error {

	user, err := s.Db.GetUser(context.Background(), s.Cfg.Current_user_name)

	feed, err := s.Db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	_, err = s.Db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}
	return nil

}

// Return all feed follows for given user, include name of feeds and user in result

func following() ([]database.Feed, error) {}
