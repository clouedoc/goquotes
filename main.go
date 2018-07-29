package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"regexp"
	"strconv"
)

var (
	// here, %s should be remplaced with the desired quote theme
	searchString  string         = "https://www.goodreads.com/quotes/search?q=%s&commit=Search"
	defaultAmount int            = 10
	contentRegexp *regexp.Regexp = regexp.MustCompile("“(.+?)”")
)

// Quote represent a quote object
type Quote struct {
	Content string
	Author  string
}

func (q *Quote) String() string {
	return fmt.Sprintf("%s ― %s\n", q.Content, q.Author)
}

func usage() {
	fmt.Println("Usage: ./goquotes <theme> [amount]")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		usage()
	}

	theme := os.Args[1]
	var amount int = defaultAmount
	var err error

	if len(os.Args) == 3 {
		amount, err = strconv.Atoi(os.Args[2])
		if err != nil {
			usage()
		}
	}

	var quotes []Quote

	c := colly.NewCollector(
		colly.AllowedDomains("www.goodreads.com"),
	)

	// extract all the quotes that are on the page
	c.OnHTML(".quoteDetails", func(e *colly.HTMLElement) {
		res := contentRegexp.FindAllStringSubmatch(e.ChildText("div.quoteText"), -1)

		// it's pretty ugly, but, works ( make sure that we can access
		// that slice's slot )
		if len(res) < 1 {
			return
		}

		if len(res[0]) < 1 {
			return
		}

		quotes = append(quotes, Quote{
			Content: res[0][0],
			Author:  e.ChildText(".authorOrTitle"),
		})

		fmt.Print(".")
	})

	// click next only if we don't have enough quotes
	c.OnHTML(".next_page", func(e *colly.HTMLElement) {
		if len(quotes) < amount {
			e.Request.Visit(e.Attr("href"))
		}
	})

	fmt.Println("Launching Scraper !")
	c.Visit(fmt.Sprintf(searchString, theme))

	fmt.Printf("Scrapped %d quotes.\n\n", len(quotes))

	for _, quote := range quotes {
		fmt.Print(quote.String())
	}
}
