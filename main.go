package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/* Add your username here */
const Username = "nelsonfigueroa"

type Commit struct {
	Sha string
	Author Author
}

type Author struct {
	Avatar_url string
}

type Link struct {

}

// func getRepos() {
// get all repos
// for each repo, getCommits()

// eventually this function will take a repo name. repo name should be dynamic, and should be done for every public repo.
// will need another function to get all repo names. public repos only
func getCommits() string {
	urlHalfOne := "https://api.github.com/repos/"
	urlHalfTwo := "/nelsonfigueroa.sh/commits?per_page=1"

	url := urlHalfOne + Username + urlHalfTwo

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	// this is all I really need
	link := resp.Header.Get("Link")
	fmt.Println(link)
	// from here just parse the string somehow to get the count

	// get json response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// deserialize json into struct
	var commit []Commit
	err = json.Unmarshal([]byte(body), &commit)
	if err != nil {
		log.Fatalln(err)
	}


	fmt.Println(commit)

	return "test"
}

// func getCount() {

// }

func main() {
	getCommits()
}
