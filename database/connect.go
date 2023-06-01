package database

import (
	"database/sql"
	"log"
	"os"
	"time"
)

var counts int64

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ConnectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Waiting for DB to start...")
			counts++
		} else {
			log.Println("DB started")
			return connection
		}

		if counts > 10 {
			log.Fatal("DB never started")
			return nil
		}
		log.Println("Backing off for 2 seconds")
		time.Sleep(2 * time.Second)
		continue
	}
}
