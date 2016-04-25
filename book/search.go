package book

type BookView struct {
	Title string
	Image string
}

type Book struct {
	Id       int
	Title    string
	ImageUrl string
}

type SearchAdapter interface {
	Find(text string) []*Book
}

func Search(text string, searchAdapter SearchAdapter) []*BookView {
	books := searchAdapter.Find(text)
	bookViews := make([]*BookView, len(books), (cap(books)+1)*2)

	for idx, book := range books {
		bookViews[idx] = buildBookViewsFor(book)
	}
	return bookViews
}

func buildBookViewsFor(book *Book) *BookView {
	return &BookView{Title: book.Title, Image: book.ImageUrl}
}
