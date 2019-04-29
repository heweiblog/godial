package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func linkScrape() {
	doc, err := goquery.NewDocument("http://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		linkTag := item
		link, _ := linkTag.Attr("href")
		linkText := linkTag.Text()
		fmt.Printf("Link #%d: '%s' - '%s'\n", index, linkText, link)
	})
}

func main() {
	linkScrape()
}
