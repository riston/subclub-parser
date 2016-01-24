package subclub

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Result is a structure of parsed movie data
type Result struct {
	ID           string
	Name         string
	SubName      string
	Views        int
	Author       string
	FPS          float64
	SubtitleLink string
	Links        []string
	Genres       []string
	Created      time.Time
}

func (r Result) String() string {

	return fmt.Sprintf("->%s [%s(%s) - %d - %0.2f by %s %v %s]",
		r.ID, r.Name, r.SubName, r.Views, r.FPS, r.Author, r.Genres, r.Created)
}

func parse(d *goquery.Document) []Result {

	// Select the tales table
	rowsSel := d.Find("#tale_list > tbody:nth-child(2) > tr")

	// var rows []Result
	rows := make([]Result, rowsSel.Length())

	rowsSel.Each(func(i int, s *goquery.Selection) {

		// Get all the rows children td tags
		tdSel := s.Children()

		rows[i] = Result{
			ID:           getMovieID(tdSel.Eq(1)),
			Name:         getMovieName(tdSel.Eq(1)),
			SubName:      getMovieSubName(tdSel.Eq(1)),
			Views:        getViews(tdSel.Eq(4)),
			Author:       getAuthor(tdSel.Eq(8)),
			FPS:          getFPS(tdSel.Eq(6)),
			SubtitleLink: getSubtitleLink(tdSel.Eq(1)),
			Links:        getMovieLinks(tdSel.Eq(3)),
			Genres:       getGenres(tdSel.Eq(2)),
			Created:      getDate(tdSel.Eq(0)),
		}
	})

	return rows
}

func getAuthor(td *goquery.Selection) (author string) {
	author = strings.TrimSpace(td.Text())
	return
}

func getFPS(td *goquery.Selection) (fps float64) {
	fmt.Sscan(td.Text(), &fps)
	return
}

func getViews(td *goquery.Selection) (views int) {
	fmt.Sscan(td.Text(), &views)
	return
}

func getGenres(td *goquery.Selection) (genres []string) {
	genres = strings.Split(td.Text(), ", ")
	return
}

func getDate(td *goquery.Selection) (date time.Time) {
	rawData := strings.TrimSpace(td.Find("font").Last().Text())
	date, err := time.Parse("02.01.2006", rawData)
	if err != nil {
		date = time.Time{}
	}
	return date
}

func getMovieSubName(td *goquery.Selection) (name string) {
	name = td.Find("span.episode_info > b").Text()
	name = strings.TrimSpace(name)
	return
}

func getMovieName(td *goquery.Selection) (name string) {
	// Get link from the anchor tag
	name = td.Find("span a.sc_link").Text()
	name = strings.Join(strings.Fields(name), " ")
	return
}

func getMovieID(td *goquery.Selection) (ID string) {
	link := getSubtitleLink(td)

	// Ge the ID from link
	re := regexp.MustCompile(`\?id=(?P<Id>\d+)$`)
	matches := re.FindStringSubmatch(link)
	ID = matches[len(matches)-1]
	return
}

func getSubtitleLink(td *goquery.Selection) (link string) {
	link = td.Find("span a.sc_link").AttrOr("href", "")
	link = strings.TrimSpace(link)
	return
}

func getMovieLinks(td *goquery.Selection) []string {
	// There are multiple links, use map instead to get all
	return td.Find("a").Map(func(i int, s *goquery.Selection) string {
		return s.AttrOr("href", "")
	})
}

// ExtractFromFile takes the reader as input and parses the Document
func ExtractFromFile(r io.Reader) ([]Result, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, errors.New("Could not parse file")
	}

	return parse(doc), nil
}
