package main

import (
	"fmt"
	"github.com/briandowns/spinner"
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

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	s.Stop()

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

func getStreak(datesAndCommits map[string]string) (int, string, int) {
	var streak int
	var bestDay string
	var bestDayCount = 0

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
			// keep track of best day
			currentCount, _ := strconv.Atoi(datesAndCommits[key])
			if currentCount > bestDayCount {
				bestDay = key
				bestDayCount = currentCount
			}
			streak++
		}
	}

	return streak, bestDay, bestDayCount
}

func main() {
	username := "nelsonfigueroa"

	// replace username if provided as subcommand
	if len(os.Args) > 1 {
		username = os.Args[1]
	}

	yearlyContributions, datesAndCommits := getContributions(username)
	currentStreak, bestDay, bestDayCount := getStreak(datesAndCommits)

	// commits in the past year
	fmt.Printf("Commits in the past year: %s \n", color.GreenString(yearlyContributions))

	// current streak
	if currentStreak == 0 {
		fmt.Println("Current streak: 0 days.")
	} else {
		// subtract (streak - 1) days from current date
		sinceDate := time.Now().AddDate(0, 0, ((currentStreak - 1) * -1)).Format("2006/01/02")
		fmt.Printf("Current streak: %s \n", color.GreenString(strconv.Itoa(currentStreak))+" day(s), since "+sinceDate)
	}

	// best day
	if bestDayCount > 0 {
		// parse date in YYYY-MM-DD format to change it to YYYY/MM/DD
		parseDate, _ := time.Parse("2006-01-02", bestDay)
		bestDayFormatted := parseDate.Format("2006/01/02")
		fmt.Printf("Best day: %s, with %s commit(s).\n", bestDayFormatted, color.GreenString(strconv.Itoa(bestDayCount)))
	}
}
