package subclub

import "fmt"
import "testing"
import "io"
import "os"
import "bufio"

const (
	MockPath = "./mock/new-movies.html"
)

func readFile(path string) io.Reader {

	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Failed to read the mock data")
	}

	return bufio.NewReader(f)
}

func TestParser(t *testing.T) {

	reader := readFile(MockPath)

	result, err := ExtractFromFile(reader)
	if err != nil {
		panic("Could not parse the data")
	}

	for _, row := range result {
		fmt.Println(row)
	}
}
