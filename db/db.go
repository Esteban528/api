package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Load() {
	dbPath := os.Getenv("DB_PATH")
	con, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		log.Panic("Database error", err)
	}

	_, err = con.Exec(`CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
		author TEXT,
		title TEXT,
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
		source_url TEXT,
		youtube_url TEXT,
		image_url TEXT
	)`)

	if err != nil {
		log.Fatal("Project model load error ", err)
	}

	_, err = con.Exec(`CREATE TABLE IF NOT EXISTS resources (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		link TEXT,
		image_url TEXT
	)`)

	if err != nil {
		log.Fatal("Resource model load error ", err)
	}

	db = con
	log.Println("Database loaded successfull")
}
