package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Load() {
	con, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Panic("Database error", err)
	}

	_, err = con.Exec(`CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
		author TEXT,
		description TEXT,
		content TEXT,
		created_at TEXT
	)`)

	if err != nil {
		log.Fatal("Post model load error ", err)
	}

	_, err = con.Exec(`CREATE TABLE IF NOT EXISTS projects (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		visit_url TEXT,
		source_url TEXT
	)`)

	if err != nil {
		log.Fatal("Project model load error ", err)
	}

	db = con
	log.Println("Database loaded successfull")
}
