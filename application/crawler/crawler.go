package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

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
	for i := uint(341); i <= 1992; i += 4 {
		wg := &sync.WaitGroup{}
		wg.Add(4)
		go fetchAndSave(i+0, wg)
		go fetchAndSave(i+1, wg)
		go fetchAndSave(i+2, wg)
		go fetchAndSave(i+3, wg)
		wg.Wait()
	}
}

func fetchAndSave(i uint, wg *sync.WaitGroup) {
	defer wg.Done()
	id := fmt.Sprintf("[%04d]", i)

	log.Printf("%s Getting ...", id)
	doc, err := getDocument(fmt.Sprintf("https://ganjoor.net/moulavi/shams/robaeesh/sh%d/", i))
	if err != nil {
		log.Println(id, err)
		return
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

	robaei := &database.Robaei{
		ID:      i,
		Content: htmlResponse,
	}
	err = robaei.Save()
	if err != nil {
		log.Println(id, err)
		return
	}
	log.Printf("%s Getting : done !", id)
}
