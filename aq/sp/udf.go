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

	_,err = db.Exec("call kaboom(:1)", "bad-input")
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Printf("call ok")
	}

 	_,err = db.Exec("call kaboom(:1)", "foo")
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Printf("call ok")
	}


	var status int
	db.QueryRow("select esdbo.tf2(:1) from dual", "foo").Scan(&status)
	log.Printf("fn status: %d",status)

	_,err = db.Exec("call ENQUEUE_AGG_SPEC(:1)", "123:2")
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Printf("enqueue ok")
	}

	var aggSpec string
	var bind = make([]byte, 100)
	err = db.QueryRow("call DEQUEUE_AGG_SPEC(:1)",bind).Scan(&aggSpec)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Printf("fn status: %d", status)
	}
}
