package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDatabase() {
	database, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		panic(err)
	}

	DB = database
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable()
	createUsersTable()
	createRegistrationsTable()
}

func createEventsTable() {
	var query = `
		CREATE TABLE IF NOT EXISTS events (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    name TEXT NOT NULL,
		    description TEXT NOT NULL,
		    location TEXT NOT NULL,
		    datetime DATETIME NOT NULL,
		    user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id)
		);
		`
	_, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}
}

func createUsersTable() {
	var query = `
		CREATE TABLE IF NOT EXISTS users (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);
		`
	_, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}
}

func createRegistrationsTable() {
	var query = `
		CREATE TABLE IF NOT EXISTS registrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_id INTEGER,
			user_id INTEGER,
			FOREIGN KEY(event_id) REFERENCES events(id),
			FOREIGN KEY(user_id) REFERENCES users(id)
		);
		`
	_, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}
}
