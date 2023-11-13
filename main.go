package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"scrape-us/tags"

	"golang.org/x/net/html"
)

type error interface {
	Error() string
}

func main() {
	fmt.Println("Hello, World!")
	html, err := fetch("https://www.freepik.com/")
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan map[string]string)
	reader := strings.NewReader(html)
	go tokenizer(reader, c)
	for {
		val, open := <-c
		if !open {
			break
		}

		for k, v := range val {
			str := fmt.Sprintf("%s:%s \n,", k, v)

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

func tokenizer(r io.Reader, c chan map[string]string) {

	tokens := html.NewTokenizer(r)
	//  a := []map[string]string{}
	for {
		m := make(map[string]string)
		token := tokens.Next()
		if token == html.ErrorToken {
			// emitToken(tokens.Token())
			if tokens.Err() == io.EOF {
				break
			}
		}

		text, val, err := tagToScrape("img", tokens)
		if err != nil {
			continue
		}
		m[text] = val
		c <- m
	}
	close(c)
}

func writeToFile(s string) {
	file, err := os.OpenFile("scrape-us.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

	fmt.Println("File saved successfully")
}

func tagToScrape(tag string, token *html.Tokenizer) (string, string, error) {
	switch tag {
	case "a":
		return tags.Anchor(token)
	case "img":
		return tags.Image(token)
	}
	return "", "", nil
}
