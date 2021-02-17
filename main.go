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

// dates = strings
// commits = ints
var datesAndCommits = make(map[string]int)
// slice for in-order iteration of datesAndCommits
var datesKeys = []string{}

func getContributions(username string) (string, map[string]int) {
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
			var commitCount int

			if t.Data == "rect" {
				for _, a := range t.Attr {
					if a.Key == "data-count" {
						commitCount, _ = strconv.Atoi(a.Val)
					} else if a.Key == "data-date" {
						date = a.Val
						datesAndCommits[date] = commitCount
						datesKeys = append(datesKeys, a.Val)
						break
					}
				}
			}

			// yearlyContributions
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

func getStreak(datesAndCommits map[string]int) int {
	var streak int
	fmt.Println(datesKeys)


	// check if GitHub contributions added an additional date. If so, remove from map
	currentDate := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

	if currentDate == "2021-02-17" {
		fmt.Println("Matches!")
	}

	// check if extra date exists in the map, delete if it does
	if _, val := datesAndCommits[currentDate]; val {
		delete(datesAndCommits, currentDate)
	}
	// fmt.Println(datesAndCommits)

	for _, key := range datesKeys {
		if datesAndCommits[key] > 0 {
			streak++
		} else {
			streak = 0
		}
	}

	return streak
}

func main() {
	yearlyContributions, datesAndCommits = getContributions("nelsonfigueroa")
	fmt.Printf("Commits in the past year: %s \n", color.GreenString(yearlyContributions))
	streak := getStreak(datesAndCommits)
	fmt.Println("Streak: ", streak)
}


