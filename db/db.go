package db

import (
	"database/sql"
	"fmt"
	"log"

	// For PostgreSQL connection
	_ "github.com/lib/pq"
)

// Conn exposes global variable referencing the database connection
var Conn *sql.DB

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	name     = "postgres"
)

// InitDB initializes a global connection to PostgreSQL
func InitDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, name)

	var err error
	Conn, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panic(err)
	}

	if err = Conn.Ping(); err != nil {
		log.Panic(err)
	}
}
