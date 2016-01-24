package subclub

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	// Postgres driver
	_ "github.com/lib/pq"
)

// ConnectDB get database instance
func ConnectDB() (db *sql.DB) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Could not connect to database ", err)
	}
	return
}

// PrepareSaveQueryStmt generate the prepared query statement
func PrepareSaveQueryStmt(db *sql.DB) (stmt *sql.Stmt) {
	query := `
        INSERT INTO
            sub.movies (movie_id, data)
        VALUES
            ($1::text, $2::jsonb)
        ON CONFLICT (movie_id) DO UPDATE SET data=excluded.data`

	// Prepare statement
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal("Could not create a prepared stmt for save ", err)
	}
	return
}

// SaveCrawlResult execute the save operation on the result set
func SaveCrawlResult(stmt *sql.Stmt, result []Result) {
	for _, dataRow := range result {
		// Encode the data part
		encode, err := json.Marshal(dataRow)
		if err != nil {
			log.Fatal("Could not encode struct ", err)
		}

		res, err := stmt.Exec(dataRow.ID, encode)
		if err != nil || res == nil {
			log.Fatal("Could not insert into movies result ", err)
		}
	}
}
