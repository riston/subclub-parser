
# Subclub.eu movies list parser

The module has been split into three sub-modules:
- Parser - responsible only for parsing HTML
- Crawler - visit pages and return the result set
- Database insert - save the result set in database

## Usage

```
fmt.Println("Run Database")

db := Connect()
stmt := PrepareSaveQueryStmt(db)

defer stmt.Close()
defer db.Close()

// Visit the first 1-10 pages
for pageSubtitles := range CrawlPages(1, 10) {
    fmt.Println("Result", pageSubtitles)

    SaveCrawlResult(stmt, pageSubtitles)
}
```

## Installation

Make sure you have installed Go v1.5+, package dependencies used:

- github.com/PuerkitoBio/goquery - query from HTML
- github.com/lib/pq - PostgreSQL driver

When using the database module, you have to set the PostgreSQL connection string
in env variable `DATABASE_URL`.

Make sure you have setup the database before inserting data `/db/deploy.sql`

## License

MIT
