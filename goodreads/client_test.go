package goodreads

import (
	"testing"

	"github.com/cored/goodreadsbot/book"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	testcases := []struct {
		Scenario string
		Text     string
		Expected []*book.BookView
	}{
		{
			"Passing a word included in various book titles",
			"hunger",
			[]*book.BookView{
				&book.BookView{Title: "Games", Image: ""},
			},
		},
	}

	goodreads := &GoodReadsAdapter{}

	for _, testcase := range testcases {
		books := book.Search(testcase.Text, goodreads)
		assert.Equal(t, testcase.Expected, books, testcase.Scenario)
	}
}
