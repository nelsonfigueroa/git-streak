package main

import (
	"fmt"
	"log"
	"net/http"
	// "strings"
	"io/ioutil"
	// "golang.org/x/net/html"
)

/* Add your username here */
const Username = "nelsonfigueroa"


func getContributions() string {
	url := "https://github.com/users/" + Username + "/contributions"

	response, err := http.Get(url)
	if err != nil {
			log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("HTML:\n\n", string(body))

	return "test"
}

func main() {
	getContributions()
}
