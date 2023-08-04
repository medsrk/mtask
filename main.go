package main

import (
	"database/sql"
	"log"
	"meditasker/cmd"
	"meditasker/domain"
	"meditasker/repository"

	_ "github.com/mattn/go-sqlite3" // Here is the SQLite driver
)

func main() {
	db := initialiseDB()
	defer db.Close()

	// Create a new task manager with a SQLite repository
	taskManager := domain.TaskManager{
		TaskRepo: repository.NewSQLiteTaskRepository(db),
	}

	cmd.TaskManager = &taskManager

	cmd.Execute()
}

func initialiseDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./meditasker.db")
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %q", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		ID TEXT,
		Description TEXT,
		Status TEXT,
		Project TEXT,
		Entered DATETIME,
		Due DATETIME,
		UUID TEXT,
		Urgency REAL
	)`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Error creating table: %q", err)
	}

	return db
}
