package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

var DB *pgx.Conn

func InitDB() {
	var err error
	connString := "postgresql://postgres:postgresql14@localhost:5432/filemanagers"

	DB, err = pgx.Connect(context.Background(), connString)
	if err != nil {
		panic("Could not connect to database." + err.Error())
	}

	createTables()

}

func createTables() {

	createFoldersTable := `
    CREATE TABLE IF NOT EXISTS folders (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL
    );
    `

	_, err := DB.Exec(context.Background(), createFoldersTable)
	if err != nil {
		log.Fatalf("Error creating events table: %v\n", err)
	}

	createFilesTable := `
    CREATE TABLE IF NOT EXISTS files (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        folder_id INTEGER,
		FOREIGN KEY (folder_id) REFERENCES folders(id)
    );`

	_, err = DB.Exec(context.Background(), createFilesTable)
	if err != nil {
		log.Fatalf("Error creating events table: %v\n", err)
	}

}
