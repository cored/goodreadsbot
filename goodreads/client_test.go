package goodreads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	testcases := []struct {
		Scenario string
		Text     string
		Expected []Book
	}{
		{
			"Passing a word included in various book titles",
			"hunger",
			[]Book{},
		},
	}

	goodreads := &GoodReadsAdapter{}

	for _, testcase := range testcases {
		books := goodreads.Find(testcase.Text)
		assert.Equal(t, testcase.Expected, books, testcase.Scenario)
	}
}
