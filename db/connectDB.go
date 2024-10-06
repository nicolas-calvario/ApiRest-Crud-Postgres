package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectionDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://calvario:root@localhost:15432/DbGo?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connecte to data base")
	return db
}
