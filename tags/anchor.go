package tags

import (
	"fmt"

	"golang.org/x/net/html"
)

func Hello(t html.Tokenizer) (string, string, error) {
	return "", "", nil
}

func Anchor(t *html.Tokenizer) (string, string, error) {
	token := t.Token()
	switch token.Type {
	case html.StartTagToken:
		if token.Type == html.StartTagToken {
			// fmt.Printf("This is a start token---> %s \n", token.Data)
			if len(token.Attr) >= 1 {
				attr := token.Attr
				for _, val := range attr {
					if val.Key == "href" {
						nextToken := t.Next()
						if nextToken == html.TextToken {
							return t.Token().Data, val.Val, nil
						}

					}
				}
			}
		}
	}

	return "", "", fmt.Errorf("No anchor tag found")
}

func Image(t *html.Tokenizer) (string, string, error) {
	token := t.Token()
	switch token.Type {
	case html.StartTagToken, html.SelfClosingTagToken:
		if token.Type == html.SelfClosingTagToken {
			// fmt.Printf("This is a start token---> %s \n", token.Data)
			// fmt.Println(token.Data)
			var src string
			if token.Data == "img" {
				if len(token.Attr) >= 1 {
					attr := token.Attr
					for _, val := range attr {

						if val.Key == "src" {
							src = val.Val
						}
						if val.Key == "alt" && val.Val != "" {
							return val.Val, src, nil
						}
					}
					return "alt", src, nil
				}
			}
		} else {
			return "", "", fmt.Errorf("No Image tag found")
		}

	}
	return "", "", fmt.Errorf("No Image tag found")
}
