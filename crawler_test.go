package subclub

import (
	"fmt"
	"testing"
)

func TestCrawler(t *testing.T) {
	fmt.Println("Run Crawler")

	for result := range CrawlPages(1, 3) {
		fmt.Println("Result", result)
	}
}
