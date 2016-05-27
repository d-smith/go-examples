package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := sql.Open("postgres", "user=esuser dbname=esdb password=password host=localhost sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	log.Println("connected...")

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
