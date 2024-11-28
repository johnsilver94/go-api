package types

import (
	"time"

	"github.com/google/uuid"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id string) (*User, error)
	CreateUser(user User) error
}

type TodoStore interface {
	GetTodoByID(todoID string) (*Todo, error)
	CreateTodo(todo Todo) error
	GetTodos() ([]*Todo, error)
}

type RegisterUserPayload struct {
	Email    string `json:"email"  validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,min=3,max=100"`
}

type LoginUserPayload struct {
	Email    string `json:"email"  validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=100"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Todo struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	UserID      uuid.UUID `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTodoPayload struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Completed   bool      `json:"completed"`
	UserID      uuid.UUID `json:"user_id" validate:"required"`
}
