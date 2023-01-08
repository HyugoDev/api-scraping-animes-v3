package JKanime

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/RomainMichau/cloudscraper_go/cloudscraper"
)

type Directory struct {
	Title    string `json:"title"`
	Episode  string `json:"episode"`
	Url      string `json:"url"`
	Img      string `json:"img"`
	Date     string `json:"date"`
	Synopsis string `json:"synopsis"`
	State    string `json:"state"`
	Type     string `json:"type"`
}

var directory []Directory

func GetDirectory() []Directory {

	directory = nil

	scrapingDirectory(1)

	return directory
}

func scrapingDirectory(page int) {

	client, err := cloudscraper.Init(false, false)
	if err != nil {
		log.Fatal(err)
	}

	// Define la URL a la que acceder
	url := fmt.Sprintf("https://jkanime.net/directorio/%d/desc/", page)

	// Accede a la página
	res, err := client.Get(url, make(map[string]string), "")
	if err != nil {
		log.Fatal(err)
	}

	html := string(res.Body)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div.custom_flex").Each(func(i int, s *goquery.Selection) {

		s.Find("div.custom_item2").Each(func(j int, li *goquery.Selection) {
			link, _ := li.Find("a").Attr("href")
			img, _ := li.Find("img").Attr("src")
			titulo := li.Find("a").Text()
			epi := strings.Split(li.Find("p.ep").Text(), " ")

			date := li.Find("small.text-muted").Text()
			state := strings.TrimSpace(li.Find("p.card-status").Text())
			typee := strings.TrimSpace(li.Find("p.card-txt").Text())
			synopsis := li.Find("p.synopsis").Text()

			Newlink := strings.Replace(strings.Split(link, "https://jkanime.net")[1], "/", "", -1)

			directory = append(directory, Directory{Title: titulo, Episode: epi[0], Url: Newlink, Img: img, Date: date, Synopsis: synopsis, State: state, Type: typee})
		})

	})

	if nexDirectory(doc) {
		fmt.Println(page)
		page++
		scrapingDirectory(page)
	}

}

func nexDirectory(doc *goquery.Document) bool {

	existe := false

	doc.Find("div.navigation").Each(func(i int, s *goquery.Selection) {

		tamaño := s.Find("a.nav-next").Length()

		if tamaño > 0 {
			existe = true
		}

	})

	return existe
}
