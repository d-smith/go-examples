package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
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

	rows, err := db.Query(`select * from sample`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		log.Println("a row...")
	}
	err = rows.Err()
}
