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
