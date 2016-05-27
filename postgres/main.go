package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
	"fmt"
)


func insertData(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("insert into sample (name, value) values ($1, $2)")
	if err != nil {
		return err
	}

	_,err = stmt.Exec(time.Now().String(), "value yes")
	stmt.Close()
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func queryData(db *sql.DB) error {
	rows, err := db.Query(`select name, value from sample`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var name, value string
	for rows.Next() {
		rows.Scan(&name,&value)
		fmt.Printf("<%s, %s>\n", name, value)
	}
	err = rows.Err()

	return err
}

func main() {
	db, err := sql.Open("postgres", "user=esuser dbname=esdb password=password host=localhost sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	log.Println("connected...")

	err = insertData(db)
	if err != nil {
		log.Fatal(err)
	}

	err = queryData(db)
	if err != nil {
		log.Fatal(err)
	}
}
