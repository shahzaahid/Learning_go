package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

var (
	host     = "localhost"
	port     = "5432"
	dbuser   = "postgres"
	password = "shahzahid"
	dbname   = "revision_db"
)

// Define the PostgreSQL connection string
func InitDB() {
	// Define the PostgreSQL connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, dbuser, password, dbname)
	fmt.Println(connStr)
	// Open a database connection
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(fmt.Errorf("error connecting to the database: %v", err))
	}

	createTables()
	createTable()

	// defer DB.Close()
}
func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("could not create users table.")
	}
}

func createTable() {
	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL, 
		description TEXT NOT NULL, 
		dateTime TIMESTAMP NOT NULL, 
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`

	_, err := DB.Exec(createEventTable)
	if err != nil {
		log.Fatal(fmt.Errorf("could not create table: %v", err))
	}
}
