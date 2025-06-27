package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func store(username string, pass string) {
	db, err := sql.Open("sqlite", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Insert username and password in a single row
	_, err = db.Exec("INSERT INTO users(username, password) VALUES(?, ?)", username, pass)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("New user inserted successfully")
}
