package main

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type CONTACT struct {
	ID                   string
	Name, Address, Phone string
	CreatedAt, UpdatedAt string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	connStr := "host=%s port=%s user=%s dbname=%s password=%s sslmode=%s"
	connStr = fmt.Sprintf(connStr, os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"), os.Getenv("SSL_MODE"))

	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	// INSERT
	var ID int
	id := 8
	name := "Futu"
	address := "Tokyo"
	phone := "090-2223-2222"
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
