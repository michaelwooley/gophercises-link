package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	fmt.Println("Hello, World")

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
			// doc, err := html.Parse(reader)
			// if err != nil {
			// 	log.Fatal(err)
			// }

			fmt.Println("Starting ", path)
			// ReadNodes(doc, 0)
			links := ExtractLinks(reader)
			fmt.Println(links)
			// z := html.NewTokenizer(reader)
			// ReadTokenizer(z)
			fmt.Println(strings.Repeat("-", 60))
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

// ReadNodes Reads through elements in list.
func ReadNodes(doc *html.Node, depth int) {
	for n := doc.FirstChild; n != nil; n = n.NextSibling {
		if n.Type != html.ElementNode {
			continue
		}
		fmt.Println(strings.Repeat("\t", depth), n.Data, n.Attr)
		ReadNodes(n, depth+1)
	}
}

// ReadTokenizer Reads through the tokenizer
func ReadTokenizer(z *html.Tokenizer) {

	depth := 0
	for {
		tt := z.Next()
		nTab := strings.Repeat("\t", depth)
		switch tt {
		case html.ErrorToken:
			return
		case html.TextToken:
			if depth > 0 {
				// emitBytes should copy the []byte it receives,
				// if it doesn't process it immediately.
				t := z.Text()
				if len(t) > 0 {
					fmt.Printf("%s%s\n", nTab, t)
				}
			}
		case html.StartTagToken, html.EndTagToken:
			fmt.Println(string(z.Raw()))
			tn, hasAttr := z.TagName()
			// k, v, moreAttr := z.TagAttr()
			// fmt.Println(z.Token()) //string(k), string(v), moreAttr)

			// if len(tn) == 1 && tn[0] == 'a' {

			if tt == html.StartTagToken {
				fmt.Printf("%s<%s>%t\n", nTab, tn, hasAttr)
				depth++
			} else {
				depth--
				nTab := strings.Repeat("\t", depth)
				fmt.Printf("%s</%s>\n", nTab, tn)
			}
			// }
		}

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
				// fmt.Println("Text token:", z.Token().Data)
			}
		}
	}
}
