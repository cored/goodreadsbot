package goodreads

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

type GoodReadsAdapter struct {
}

type Response struct {
	User   User   `xml:"user"`
	Search Search `xml:"search"`
}

type Search struct {
	Books []Book `xml:"work>best_book"`
}

type Author struct {
	ID   string `xml:"id"`
	Name string `xml:"name"`
	Link string `xml:"link"`
}

type Book struct {
	ID       string   `xml:"id"`
	Title    string   `xml:"title"`
	Link     string   `xml:"link"`
	ImageURL string   `xml:"image_url"`
	NumPages string   `xml:"num_pages"`
	Format   string   `xml:"format"`
	Authors  []Author `xml:"authors>author"`
	IBSN     string   `xml:"isbn"`
}

type User struct {
	ID            string       `xml:"id"`
	Name          string       `xml:"name"`
	About         string       `xml:"about"`
	Link          string       `xml:"link"`
	ImageURL      string       `xml:"image_url"`
	SmallImageURL string       `xml:"small_image_url"`
	Location      string       `xml:"location"`
	LastActive    string       `xml:"last_active"`
	ReviewCount   int          `xml:"reviews_count"`
	Statuses      []UserStatus `xml:"user_statuses>user_status"`
	Shelves       []Shelf      `xml:"user_shelves>user_shelf"`
	LastRead      []Review
}

type Shelf struct {
	ID        string `xml:"id"`
	BookCount string `xml:"book_count"`
	Name      string `xml:"name"`
}

type UserStatus struct {
	Page    int    `xml:"page"`
	Percent int    `xml:"percent"`
	Updated string `xml:"updated_at"`
	Book    Book   `xml:"book"`
}

type Review struct {
	Book   Book   `xml:"book"`
	Rating int    `xml:"rating"`
	ReadAt string `xml:"read_at"`
	Link   string `xml:"link"`
}

const (
	ApiKey  = "ce2k5tJzC59duwqQMvhXwg"
	ApiRoot = "http://www.goodreads.com/"
)

func (ga *GoodReadsAdapter) Find(text string) []Book {
	response := &Response{}
	uri := ApiRoot + "search/index.xml?q=" + text + "&key=" + ApiKey

	getData(uri, response)

	return response.Search.Books
}

func getData(uri string, i interface{}) {
	data := getRequest(uri)
	xmlUnmarshal(data, i)
}

func getRequest(uri string) []byte {
	res, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func xmlUnmarshal(b []byte, i interface{}) {
	err := xml.Unmarshal(b, i)
	if err != nil {
		log.Fatal(err)
	}
}
