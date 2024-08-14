package handlers

import (
	"htmx-go/internal/app/services"
	"htmx-go/internal/pkg/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	service *services.TodoService
}

func NewTodoHandler(service *services.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) CreateTodoHandler(c *gin.Context) {
	text := c.PostForm("todo")
	todo, err := h.service.CreateTodo(c.Request.Context(), text)
	if err != nil {
		errors.HandleHTTPError(c, err)
		return
	}
	c.HTML(http.StatusOK, "todo.html", todo)
}

func (h *TodoHandler) MarkTodoDoneHandler(c *gin.Context) {
	todoIDStr := c.PostForm("id")
	todoID, err := strconv.Atoi(todoIDStr) // Conve
	if err != nil {
		errors.HandleHTTPError(c, err)
		return
	}

	if err := h.service.MarkTodoDone(c.Request.Context(), todoID); err != nil {
		errors.HandleHTTPError(c, err)
		return
	}
	c.Status(http.StatusOK)
}

func (h *TodoHandler) DeleteTodoHandler(c *gin.Context) {
	todoIDStr := c.PostForm("id")
	todoID, err := strconv.Atoi(todoIDStr) // Conve
	if err != nil {
		errors.HandleHTTPError(c, err)
		return
	}
	if err := h.service.DeleteTodo(c.Request.Context(), todoID); err != nil {
		errors.HandleHTTPError(c, err)
		return
	}
	c.Status(http.StatusOK)
}

func (h *TodoHandler) ListTodosHandler(c *gin.Context) {
	todos, err := h.service.ListTodos(c.Request.Context())
	if err != nil {
		errors.HandleHTTPError(c, err)
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{"todos": todos})
}
