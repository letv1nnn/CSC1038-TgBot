package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

type LabTask struct {
	ID       int
	TaskName string
	Code     string
}

func New(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not connect database: %w", err)
	}

	schema := `
    CREATE TABLE IF NOT EXISTS labs (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        task_name TEXT NOT NULL,
        code TEXT NOT NULL
    );`

	if _, err := db.Exec(schema); err != nil {
		return nil, fmt.Errorf("could not create table: %w", err)
	}

	return &Storage{db}, nil
}

func (s *Storage) Save(lab LabTask) error {
	_, err := s.db.Exec(
		"INSERT INTO labs (task_name, code) VALUES (?, ?)",
		lab.TaskName, lab.Code,
	)
	if err != nil {
		return fmt.Errorf("could not insert lab: %w", err)
	}
	return nil
}

func (s *Storage) Get(task_name string) (*LabTask, error) {
	row := s.db.QueryRow(
		"SELECT id, task_name, code FROM labs WHERE task_name = ?",
		task_name,
	)
	var lab LabTask
	if err := row.Scan(&lab.ID, &lab.TaskName, &lab.Code); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not scan lab: %w", err)
	}
	return &lab, nil
}
