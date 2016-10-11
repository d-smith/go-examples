package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "gopkg.in/rana/ora.v3"
	"log"
	"time"
)

type ABC struct {
	A string
	B int
	C string
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := sql.Open("ora", "user/password@//localhost:1521/xe.oracle.docker")
	fatal(err)

	defer db.Close()

	err = db.Ping()
	fatal(err)

	id := time.Now().Unix()
	err = insertBlobData(db, id)
	fatal(err)

	abc, err := readBlobData(db, id)
	fatal(err)
	fmt.Printf("%v+v\n", abc)
}


func insertBlobData(db *sql.DB, id int64) error {
	abc := ABC{
		A: "a part", B: 42, C: "you see",
	}

	abcBytes, err := json.Marshal(&abc)
	if err != nil {
		return err
	}

	_, err = db.Exec("insert into blob_sample (id, payload) values (:1,:2)", id, abcBytes)
	return err
}

func readBlobData(db *sql.DB, id int64) (ABC, error) {
	var abc ABC

	rows, err := db.Query("select payload from blob_sample where id = :1", id)
	if err != nil {
		return abc, err
	}

	defer rows.Close()

	var abcBytes []byte

	for rows.Next() {
		err := rows.Scan(&abcBytes)
		if err != nil {
			return abc, err
		}
	}

	if err = rows.Err(); err != nil {
		return abc,err
	}


	err = json.Unmarshal(abcBytes, &abc)

	return abc, err
}
