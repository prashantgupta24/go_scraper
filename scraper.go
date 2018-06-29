package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const LINK = "https://www.cars-data.com/en/sport-cars.html"

func main() {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	// Make HTTP GET request
	response, err := client.Get(LINK)
	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
	}

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	allLinks := make(map[string]int)
	// Find all links and process them with the function
	// defined earlier
	document.Find(".links a").Each(func(index int, element *goquery.Selection) {
		// See if the href attribute exists on the element
		//a := element.Find("a")
		href, _ := element.Attr("href")
		_, ok := allLinks[href]
		if !ok {
			allLinks[href] = 1
			fmt.Println(href)
		}

	})

	// // Copy data from the response to standard output
	// n, err := io.Copy(os.Stdout, response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("Number of bytes copied to STDOUT:", n)
}
