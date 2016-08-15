package main

import (
	"database/sql"

	_ "gopkg.in/rana/ora.v3"
	"log"
	"fmt"
	"time"
)

func main() {
	db, err := sql.Open("ora", "system/oracle@//localhost:1521/xe.oracle.docker")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = testDate(db)
	if err != nil {
		log.Fatal(err.Error())
	}
}


func testDate(db *sql.DB) error {
	query := "select systimestamp from dual"
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()


	for rows.Next() {
		var ts time.Time
		rows.Scan(&ts)
		fmt.Println("current timestamp", ts)
	}

	return nil
}