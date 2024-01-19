package db

import (
	"database/sql"
	"fmt"

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
		panic(fmt.Errorf("error connecting to the database: %v", err))
	}

}
