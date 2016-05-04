package book

import "github.com/cored/goodreadsbot/goodreads"

type BookView struct {
	Title string
	Image string
}

type SearchAdapter interface {
	Find(text string) []*goodreads.Book
}

func Search(text string, searchAdapter SearchAdapter) []*BookView {
	books := searchAdapter.Find(text)
	bookViews := make([]*BookView, len(books), (cap(books)+1)*2)

	for idx, book := range books {
		bookViews[idx] = buildBookViewsFor(book)
	}
	return bookViews
}

func buildBookViewsFor(book *goodreads.Book) *BookView {
	return &BookView{Title: book.Title, Image: book.ImageURL}
}
