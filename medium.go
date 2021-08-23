package main

import (
	"log"

	"github.com/Medium/medium-sdk-go"
)

// MediumPublisher - publisher for Medium platform.
type MediumPublisher struct {
	apiKey string
	client *medium.Medium
}

// NewMediumPublisher initiates a new instance.
func NewMediumPublisher(apiKey string) *MediumPublisher {
	client := medium.NewClientWithAccessToken(apiKey)
	return &MediumPublisher{apiKey: apiKey, client: client}
}

// MyArticles lists user articles.
func (m *MediumPublisher) MyArticles() ([]Article, error) {
	u, err := m.GetUser()
	if err != nil {
		log.Fatalf("failed to fetch user: %v", err)
	}
	posts, err := m.client.GetUserPublications(u.ID)
	// TODO: does not work
	if err != nil {
		log.Fatalf("retrieving publications: %v", err)
	}
	if posts == nil {
		log.Fatalf("no publications found")
	}

	var articles []Article
	for _, p := range posts.Data {
		article := Article{
			Title: p.Name,
		}
		articles = append(articles, article)
	}

	return articles, nil
}

// Publish an article.
func (m *MediumPublisher) Publish(article Article) (CreatedArticle, error) {
	u, err := m.GetUser()
	if err != nil {
		log.Fatalf("failed to fetch user: %v", err)
	}
	p, err := m.client.CreatePost(medium.CreatePostOptions{
		UserID:        u.ID,
		Title:         article.Title,
		Content:       article.Text,
		ContentFormat: medium.ContentFormatMarkdown,
		PublishStatus: medium.PublishStatusDraft,
	})
	if err != nil {
		log.Fatal(err)
	}

	return CreatedArticle{URL: p.URL}, nil
}

// GetUser returns current user.
func (m *MediumPublisher) GetUser() (User, error) {
	u, err := m.client.GetUser("")
	if err != nil {
		log.Fatalf("getting user from platform: %v", err)
	}
	return User{
		ID:       u.ID,
		Username: u.Username,
		Name:     u.Name,
	}, nil
}
