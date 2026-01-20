package db

import (
	"context"
	"time"
)

type Task struct {
	ID          int32     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

var TaskRepository = Task{}

// list all tasks
func (t Task) ReadTasks() ([]Task, error) {
	var tasks []Task

	query := `Select * 
				from tasks 
				order by created_at DESC 
				limit 10`

	rows, err := DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// get task by id
func (t Task) ReadTask(id int) (*Task, error) {
	var task Task

	query := `
		Select * 
		from tasks 
		where id = $1
	`

	err := DB.QueryRow(context.Background(), query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

type PostTaskPayload struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status"`
}

func (t *Task) AddTask(payload PostTaskPayload) (int, error) {
	var id int

	query := `
		Insert into tasks (title, description, status) 
		VALUES ($1, $2, $3) 
		RETURNING id;
	`

	err := DB.QueryRow(context.Background(), query, payload.Title, payload.Description, payload.Status).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

type PatchTaskPayload struct {
	ID          int    `json:"id" binding:"required"`
	Title       string `json:"title" binding:"max=100"`
	Description string `json:"description" binding:"max=1000"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
}

func (t *Task) UpdateTask(payload PatchTaskPayload) error {
	query := `
		UPDATE tasks
		SET title = $1, description = $2, status = $3 
		WHERE id = $4
	`

	_, err := DB.Exec(context.Background(), query, payload.Title, payload.Description, payload.Status, payload.ID)

	return err
}

func (t Task) DeleteTask(id int) error {

	query := `
		Delete from tasks 
		WHERE id = $1
	`
	_, err := DB.Exec(context.Background(), query, id)

	return err
}
