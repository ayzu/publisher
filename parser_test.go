package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var exp = Article{
	Title: "My Wonderful Article",
	Text:  "## And the story goes\n\nas\n\nfollowing\n",
	Tags:  []string{"me", "in", "shiny", "world"},
}

func TestParser(t *testing.T) {
	got, err := parseFile("testdata/article.md")
	require.NoError(t, err)
	require.Equal(t, exp, got)
}

func TestParseText(t *testing.T) {
	inp := `# My Wonderful Article

## Meta

tags: me, in, shiny, world

## And the story goes

as

following
`

	got, err := parseText(inp)
	require.NoError(t, err)
	require.Equal(t, exp, got)
}
