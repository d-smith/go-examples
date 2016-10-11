package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-oci8"
	"log"
	"strings"
	"time"
)

func openAndConnectToDb() *sql.DB {
	log.Println("Open the database")
	db, err := sql.Open("oci8", "user/password@//localhost:1521/xe.oracle.docker")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Ping the db as open might not actually connect")

	//Use a backoff/retry strategy - we can start this client before
	//the database is started, and see it eventually connect and process
	//queries
	var dbError error
	maxAttempts := 20
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		log.Println("ping database...")
		dbError = db.Ping()
		if dbError == nil {
			break
		}
		log.Println("Ping failed: ", dbError, "retry in ", attempts, " seconds.")
		time.Sleep(time.Duration(attempts) * time.Second)
	}
	if dbError != nil {
		log.Fatal(dbError)
	}

	return db
}

func connectionErr(err error) bool {
	errStr := err.Error()
	return strings.HasPrefix(errStr, "ORA-03114") || strings.HasPrefix(errStr, "ORA-03113")
}

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

	var db *sql.DB = openAndConnectToDb()

	log.Println("Do select")
	if err := testSelect(db); err != nil {
		log.Println(err)
	}

	for i := 0; i < 1; i++ {
		if err := testDate(db); err != nil {
			log.Println(err)

			if connectionErr(err) {
				log.Println("Connection to database failed... attempt reconnect.")
				db = openAndConnectToDb()
			}
		}

		time.Sleep(5 * time.Second)
	}

	id := time.Now().Unix()
	err := insertBlobData(db, id)
	fatal(err)

	abc, err := readBlobData(db, id)
	fatal(err)
	fmt.Printf("%v+v\n", abc)

}

func testSelect(db *sql.DB) error {
	query := "select 'foo', 3 from dual"
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var foo string
		var three int
		rows.Scan(&foo, &three)
		log.Println(foo, three)
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

	for rows.Next() {
		var ts time.Time
		rows.Scan(&ts)
		log.Println("systimestamp from dual is ", ts)
	}

	return nil
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

	err = json.Unmarshal(abcBytes, &abc)

	return abc, err
}
