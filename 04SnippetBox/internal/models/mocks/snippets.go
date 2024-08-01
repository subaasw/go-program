package mocks

import (
	"snippetbox-app/internal/models"
	"time"
)

var mockSnippets = &models.Snippet{
	ID:      1,
	Title:   "An old silent pond",
	Content: "An old silent pod...",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 2, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippets, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippets}, nil
}
