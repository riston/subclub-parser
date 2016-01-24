package subclub

import (
	"fmt"
	"testing"
)

func TestDatabaseSave(t *testing.T) {
	fmt.Println("Run Database")

	db := ConnectDB()
	stmt := PrepareSaveQueryStmt(db)

	defer stmt.Close()
	defer db.Close()

	for pageSubtitles := range CrawlPages(1, 10) {
		fmt.Println("Result", pageSubtitles)

		SaveCrawlResult(stmt, pageSubtitles)
	}
}
