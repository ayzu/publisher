package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// DevtoPublisher - publisher for Devto platform.
type DevtoPublisher struct {
	apiKey string
	client *Client
}

// NewDevtoPublisher constructor.
func NewDevtoPublisher(apiKey string) *DevtoPublisher {
	if apiKey == "" {
		log.Fatal("init DevTo client: apiKey wasn't provided")
	}
	client, err := NewClient("https://dev.to/api")
	if err != nil {
		log.Fatal("init DevTo client: ", err)
	}
	return &DevtoPublisher{apiKey: apiKey, client: client}
}

// MyArticles lists all articles for the current user.
func (dp *DevtoPublisher) MyArticles() ([]Article, error) {
	resp, err := dp.client.GetUserAllArticles(
		context.Background(),
		&GetUserAllArticlesParams{},
		func(ctx context.Context, req *http.Request) error {
			req.Header.Add("api-key", dp.apiKey)
			return nil
		})
	if err != nil {
		return nil, fmt.Errorf("obtaining articles: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}

	var data []ArticleMe
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("decoding server response: %v", err)
	}

	var out []Article
	for _, am := range data {
		out = append(out, articleMeToArticle(am))
	}

	return out, nil
}

func articleMeToArticle(am ArticleMe) Article {
	return Article{
		Title:     am.Title,
		Tags:      am.TagList,
		Published: am.Published,
		Text:      am.BodyMarkdown,
	}
}

// Publish an article.
func (dp *DevtoPublisher) Publish(a Article) (CreatedArticle, error) {
	body := CreateArticleJSONRequestBody{
		Article: &struct {
			BodyMarkdown   *string   `json:"body_markdown,omitempty"`
			CanonicalUrl   *string   `json:"canonical_url,omitempty"`
			Description    *string   `json:"description,omitempty"`
			MainImage      *string   `json:"main_image,omitempty"`
			OrganizationId *int32    `json:"organization_id,omitempty"`
			Published      *bool     `json:"published,omitempty"`
			Series         *string   `json:"series,omitempty"`
			Tags           *[]string `json:"tags,omitempty"`
			Title          *string   `json:"title,omitempty"`
		}{
			BodyMarkdown:   &a.Text,
			CanonicalUrl:   nil,
			Description:    nil,
			MainImage:      nil,
			OrganizationId: nil,
			Published:      &a.Published,
			Series:         nil,
			Tags:           nil,
			Title:          &a.Title,
		},
	}

	resp, err := dp.client.CreateArticle(
		context.Background(),
		body,
		func(ctx context.Context, req *http.Request) error {
			req.Header.Add("api-key", dp.apiKey)
			return nil
		})
	if err != nil {
		return CreatedArticle{}, fmt.Errorf("posting article: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return CreatedArticle{}, fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}

	return CreatedArticle{}, nil
}

func (dp *DevtoPublisher) GetUser() (User, error) {
	return User{}, errors.New("operation is not implmemented for the platform")
}
