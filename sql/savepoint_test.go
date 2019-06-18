// -*- mode:go;mode:go-playground -*-
// snippet of code @ 2019-06-17 13:46:57

// === Go Playground ===
// Execute the snippet with Ctl-Return
// Remove the snippet completely with its dir and all files M-x `go-playground-rm`

package sql

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	// initialize mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func TestSavepint(t *testing.T) {
	// database test is like
	// +----+------------+--------+
	// | id | name       | number |
	// +----+------------+--------+
	// |  1 | nan        |     10 |
	// |  2 | pineapple  |      5 |
	// |  3 | tomato     |      5 |
	// |  4 | watermelon |      5 |
	// |  5 | peach      |      1 |
	// +----+------------+--------+

	// IGNORES error handling
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("start transaction")
	_, err = db.Exec("savepoint updatetomato")
	if err != nil {
		log.Fatal("start transaction failed:", err)
	}
	var number int
	_ = db.QueryRow("select number from fruit where name=?", "tomato").Scan(&number)
	fmt.Println("original tomato number is", number)

	_ = db.QueryRow("select number from fruit where name=?", "peach").Scan(&number)
	fmt.Println("original peach number is", number)

	// update peach in another goroutine
	go updatePeach(db)

	// let updatePeach run first
	time.Sleep(3 * time.Second)
	if err = updateTomato(db); err != nil {
		log.Fatal("update tomato failed:", err)
	}

	_ = db.QueryRow("select number from fruit where name=?", "tomato").Scan(&number)
	fmt.Println("updated tomato number is", number)

	_ = db.QueryRow("select number from fruit where name=?", "peach").Scan(&number)
	fmt.Println("unknown peach number is", number)

	// Attention: any modification after savepoint will be
	// discarded this causes a concurrency race of
	// non-transactional operations in other goroutines.
	if _, err = db.Exec("rollback to updatetomato"); err != nil {
		log.Fatal("rollback failed")
	}
	db.Exec("commit")

	_ = db.QueryRow("select number from fruit where name=?", "tomato").Scan(&number)

	fmt.Println("tomato number finally is ", number)

	_ = db.QueryRow("select number from fruit where name=?", "peach").Scan(&number)
	fmt.Println("peach number is finally is", number)
	_, err = db.Exec("Update fruit set number=1 where name=?", "peach")
	fmt.Println("set peach number to 1")
}

func updateTomato(db *sql.DB) error {
	_, err := db.Exec("Update fruit set number=1000 where name=?", "tomato")
	return err
}

func updatePeach(db *sql.DB) error {
	_, err := db.Exec("Update fruit set number=-100 where name=?", "peach")
	return err
}
