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

// slice for in-order iteration of datesAndCommits map
var datesKeys = []string{}

// to add color where needed
var colorize = color.New(color.FgGreen, color.Bold).SprintFunc()

func getContributions(username string) (string, map[string]string) {
	var yearlyContributions string
	var datesAndCommits = make(map[string]string)
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

		switch token {
		case html.StartTagToken:
			t := tokenizer.Token()

			// iterate through HTML attributes of <rect>
			var date string
			var commitCount string

			if t.Data == "rect" {
				// text inside <rect> </rect>
				tokenizer.Next()
				text := string(tokenizer.Text())
				// get the contributions digit (or "No" if no contributions)
				commitCount = strings.Split(text, " ")[0]

				// parse HTML attributes to get date
				for _, a := range t.Attr {
					if a.Key == "data-date" {
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
				// strings.Fields splits the string on one or more consecutive white characters
				yearlyContributions = strings.Fields(t.Data)[0]
			}

		case html.ErrorToken:
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
		if datesAndCommits[key] == "No" {
			streak = 0
		} else {
			// keep track of best day
			currentCount, _ := strconv.Atoi(datesAndCommits[key]) // this *might* need to be updated
			if currentCount > bestDayCount {
				bestDay = key
				bestDayCount = currentCount
			}
			streak++
		}
	}

	return streak, bestDay, bestDayCount
}

func printYearlyContributions(yearlyContributions string) {
	if yearlyContributions == "0" {
		fmt.Println("No commits in the past year")
	} else {
		fmt.Printf("Commits in the past year: %s \n", colorize(yearlyContributions))
	}
}

func printCurrentStreak(currentStreak int) {
	if currentStreak == 0 {
		fmt.Println("Current streak: 0 days.")
	} else {
		// subtract (streak - 1) days from current date to get the correct date
		sinceDate := time.Now().AddDate(0, 0, ((currentStreak - 1) * -1)).Format("2006/01/02")

		// determine whether to use "day" or "days"
		if currentStreak > 1 {
			fmt.Printf("Current streak: %s \n", colorize(strconv.Itoa(currentStreak))+" days, since "+sinceDate)
		} else {
			fmt.Printf("Current streak: %s \n", colorize(strconv.Itoa(currentStreak))+" day, since "+sinceDate)
		}
	}
}

func printBestDay(bestDayCount int, bestDay string) {
	// parse date in YYYY-MM-DD format to change it to YYYY/MM/DD
	parseDate, _ := time.Parse("2006-01-02", bestDay)
	bestDayFormatted := parseDate.Format("2006/01/02")

	// determine whether to use "commit" or "commits"
	if bestDayCount > 1 {
		fmt.Printf("Best day in the past year: %s with %s commits.\n", bestDayFormatted, colorize(strconv.Itoa(bestDayCount)))
	} else if bestDayCount == 1 {
		fmt.Printf("Best day in the past year: %s with %s commit.\n", bestDayFormatted, colorize(strconv.Itoa(bestDayCount)))
	}
}

func main() {
	fmt.Println(`
	 ██████╗ ██╗████████╗    ███████╗████████╗██████╗ ███████╗ █████╗ ██╗  ██╗
	██╔════╝ ██║╚══██╔══╝    ██╔════╝╚══██╔══╝██╔══██╗██╔════╝██╔══██╗██║ ██╔╝
	██║  ███╗██║   ██║       ███████╗   ██║   ██████╔╝█████╗  ███████║█████╔╝ 
	██║   ██║██║   ██║       ╚════██║   ██║   ██╔══██╗██╔══╝  ██╔══██║██╔═██╗ 
	╚██████╔╝██║   ██║       ███████║   ██║   ██║  ██║███████╗██║  ██║██║  ██╗
	 ╚═════╝ ╚═╝   ╚═╝       ╚══════╝   ╚═╝   ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝
		`)
	// replace this with your username if you'd rather have the username hardcoded
	username := "nelsonfigueroa"

	// replace username if provided as subcommand
	if len(os.Args) > 1 {
		username = os.Args[1]
	}

	fmt.Printf("Getting stats for %s\n", username)

	yearlyContributions, datesAndCommits := getContributions(username)
	currentStreak, bestDay, bestDayCount := getStreak(datesAndCommits)

	printYearlyContributions(yearlyContributions)
	printCurrentStreak(currentStreak)
	printBestDay(bestDayCount, bestDay)
}
