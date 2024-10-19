package models

import (
	"context"
	"go-explorer/db"
)

type Folder struct {
	ID   int64
	Name string `binding:"required"`
}

type File struct {
	ID        int64
	Name      string `binding:"required"`
	Folder_id int64  `binding:"required"`
}

func GetAllFolders() ([]Folder, error) {
	query := "select id, name from folders"
	rows, err := db.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Folder

	for rows.Next() {
		var event Folder
		err := rows.Scan(&event.ID, &event.Name)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetFileById(id int64) ([]File, error) {
	query := "select * from files where folder_id = $1"
	rows, err := db.DB.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []File

	for rows.Next() {
		var event File
		err := rows.Scan(&event.ID, &event.Name, &event.Folder_id)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
