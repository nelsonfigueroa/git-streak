package main

import (
	"fmt"
	"golang.org/x/net/html"
	"github.com/fatih/color"
	"log"
	"net/http"
	"os"
	"strings"
)

var yearlyContributions string

func getContributions(username string) string {
	url := "https://github.com/users/" + username + "/contributions"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if response.Status == "404 Not Found" {
		fmt.Println("404 Not Found. Are you sure the username exists?")
		os.Exit(1)
	}

	tokenizer := html.NewTokenizer(response.Body)

	for {
		token := tokenizer.Next()

		switch {
		case token == html.StartTagToken:
			t := tokenizer.Token()

			// trying to get data-count for streak 
			if t.Data == "rect" {
				// fmt.Println("in rect")
				for _, a := range t.Attr {
					if a.Key == "data-count" {
						// fmt.Println("Found data-count:", a.Val)
						break
					}
				}
			}

			if t.Data == "h2" {
				token = tokenizer.Next()
				t := tokenizer.Token()
				// strings.Fields splits the string s around each instance of
				// one or more consecutive white space characters
				yearlyContributions = strings.Fields(t.Data)[0]
			}

		case token == html.ErrorToken:
			// end of the document
			return yearlyContributions
		}
	}
}

func main() {
	yearlyContributions := getContributions("nelsonfigueroa")
	fmt.Printf("Commits in the past year: %s \n", color.GreenString(yearlyContributions))
}
