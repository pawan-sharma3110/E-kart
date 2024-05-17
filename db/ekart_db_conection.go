package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DbInIt() *sql.DB {
	connStr := `host=localhost port=8080 user=postgres dbname=E-kart password=Pawan@2003 sslmode=disable`
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Something went wrong in connection string %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("You are not conect with database %v", err)
	}
	return db
}
