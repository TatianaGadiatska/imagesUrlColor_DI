package find_images

import (
	st "../data_struct"
	"github.com/cenkalti/dominantcolor"
	"golang.org/x/image/draw"
	"golang.org/x/net/html"
	"image"
	_ "image/jpeg"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Result struct {
	resultUrlColor []*st.UrlImage
}

func (result_ *Result) MakeUrlColor() []st.UrlImage {
	links := generate(2)

	var resultUrlColor []st.UrlImage
	var result st.UrlImage
	for _, url := range links {
		result.UrlImg = url
		result.Color = findFromUrl(url)
		resultUrlColor = append(resultUrlColor, result)
	}
	log.Print(resultUrlColor)
	return resultUrlColor
}

func generate(workers int) []string {
	//https://wallpaperstock.net/wallpapers_p2.html
	s := "https://wallpaperstock.net"
	var allUrl []string
	ch := make(chan []string)
	for i := 2; i < workers+2; i++ {
		go findLinks(s, ch)
		s = s + "/wallpapers_p" + strconv.Itoa(i) + ".html"
		allUrl = append(allUrl, <-ch...)
	}
	log.Printf("\nСкачено %v ссылок \n", len(allUrl))

	return allUrl
}

func findFromUrl(url string) string {

	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		log.Print(err)
	}
	// создаём пустое изображение для записи необходимого размера
	dst := image.NewRGBA(image.Rect(0, 0, 200, 200))
	// изменение размера
	draw.CatmullRom.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

	return dominantcolor.Hex(dominantcolor.Find(dst))
}

func findLinks(url string, c chan []string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()

	c <- visit(nil, doc)
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "img" {
		for _, a := range n.Attr {
			if a.Key == "src" && strings.Contains(a.Val, "wallpapers/thumbs") {
				links = append(links, "https:"+a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
