package models

import (
	"database/sql"
	"errors"
	"time"
)

type SnippetModelInterface interface {
	Insert(title string, content string, expires int) (int, error)
	Get(id int) (Snippet, error)
	Latest() ([]Snippet, error)
	GetIDs() ([]int, error)
	Delete(id int) (Snippet error)
}

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

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	query := `INSERT INTO snippets (title, content, expires)
	VALUES ($1, $2, CURRENT_TIMESTAMP + $3 * INTERVAL '1 day')
	RETURNING id`

	result := m.DB.QueryRow(query, title, content, expires)

	var id int
	err := result.Scan(
		&id,
	)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *SnippetModel) Delete(id int) error {
	query := `DELETE FROM snippets WHERE id = $1`

	_, err := m.DB.Exec(query, id)

	if err != nil {
		return err
	}
	return nil
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

func (m *SnippetModel) GetIDs() ([]int, error) {
	query := `SELECT DISTINCT id FROM snippets ORDER BY id ASC`
	result, err := m.DB.Query(query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []int{}, ErrNoRecord
		}
		return []int{}, err
	}

	var ids []int
	for result.Next() {
		var id int
		err = result.Scan(&id)
		ids = append(ids, id)
	}

	return ids, nil
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
