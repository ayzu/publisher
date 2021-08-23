package main

import (
	"log"

	"github.com/alexflint/go-arg"
	"github.com/joho/godotenv"
)

// Registered platforms
const (
	PLATFORM_DEVTO  = "devto"
	PLATFORM_MEDIUM = "medium"
)

// Actions
const (
	PUBLISH = "publish"
	LIST    = "list"
)

// Publisher for a platform.
type Publisher interface {
	MyArticles() ([]Article, error)
	Publish(article Article) (CreatedArticle, error)
	GetUser() (User, error)
}

// Article represents an article.
type Article struct {
	Title     string   `json:"title"`
	Text      string   `json:"text"`
	Tags      []string `json:"tags"`
	Published bool     `json:"published"`
	Series    string   `json:"series"`
	User      string   `json:"user"`
}

// User for a platform.
type User struct {
	ID       string
	Username string
	Name     string
}

// CreatedArticle on platform
type CreatedArticle struct {
	URL string
}

// Command-line args.
type args struct {
	DevtoApiKey string `arg:"env" help:"API key for dev.to platform"`

	MediumApiKey string `arg:"env" help:"API key for medium.io platform"`

	Platform string `arg:"--platform,required" help:"Platform to publish an article. Currently supported medium.io and dev.to"`
	Action   string `arg:"--action,required" help:"Action to perform: list or publish"`
	Filepath string `arg:"--filepath" help:"Filepath to an article for publishing"`
}

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	var a args
	arg.MustParse(&a)

	var p Publisher
	switch a.Platform {
	case PLATFORM_MEDIUM:
		p = NewMediumPublisher(a.MediumApiKey)
	case PLATFORM_DEVTO:
		p = NewDevtoPublisher(a.DevtoApiKey)
	default:
		log.Fatalf("unknown platform %v", a.Platform)
	}

	switch a.Action {
	case LIST:
		articles, err := p.MyArticles()
		if err != nil {
			log.Fatalf("retrieving my articles: %v", err)
		}
		log.Println(articles)
	case PUBLISH:
		article, err := parseFile(a.Filepath)
		if err != nil {
			log.Fatal(err)
		}

		created, err := p.Publish(article)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("successfully published article ", created)
	default:
		log.Fatalf("unrecognized action: %v", a.Action)
	}
}
