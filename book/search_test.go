package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type FakeSearch struct {
	Books []*Book
}

func (fs *FakeSearch) Find(text string) []*Book {
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
				Books: []*Book{
					&Book{
						Id:       -1,
						Title:    "The Hunger Games",
						ImageUrl: "http://hungergames1.com",
					},
					&Book{
						Id:       -1,
						Title:    "The Hunger Games 2",
						ImageUrl: "http://hungergames2.com",
					},
					&Book{
						Id:       -1,
						Title:    "The Hunger Games 3",
						ImageUrl: "http://hungergames3.com",
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
