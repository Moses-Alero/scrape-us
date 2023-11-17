package scraper

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"scrape-us/tags"
	"strings"

	"golang.org/x/net/html"
)

var htmlTags = make(map[string][]string)

func Scrape(url string, tag string, ext string) {
	html, err := fetch(url)
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan map[string]string)
	reader := strings.NewReader(html)
	go tokenizer(reader, tag, c)
	for val := range c {
		for k, v := range val {
			str := fmt.Sprintf("%s:%s \n", strings.TrimSpace(k), strings.TrimSpace(v))

			go writeToFile(str)
		}
	}
	fmt.Println("DONE")
}

func fetch(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return string(body), nil

}

func tokenizer(r io.Reader, tag string, c chan map[string]string) {

	tokens := html.NewTokenizer(r)

	for {
		m := make(map[string]string)
		token := tokens.Next()
		if token == html.ErrorToken {
			if tokens.Err() == io.EOF {
				break
			}
		}

		text, val, err := tagToScrape(tag, tokens)
		if err != nil {
			continue
		}
		m[text] = val
		c <- m
	}
	close(c)
}

func writeToFile(s string) {
	file, err := os.OpenFile("scrape-us.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(s)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func tagToScrape(tag string, token *html.Tokenizer) (string, string, error) {
	switch tag {
	case "a":
		return tags.Anchor(token)
	case "img":
		return tags.Image(token)
	case "p", "big", "strong", "span":
		return tags.TextFormatter(token, tag)
	}

	return "", "", nil

}
