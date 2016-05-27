package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
	"fmt"
	"encoding/json"
)


func insertData(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("insert into sample (name, value, dablob) values ($1, $2, $3)")
	if err != nil {
		return err
	}

	_,err = stmt.Exec(time.Now().String(), "value yes", []byte(jsonDoc))
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

func deleteRecord(db *sql.DB, recno int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("delete from sample where recno = $1")
	if err != nil {
		return err
	}

	_,err = stmt.Exec(recno)
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

func updateRecord(db *sql.DB, recno int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("update sample set value = $1 where recno = $2")
	if err != nil {
		return err
	}

	_,err = stmt.Exec(fmt.Sprintf("updated value for rec %d",recno), recno)
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

func queryAndUpdateData(db *sql.DB) error {
	rows, err := db.Query(`select recno, name, value from sample`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var name, value string
	var recordNo int

	for rows.Next() {
		rows.Scan(&recordNo, &name,&value)
		fmt.Printf("Read row <%d, %s, %s>\n", recordNo, name, value)
		if recordNo % 2 == 0 {
			fmt.Println("-> delete")
			deleteRecord(db, recordNo)
		} else {
			fmt.Println("-> update")
			updateRecord(db, recordNo)
		}


	}
	err = rows.Err()

	return err
}

func queryData(db *sql.DB) error {
	rows, err := db.Query(`select recno, name, value, dablob from sample  order by recno desc`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var name, value string
	var recordNo int
	var dablob []byte

	for rows.Next() {
		rows.Scan(&recordNo, &name,&value,&dablob)
		fmt.Printf("<%d, %s, %s>\n", recordNo, name, value)

		if len(dablob) > 0 {
			fmt.Println("...got some bytes...")
			ms := new(managedService)
			err = json.Unmarshal([]byte(jsonDoc), ms)
			if err != nil {
				fmt.Println("...can't unmarshall da bytes", err.Error())
			} else {
				fmt.Printf("\t%v\n",ms)
			}
		}
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

	fmt.Println("query and update...")
	err = queryAndUpdateData(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("query...")
	err = queryData(db)
	if err != nil {
		log.Fatal(err)
	}
}
