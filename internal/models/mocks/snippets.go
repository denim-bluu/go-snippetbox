package mocks

import (
	"time"

	"snippetbox.joonkang.net/internal/models"
)

type SnippetModel struct{}

var mockSnippet = models.Snippet{
	ID:      1,
	Title:   "Mock Snippet",
	Content: "Mock Content",
	Created: time.Now(),
	Expires: time.Now().Add(time.Hour),
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 2, nil
}

func (m *SnippetModel) Get(id int) (models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return models.Snippet{}, models.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() ([]models.Snippet, error) {
	return []models.Snippet{mockSnippet}, nil
}

func (m *SnippetModel) GetIDs() ([]int, error) {
	return []int{1}, nil
}

func (m *SnippetModel) Delete(id int) error {
	return nil
}
