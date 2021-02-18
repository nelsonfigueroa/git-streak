package main

import (
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var yearlyContributions string
var datesAndCommits = make(map[string]string)

// slice for in-order iteration of datesAndCommits
var datesKeys = []string{}

func getContributions(username string) (string, map[string]string) {
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

			// iterate through HTML attributes of <rect>
			var date string
			var commitCount string

			if t.Data == "rect" {
				for _, a := range t.Attr {
					if a.Key == "data-count" {
						commitCount = a.Val
					} else if a.Key == "data-date" {
						date = a.Val
						datesAndCommits[date] = commitCount
						datesKeys = append(datesKeys, a.Val)
						break
					}
				}
			}

			// get yearlyContributions
			if t.Data == "h2" {
				token = tokenizer.Next()
				t := tokenizer.Token()
				// strings.Fields splits the string s around each instance of
				// one or more consecutive white space characters
				yearlyContributions = strings.Fields(t.Data)[0]
			}

		case token == html.ErrorToken:
			// end of the document
			return yearlyContributions, datesAndCommits
		}
	}
}

func getStreak(datesAndCommits map[string]string) int {
	var streak int

	// check if GitHub contributions added tomorrow's date. If so, remove from map
	currentDate := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

	// check if extra date exists in the map, delete if it does
	if _, val := datesAndCommits[currentDate]; val {
		// remove from map
		delete(datesAndCommits, currentDate)
		// remove from keys slice
		datesKeys = datesKeys[:len(datesKeys)-1]
	}

	// count streak
	for _, key := range datesKeys {
		if datesAndCommits[key] == "0" {
			streak = 0
		} else {
			streak++
		}
	}

	return streak
}

func main() {
	username := "nelsonfigueroa"

	if len(os.Args) > 1 {
		username = os.Args[1]
	}

	yearlyContributions, datesAndCommits := getContributions(username)
	streak := getStreak(datesAndCommits)

	fmt.Printf("Commits in the past year: %s \n", color.GreenString(yearlyContributions))

	if streak == 0 {
		fmt.Println("Current streak: 0 days.")
	} else {
		fmt.Printf("Current streak: %s \n", color.GreenString(strconv.Itoa(streak))+" days, since "+time.Now().AddDate(0, 0, (streak*-1)).Format("2006/01/02"))
	}
}
