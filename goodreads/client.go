package goodreads

import "github.com/cored/goodreadsbot/book"

type GoodReadsAdapter struct {
}

func (ga *GoodReadsAdapter) Find(text string) []*book.Book {
	return []*book.Book{}
}
