package main

import (
	_ "github.com/mattn/go-oci8"
	"database/sql"
	"log"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println("Open the database")
	db, err := sql.Open("oci8", "esdbo/password@//localhost:1521/xe.oracle.docker")
	fatal(err)

	defer db.Close()

	err = db.Ping()
	fatal(err)

	var fnout string
 	db.QueryRow("select esdbo.test_function() from dual").Scan(&fnout)

	log.Printf("fn out: %s",fnout)

	var status int
	db.QueryRow("select esdbo.tf2(:1) from dual", "foo").Scan(&status)
	log.Printf("fn status: %d",status)

	stmt, err := db.Prepare("")
}
