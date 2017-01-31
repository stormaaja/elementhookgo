package element_scraper

import (
  "github.com/PuerkitoBio/goquery"
  "io"
)

type ElementScraper struct {
  Document *goquery.Document
}

func NewElementScraperFromReader(reader io.Reader) (ElementScraper, error) {
  doc, err := goquery.NewDocumentFromReader(reader)
  if err != nil {
    return ElementScraper{nil}, err
  }

  return ElementScraper{doc}, nil
}

func (elementScraper *ElementScraper) Find(element string) (string)  {
  return elementScraper.Document.Find(element).Contents().Text()
}
