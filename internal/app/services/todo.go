package services

import (
	"context"
	"htmx-go/internal/app/models"
	"htmx-go/internal/app/repositories"
)

type TodoService struct {
	repo *repositories.TodoRepository
}

func NewTodoService(repo *repositories.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) CreateTodo(ctx context.Context, text string) (*models.Todo, error) {
	todo := &models.Todo{Text: text, Done: false}
	if err := s.repo.Create(ctx, todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *TodoService) MarkTodoDone(ctx context.Context, todoID int) error {
	return s.repo.Update(ctx, todoID, true)
}

func (s *TodoService) DeleteTodo(ctx context.Context, todoID int) error {
	return s.repo.Delete(ctx, todoID)
}

func (s *TodoService) ListTodos(ctx context.Context) ([]models.Todo, error) {
	return s.repo.GetAll(ctx)
}
