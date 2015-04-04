package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-oci8"
)

func main() {
	db, err := sql.Open("oci8", "b2bnext/password@localhost:1521/XE")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	if err = testSelect(db); err != nil {
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
