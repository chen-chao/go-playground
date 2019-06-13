package sql

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
)

func TestSQLQuery(t *testing.T) {
	// open a driver typically not attempt to connect to the database
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/test")
	// check error first
	if err != nil {
		log.Fatal(err)
	}
	// then ensure closing the database
	defer db.Close()

	// test the connection
	if err = db.Ping(); err != nil {
		// cannot connect to the database
		log.Fatal(err)
	}

	// query data
	rows, err := db.Query("select svpsubid, subhash, language, votescore from svpsub_info where svpsubid < 100")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var (
		svpsubid  int
		subhash   string
		language  string
		votescore sql.NullInt64 // to be compatible with SQL `default null`
	)
	// loop through the rows
	for rows.Next() {
		err := rows.Scan(&svpsubid, &subhash, &language, &votescore)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d, %s, %s, %v\n", svpsubid, subhash, language, votescore)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
