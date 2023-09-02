package entity_test

import (
	"testing"
	"time"

	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewBook(t *testing.T) {
	b, err := entity.NewBook("The Trial", "Franz Kafka", 160, "Schocken", 1925, "978-0805209990")
	assert.Nil(t, err)
	assert.NotNil(t, b)
	assert.Equal(t, b.Title, "The Trial")
}

func TestValidadeBook(t *testing.T) {
	type test struct {
		title      string
		author     string
		pages      int
		publisher  string
		year       int
		isbn       string
		created_at time.Time
		expect     error
	}

	tests := []test{
		{
			title:      "The Trial",
			author:     "Franz Kafka",
			pages:      160,
			publisher:  "Schocken",
			year:       1925,
			isbn:       "978-0805209990",
			created_at: time.Now(),
			expect:     nil,
		},
		{
			title:      "",
			author:     "Franz Kafka",
			pages:      160,
			publisher:  "Schocken",
			year:       1925,
			isbn:       "978-0805209990",
			created_at: time.Now(),
			expect:     entity.ErrInvalidEntity,
		},
		{
			title:      "The Trial",
			author:     "",
			pages:      160,
			publisher:  "Schocken",
			year:       1925,
			isbn:       "978-0805209990",
			created_at: time.Now(),
			expect:     entity.ErrInvalidEntity,
		},
		{
			title:      "The Trial",
			author:     "Franz Kafka",
			pages:      0,
			publisher:  "Schocken",
			year:       1925,
			isbn:       "978-0805209990",
			created_at: time.Now(),
			expect:     entity.ErrInvalidEntity,
		},
		{
			title:      "The Trial",
			author:     "Franz Kafka",
			pages:      160,
			publisher:  "",
			year:       1925,
			isbn:       "978-0805209990",
			created_at: time.Now(),
			expect:     entity.ErrInvalidEntity,
		},
		{
			title:      "The Trial",
			author:     "Franz Kafka",
			pages:      160,
			publisher:  "Schocken",
			year:       0,
			isbn:       "978-0805209990",
			created_at: time.Now(),
			expect:     entity.ErrInvalidEntity,
		},
		{
			title:      "The Trial",
			author:     "Franz Kafka",
			pages:      160,
			publisher:  "Schocken",
			year:       1925,
			isbn:       "",
			created_at: time.Now(),
			expect:     entity.ErrInvalidEntity,
		},
		{
			title:      "The Trial",
			author:     "Franz Kafka",
			pages:      -20,
			publisher:  "Schocken",
			year:       1925,
			isbn:       "978-0805209990",
			created_at: time.Now(),
			expect:     entity.ErrInvalidEntity,
		},
		{
			title:      "The Trial",
			author:     "Franz Kafka",
			pages:      160,
			publisher:  "Schocken",
			year:       -1925,
			isbn:       "978-0805209990",
			created_at: time.Now(),
			expect:     entity.ErrInvalidEntity,
		},
		{
			title:     "The Trial",
			author:    "Franz Kafka",
			pages:     160,
			publisher: "Schocken",
			year:      -1925,
			isbn:      "978-0805209990",
			expect:    entity.ErrInvalidEntity,
		},
	}

	for _, tst := range tests {
		_, err := entity.NewBook(tst.title, tst.author, tst.pages, tst.publisher, tst.year, tst.isbn)
		assert.Equal(t, tst.expect, err)
	}
}
