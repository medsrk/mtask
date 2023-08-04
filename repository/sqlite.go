package repository

import (
	"database/sql"
	"meditasker/domain"
)

type SQLiteTaskRepository struct {
	db *sql.DB
}

func NewSQLiteTaskRepository(db *sql.DB) *SQLiteTaskRepository {
	return &SQLiteTaskRepository{db}
}

func (repo *SQLiteTaskRepository) Store(task domain.Task) error {
	_, err := repo.db.Exec("INSERT INTO tasks (id, description, status, project, entered, due, uuid, urgency) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		task.ID,
		task.Description,
		task.Status,
		task.Project,
		task.Entered,
		task.Due,
		task.UUID,
		task.Urgency,
	)
	if err != nil {
		return err
	}

	return nil
}

func (repo *SQLiteTaskRepository) GetTasks() ([]domain.Task, error) {
	rows, err := repo.db.Query("SELECT id, description, status, project, entered, due, uuid, urgency FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		err := rows.Scan(
			&task.ID,
			&task.Description,
			&task.Status,
			&task.Project,
			&task.Entered,
			&task.Due,
			&task.UUID,
			&task.Urgency,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (repo *SQLiteTaskRepository) Count() (int, error) {
	var count int
	err := repo.db.QueryRow("SELECT COUNT(*) FROM tasks").Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
