package db

import (
	"database/sql"
	"fmt"
	"go-rest-api/configs"

	_ "github.com/mattn/go-sqlite3"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	conn, err := sql.Open("sqlite3", conf.FilePath)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		done BOOLEAN DEFAULT FALSE
	);`

	_, err = conn.Exec(createTableSQL)
	if err != nil {
		return nil, fmt.Errorf("error creating table: %v", err)
	}

	return conn, nil
}
