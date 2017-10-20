package main

import (
	"database/sql"
	"log"
	"time"
	"fmt"
	_ "github.com/mattn/go-oci8"
	"os"
)

func openAndConnectToDb() *sql.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	log.Println("Open the database")
	db, err := sql.Open("oci8", user + "/" + password + "@//localhost:1521/xe.oracle.docker")
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

func main() {

	var db *sql.DB = openAndConnectToDb()

	query := `select aggregate_id, version from t_aeev_events where id = (select max(id) from t_aeev_events)`

	for {
		var aggregateID string
		var version int

		err := db.QueryRow(query).Scan(&aggregateID, &version)
		if err != nil {
			fmt.Printf("Error scanning for aggregate and version: %s\n", err.Error())
		}

		log.Printf("Last event: %s %d", aggregateID, version)


		var d string
		err = db.QueryRow("select dummy from dual").Scan(&d)
		if err != nil {
			fmt.Printf("Error scanning dual: %s\n", err.Error())
		}

		time.Sleep(time.Second)

	}

}

