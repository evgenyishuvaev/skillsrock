package sql_storage

import (
	"fmt"
	"skillsrock/internal/storage"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type SQLStorage struct {
	db *sqlx.DB
}

var (
	createTask = `INSERT INTO tasks ("title", "description", "status") VALUES ($1,$2,$3)`
	updateTask = `UPDATE tasks SET title = $1, description = $2, status = $3, updated_at=now() WHERE id = $4`
	deleteTask = `DELETE FROM tasks WHERE id = $1`
	listTask   = `SELECT * from tasks`
)

func (s *SQLStorage) CreateTask(task storage.Task) error {
	_, err := s.db.Exec(createTask, task.Title, task.Description, task.Status)
	if err != nil {
		return fmt.Errorf("Cannot create task: %w", err)
	}
	return nil
}

func (s *SQLStorage) UpdateTask(id string, task storage.Task) error {
	_, err := s.db.Exec(updateTask, task.Title, task.Description, task.Status, id)
	if err != nil {
		return fmt.Errorf("Cannot update task: %w", err)
	}
	return nil
}

func (s *SQLStorage) DeleteTask(id string) error {
	_, err := s.db.Exec(deleteTask, id)
	if err != nil {
		return fmt.Errorf("Cannot delete task: %w", err)
	}
	return nil
}

func (s *SQLStorage) ListTask() ([]storage.Task, error) {
	res := []storage.Task{}
	err := s.db.Select(&res, listTask)
	if err != nil {
		return nil, fmt.Errorf("Cannot get tasks: %w", err)
	}
	return res, nil
}

func NewSQLStorage(dsn string) (*SQLStorage, error) {
	s := &SQLStorage{}
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		err = fmt.Errorf("Can't connect to db: %w", err)
		return s, err
	}

	err = db.Ping()
	if err != nil {
		err = fmt.Errorf("Can't ping db: %w", err)
		return s, err
	}
	s.db = db
	return s, nil
}
