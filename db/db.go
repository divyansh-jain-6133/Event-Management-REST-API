package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") // Make sure you're assigning to the package-level DB
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	createTables()
}

func createTables() {

	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
    );
	`
	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic("Could not create user table:" + err.Error())
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`

	_, err = DB.Exec(createEventsTable) // <- This line panics if DB is nil
	if err != nil {
		// log.Fatal("Could not create table:", err)
		panic("Could not create table:" + err.Error())
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
	 id INTEGER PRIMARY KEY AUTOINCREMENT,
	 event_id INTEGER,
	 user_id INTEGER,
	 FOREIGN KEY(event_id) REFERENCES events(id),
	 FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("Could not registration table:" + err.Error())
	}

}
