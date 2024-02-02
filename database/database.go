package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "caas"
)

func ConnectDB() error {
	// Connect to the database
	var connStr string = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	// Check if the connection is successful
	if err := db.Ping(); err != nil {
		return err
	}

	DB = db
	fmt.Println("😁 Connected to database")

	return nil
}
