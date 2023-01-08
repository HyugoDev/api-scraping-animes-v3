package JKanime

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

	client, err := cloudscraper.Init(false, false)
	if err != nil {
		log.Fatal(err)
	}

	// Define la URL a la que acceder
	url := "https://jkanime.net/"

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

	doc.Find("div.maximoaltura").Each(func(i int, s *goquery.Selection) {

		s.Find("a").Each(func(j int, li *goquery.Selection) {
			link, _ := li.Attr("href")
			img, _ := li.Find("img").Attr("src")
			titulo := li.Find("h5").Text()
			epi := li.Find("h6").Text()

			Newlink := strings.Split(link, "https://jkanime.net")

			// fmt.Println(Newlink)

			eRecientes = append(eRecientes, EReciente{Title: titulo, Episode: epi, Url: Newlink[1], Img: img})
		})
	})

	return eRecientes
}
