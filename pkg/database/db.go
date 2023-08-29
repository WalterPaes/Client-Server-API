package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Database struct {
	conn *sql.DB
}

func NewDatabaseConnection() *Database {
	conn, err := sql.Open("sqlite3", "exchange.db")
	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		conn: conn,
	}
}

func (db Database) GetConnection() *sql.DB {
	return db.conn
}

func (db Database) Migrate() {
	createTable := `
	  CREATE TABLE IF NOT EXISTS exchanges (
	  uuid TEXT NOT NULL PRIMARY KEY,
	  code TEXT NOT NULL,
	  code_in TEXT NOT NULL,
	  name TEXT NOT NULL,
	  high REAL NOT NULL,
	  low REAL NOT NULL,
	  var_bid REAL NOT NULL,
	  pct_change REAL NOT NULL,
	  bid REAL NOT NULL,
	  ask REAL NOT NULL,
	  timestamp INTEGER NOT NULL,
	  create_date DATETIME NOT NULL,
	  createdAt DATETIME NOT NULL
	  );`

	if _, err := db.conn.Exec(createTable); err != nil {
		log.Fatal("[DATABASE MIGRATION]", err)
	}
	log.Println("[DATABASE MIGRATION] SUCCESS")
}
