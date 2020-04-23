package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host = "localhost"
	port = 5432
	user = "nathanielrand" 
	password = "" 
	dbname = "goblog"
)

func main() {
	// A connection string that we use to connect to our database.
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)
	
	// Open a database connection.
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	var id int
	var first_name, email string 
	
	rows, err := db.Query(`
		SELECT id, first_name, email FROM users
		WHERE email=$1
		OR id >$2`,
		"nathanieljrand@gmail.com", 3)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		rows.Scan(&id, &first_name, &email)
		fmt.Println("ID:", id, "Name:", first_name, "Email:", email)
	}
	
	// Ping the database.
	err = db.Ping() 
	if err != nil {
		panic(err) 
	}

	// Print success message to terminal.
	fmt.Println("Successfully connected!")

	// Close database connection.
	db.Close()
}