package FLV

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/RomainMichau/cloudscraper_go/cloudscraper"
)

type EReciente struct {
	Title   string `json:"title"`
	Episode string `json:"episode"`
	Url     string `json:"url"`
	Img     string `json:"img"`
}

func GetERecientes() []EReciente {

	// Inicializa cloudscraper_go
	client, err := cloudscraper.Init(false, false)
	if err != nil {
		log.Fatal(err)
	}

	// Define la URL a la que acceder
	url := "https://www3.animeflv.net"

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

	var eRecientes []EReciente

	doc.Find("ul.ListEpisodios").Each(func(i int, s *goquery.Selection) {
		s.Find("li").Each(func(j int, li *goquery.Selection) {
			link, _ := li.Find("a.fa-play").Attr("href")
			img, _ := li.Find("img").Attr("src")
			titulo := li.Find("strong").Text()
			epi := li.Find("span.Capi").Text()

			eRecientes = append(eRecientes, EReciente{Title: titulo, Episode: epi, Url: link, Img: url + img})
		})
	})

	return eRecientes
}
