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
	Search Search `xml:"search"`
}

type Search struct {
	Results Results `xml:"results"`
}

type Results struct {
	Books []Book `xml:"work>best_book"`
}

type Author struct {
	ID   string `xml:"id"`
	Name string `xml:"name"`
}

type Book struct {
	ID       string `xml:"id"`
	Title    string `xml:"title"`
	ImageURL string `xml:"image_url"`
	Author   Author `xml:"author"`
}

const (
	ApiKey  = "ce2k5tJzC59duwqQMvhXwg"
	ApiRoot = "http://www.goodreads.com/"
)

func (ga *GoodReadsAdapter) Find(text string) []Book {
	response := &Response{}
	uri := ApiRoot + "search/index.xml?q=" + text + "&key=" + ApiKey

	getData(uri, response)

	return response.Search.Results.Books
}

func getData(uri string, i *Response) {
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
