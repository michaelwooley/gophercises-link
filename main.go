package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	err := filepath.Walk("./testcases",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			} else if info.IsDir() {
				return nil
			}
			s, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatal(err)
			}
			reader := strings.NewReader(string(s))

			log.Printf("Starting: %s", path)
			links := ExtractLinks(reader)
			log.Printf("Result: %v", links)
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
}

// Link Html <a> tag href + inner text
type Link struct {
	Href string
	Text string
}

// ExtractLinks Extract all <a> tags w/ inner text from HTML doc.
func ExtractLinks(reader *strings.Reader) []Link {

	z := html.NewTokenizer(reader)
	var links []Link
	var link Link
	isOpen := false

	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if token.Data == "a" {
				if tt == html.StartTagToken {
					isOpen = true
					for _, attr := range token.Attr {
						if attr.Key == "href" {
							link = Link{Href: attr.Val, Text: ""}
						}
					}
				} else {
					isOpen = false
					links = append(links, link)
				}
			}
		case html.TextToken:
			if isOpen {
				link.Text += z.Token().Data
			}
		}
	}
}
