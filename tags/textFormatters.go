package tags

import (
	"fmt"

	"golang.org/x/net/html"
)

func TextFormatter(t *html.Tokenizer, tag string) (string, string, error) {
	token := t.Token()

	switch token.Type {
	case html.StartTagToken:
		if token.Data == tag {
			for {
				nextToken := t.Next()
				if nextToken == html.TextToken {
					return tag + "____>", t.Token().Data, nil
				} else if token.Type == html.EndTagToken {
					break
				}
			}
		}
	}
	return "", "", fmt.Errorf("%s tag not found", tag)
}
