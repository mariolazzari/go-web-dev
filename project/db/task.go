package db

import "context"

type Task struct{}

var TaskRepository = Task{}

type PostTaskPayload struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status"`
}

func (t *Task) SaveTaskQuery(payload PostTaskPayload) (int, error) {
	var id int

	query := `Insert into tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id;`

	err := DB.QueryRow(context.Background(), query, payload.Title, payload.Description, payload.Status).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
