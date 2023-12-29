package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (Snippet, error) {
	query := `INSERT INTO snippets (title, content, expires)
	VALUES ($1, $2, CURRENT_TIMESTAMP + $3 * INTERVAL '1 day')
	RETURNING id, title, content, created, expires`

	result := m.DB.QueryRow(query, title, content, expires)

	var snip Snippet
	err := result.Scan(
		&snip.ID,
		&snip.Title,
		&snip.Content,
		&snip.Created,
		&snip.Expires,
	)
	if err != nil {
		return Snippet{}, err
	}
	return snip, nil
}

func (m *SnippetModel) Get(id int) (Snippet, error) {
	query := `SELECT * FROM snippets WHERE id = $1 LIMIT 1`
	result := m.DB.QueryRow(query, id)

	var snip Snippet
	err := result.Scan(
		&snip.ID,
		&snip.Title,
		&snip.Content,
		&snip.Created,
		&snip.Expires,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}
	return snip, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	query := `SELECT * FROM snippets ORDER BY created LIMIT 5`
	result, err := m.DB.Query(query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []Snippet{}, ErrNoRecord
		}
		return []Snippet{}, err
	}
	var snips []Snippet
	for result.Next() {
		var snip Snippet
		err = result.Scan(
			&snip.ID,
			&snip.Title,
			&snip.Content,
			&snip.Created,
			&snip.Expires,
		)
		snips = append(snips, snip)
	}

	if err != nil {
		return []Snippet{}, err
	}
	return snips, nil
}
