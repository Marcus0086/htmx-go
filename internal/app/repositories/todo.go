package repositories

import (
	"context"
	"htmx-go/internal/app/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoRepository struct {
	db *pgxpool.Pool
}

func NewTodoRepository(db *pgxpool.Pool) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(ctx context.Context, todo *models.Todo) error {
	err := r.db.QueryRow(ctx, "INSERT INTO todos (text, done) VALUES ($1, false) RETURNING id", todo.Text).Scan(&todo.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) Update(ctx context.Context, todoID int, done bool) error {
	_, err := r.db.Exec(ctx, "UPDATE todos SET done = $1 WHERE id = $2", done, todoID)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) Delete(ctx context.Context, todoID int) error {
	_, err := r.db.Exec(ctx, "DELETE FROM todos WHERE id = $1", todoID)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) GetAll(ctx context.Context) ([]models.Todo, error) {
	rows, err := r.db.Query(ctx, "SELECT id, text, done FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Text, &todo.Done); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
