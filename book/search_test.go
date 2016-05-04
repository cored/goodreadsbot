package book

import (
	"testing"

	"github.com/cored/goodreadsbot/goodreads"
	"github.com/stretchr/testify/assert"
)

type FakeSearch struct {
	Books []*goodreads.Book
}

func (fs *FakeSearch) Find(text string) []*goodreads.Book {
	return fs.Books
}

func TestSearch(t *testing.T) {
	testcases := []struct {
		Scenario   string
		Text       string
		FakeSearch SearchAdapter
		Expected   []*BookView
	}{
		{
			"Passing a word included in various book titles",
			"hunger",
			&FakeSearch{
				Books: []*goodreads.Book{
					&goodreads.Book{
						ID:       "-1",
						Title:    "The Hunger Games",
						ImageURL: "http://hungergames1.com",
					},
					&goodreads.Book{
						ID:       "-1",
						Title:    "The Hunger Games 2",
						ImageURL: "http://hungergames2.com",
					},
					&goodreads.Book{
						ID:       "-1",
						Title:    "The Hunger Games 3",
						ImageURL: "http://hungergames3.com",
					},
				},
			},
			[]*BookView{
				&BookView{Title: "The Hunger Games", Image: "http://hungergames1.com"},
				&BookView{Title: "The Hunger Games 2", Image: "http://hungergames2.com"},
				&BookView{Title: "The Hunger Games 3", Image: "http://hungergames3.com"},
			},
		},
	}

	for _, testcase := range testcases {
		books := Search(testcase.Text, testcase.FakeSearch)
		assert.Equal(t, testcase.Expected, books, testcase.Scenario)
	}
}
