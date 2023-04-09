package main

import (
	"fmt"
	"time"
	// "os"
	"io"
	// "log"
	"net/http"
	"golang.org/x/net/html"
)

var root string = ""

func main() {
	// fmt.Println("Starting...")
	// start := "https://en.wikipedia.org/wiki/Wikipedia"
	// root = start
	// traverse(start)
	var my_set *set = &set{};
	fmt.Println(my_set);
	add(my_set, 1);
	// fmt.Println("contains? ",  contains(my_set, 1));
}

func traverse(url string) {
	fmt.Println("Traversing: " + url)
	resp, _ := http.Get(url)
	links := extractAllLinksFromHtml(resp.Body)
	for _,link := range links {
		time.Sleep(4 * time.Second)
		link = root + link
		fmt.Println("Found " + link)
		traverse(link)
	}
}


func extractAllLinksFromHtml(htmlBody io.Reader) []string {
	links := make([]string, 0)
	tokenizer := html.NewTokenizer(htmlBody)
	for {
		tt := tokenizer.Next()
		if tt == html.ErrorToken {
			// ...
			return links
		} else if tt == html.StartTagToken || tt == html.EndTagToken {
			tkn := tokenizer.Token()
			if tkn.Data == "a" {
				if len(tkn.Attr) > 0  {
					hrefLink := tkn.Attr[0].Val
					if tkn.Attr[0].Key == "href" {
						links = append(links, hrefLink)
					}
				}
			}
		}
		// Process the current token.
	}
}