package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/ahmdrz/divan-e-shams/database"
)

func getDocument(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Bad status code %d", res.StatusCode)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	return doc, err
}

func main() {
	err := database.Open("database.boltdb")
	if err != nil {
		log.Println(err)
		return
	}
	for i := uint(1); i <= 3229; i++ {
		log.Printf("Getting %04d ...", i)
		doc, err := getDocument(fmt.Sprintf("https://ganjoor.net/moulavi/shams/ghazalsh/sh%d/", i))
		if err != nil {
			log.Println(err)
			continue
		}
		htmlResponse := ""
		query := doc.Find(".b p")
		length := query.Length()
		query.Each(func(i int, s *goquery.Selection) {
			htmlResponse += s.Text()
			if i < length-1 {
				htmlResponse += "<br/>"
				if i%2 != 0 {
					htmlResponse += "<br/>"
				}
			}
		})

		poem := &database.Poem{
			ID:      i,
			Content: htmlResponse,
		}
		err = poem.Save()
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
