package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type CONTACT struct {
	ID                   string
	Name, Address, Phone string
	CreatedAt, UpdatedAt string
}

func main() {
	connStr := "host=127.0.0.1 port=5432 user=admin dbname=admin password=admin sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// INSERT
	var ID int
	id := 2
	name := "Akito"
	address := "Tokyo"
	phone := "090-1111-1111"
	updatedAt := "2019-11-1"
	createdAt := "2014-10-1"
	err = db.QueryRow("INSERT INTO contacts(id, name, address, phone, updatedAt, createdAt) VALUES($1,$2,$3,$4,$5, $6)", id, name, address, phone, updatedAt, createdAt).Scan(&ID)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ID)

	// SELECT
	rows, err := db.Query("SELECT * FROM contacts")
	if err != nil {
		log.Fatal(err)
	}

	var contacts []CONTACT
	for rows.Next() {
		var e CONTACT
		rows.Scan(&e.ID, &e.Name, &e.Address, &e.Phone, &e.CreatedAt, &e.UpdatedAt)
		contacts = append(contacts, e)
	}
	fmt.Printf("%v", contacts)
}
