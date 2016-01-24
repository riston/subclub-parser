package subclub

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	// UserAgent is a header used for requests
	UserAgent = "LoveBot"
	// BaseURL page to get data
	BaseURL = "http://www.subclub.eu/jutud.php?sort=uued&page=%d"
)

func fetchPage(url string) io.Reader {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	// The user agent header should be randomnized
	req.Header.Set("User-Agent", UserAgent)

	if err != nil {
		log.Fatal("Failed to create request")
	}

	response, err := client.Do(req)

	if err != nil {
		log.Fatalf("http get => %v", err.Error())
	}

	return response.Body
}

func parsePage(url string) []Result {
	results, err := ExtractFromFile(fetchPage(url))
	if err != nil {
		log.Fatal("Could not extract page content")
	}

	return results
}

// CrawlPages is to getting pages from subclub as parsed result
func CrawlPages(from, to int) chan []Result {
	len := to - from
	results := make(chan []Result, len)

	go func() {
		for i := from; i <= to; i++ {
			url := fmt.Sprintf(BaseURL, i)
			results <- parsePage(url)
		}

		// Close channel when we are done
		close(results)
	}()

	return results
}
