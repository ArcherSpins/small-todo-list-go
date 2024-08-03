package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

type Task struct {
	ID       int     `json:"id"`
	Title    *string `json:"title,omitempty"`
	Resolved *bool   `json:"resolved,omitempty"`
}

var db *sql.DB

func InitDB(dataSourceName string) error {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	return db.Ping()
}

func GetTasks() ([]Task, error) {
	rows, err := db.Query("SELECT id, title, resolved FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Resolved); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTask(id int) (Task, error) {
	var task Task
	err := db.QueryRow("SELECT id, title, resolved FROM tasks WHERE id = $1", id).Scan(&task.ID, &task.Title, &task.Resolved)
	if err != nil {
		if err == sql.ErrNoRows {
			return task, errors.New("task not found")
		}
		return task, err
	}
	return task, nil
}

func CreateTask(task *Task) error {
	err := db.QueryRow(
		"INSERT INTO tasks (title, resolved) VALUES ($1, $2) RETURNING id",
		task.Title, task.Resolved).Scan(&task.ID)
	return err
}

func UpdateTask(id int, task *Task) error {
	var fields []string
	var args []interface{}
	argID := 1

	if task.Title != nil {
		fields = append(fields, fmt.Sprintf("title = $%d", argID))
		args = append(args, *task.Title)
		argID++
	}

	if task.Resolved != nil {
		fields = append(fields, fmt.Sprintf("resolved = $%d", argID))
		args = append(args, *task.Resolved)
		argID++
	}

	if len(fields) == 0 {
		return errors.New("no fields to update")
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE tasks SET %s WHERE id = $%d", strings.Join(fields, ", "), argID)

	_, err := db.Exec(query, args...)
	return err
}

func DeleteTask(id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}
