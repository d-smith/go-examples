package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func deleteData(db *sql.DB) {
	//Clear out any data, then insert some
	_, err := db.Exec("delete from sample")
	if err != nil {
		log.Fatal(err)
	}
}

func insertData(db *sql.DB) {
	stmt, err := db.Prepare("insert into sample(name,value) values(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("Thing1", "Thing2")
	if err != nil {
		log.Fatal(err)
	}
}

func duplicateInsert(db *sql.DB) {
	stmt, err := db.Prepare("insert into sample(name,value) values(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("Thing1", "Thing2")

	log.Println("Error:", err)
	sqlErr := err.(*mysql.MySQLError)
	log.Println("Error code:", sqlErr.Number)
}

func selectData(db *sql.DB) {
	rows, err := db.Query("select name, value from sample")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var name, value string
	for rows.Next() {
		err := rows.Scan(&name, &value)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(name, value)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func selectNonexistentData(db *sql.DB) {
	var value string
	err := db.QueryRow("select value from sample where name = ?", "pointy birds").Scan(&value)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Value is", value)

}

func updateData(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("update sample set value = ? where name = ?")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("A new value, certainly not Thing2", "Thing1")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}

func main() {
	//Create database object
	db, err := sql.Open("mysql",
		"sampleusr:samplepw@tcp(127.0.0.1:3306)/sample")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Verify we can connect via our database object
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	deleteData(db)

	insertData(db)
	duplicateInsert(db)
	selectNonexistentData(db)

	selectData(db)

	updateData(db)

	selectData(db)

	log.Println(db.Stats())

}
