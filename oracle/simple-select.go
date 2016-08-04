package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-oci8"
	"time"
)

func main() {
	db, err := sql.Open("oci8", "system/oracle@//localhost:1521/xe.oracle.docker")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	if err = testSelect(db); err != nil {
		fmt.Println(err)
		return
	}

	if err = testDate(db); err != nil {
		fmt.Println(err)
		return
	}

}

func testSelect(db *sql.DB) error {
	query := "select 'foo', 3 from dual"
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("rows", rows)
	fmt.Println(rows.Columns())
	for rows.Next() {
		var foo string
		var three int
		rows.Scan(&foo, &three)
		fmt.Println(foo, three)
	}

	return nil
}

func testDate(db *sql.DB) error {
	query := "select systimestamp from dual"
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("rows", rows)
	fmt.Println(rows.Columns())

	for rows.Next() {
		var ts time.Time
		rows.Scan(&ts)
		fmt.Println(ts)
	}

	return nil
}
