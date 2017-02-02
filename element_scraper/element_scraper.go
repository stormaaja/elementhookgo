package elementscraper

import (
	"io"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type ElementScraper struct {
	Document *goquery.Document
}

func NewElementScraperFromResponse(resp *http.Response) (ElementScraper, error) {
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return ElementScraper{nil}, err
	}

	return ElementScraper{doc}, nil
}

func NewElementScraperFromReader(reader io.Reader) (ElementScraper, error) {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return ElementScraper{nil}, err
	}

	return ElementScraper{doc}, nil
}

func (elementScraper *ElementScraper) Find(element string) string {
	return elementScraper.Document.Find(element).Contents().Text()
}

func (elementScraper *ElementScraper) FindInside(element string, children []string) []map[string]string {
	childrenElements := make([]map[string]string, 0)
	elementScraper.Document.Find(element).Each(func(i int, s *goquery.Selection) {
		row := make(map[string]string)
		found := false
		for _, child := range children {
			row[child] = s.Find(child).Text()
			if len(row[child]) > 0 {
				found = true
			}
		}
		if found {
			childrenElements = append(childrenElements, row)
		}
	})
	return childrenElements
}
