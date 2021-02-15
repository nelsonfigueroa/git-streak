package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

func getContributions(username string) string {
	url := "https://github.com/users/" + username + "/contributions"
	var yearlyContributions string

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

			if t.Data == "h2" {
				token = tokenizer.Next()
				t := tokenizer.Token()
				// strings.Fields splits the string s around each instance of
				// one or more consecutive white space characters
				yearlyContributions := strings.Fields(t.Data)[0]
				return yearlyContributions
			}

		case token == html.ErrorToken:
			// end of the document
			return yearlyContributions
		}
	}
}

func main() {
	yearlyContributions := getContributions("nelsonfigueroa8a9dwj")
	fmt.Println("Commits in the past year:", yearlyContributions)
}
