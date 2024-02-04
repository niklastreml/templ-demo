package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/charmbracelet/log"
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

func Connect(ctx context.Context) error {
	l := log.FromContext(ctx)
	// Connect to the database
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		l.Error("Error while opening database", "error", err)
		return err
	}

	// Check if the connection is successful
	if err := db.PingContext(ctx); err != nil {
		l.Error("Error while checking database connection", "error", err)
		return err
	}

	DB = db
	l.Info("Database connection established")

	return nil
}
