package todo

import (
	"database/sql"

	"github.com/johnsilver94/go-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateTodo(todo types.Todo) error {
	_, err := s.db.Exec("INSERT INTO todos (id, title, description, completed, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)", todo.ID, todo.Title, todo.Description, todo.Completed, todo.UserID, todo.CreatedAt, todo.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetTodos() ([]*types.Todo, error) {
	rows, err := s.db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}

	todos := make([]*types.Todo, 0)
	for rows.Next() {
		p, err := scanRowsIntoTodo(rows)
		if err != nil {
			return nil, err
		}

		todos = append(todos, p)
	}

	return todos, nil
}

func (s *Store) GetTodoByID(todoID string) (*types.Todo, error) {
	rows, err := s.db.Query("SELECT * FROM todos WHERE id = $1", todoID)
	if err != nil {
		return nil, err
	}

	p := new(types.Todo)
	for rows.Next() {
		p, err = scanRowsIntoTodo(rows)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func scanRowsIntoTodo(rows *sql.Rows) (*types.Todo, error) {
	product := new(types.Todo)

	err := rows.Scan(
		&product.ID,
		&product.Title,
		&product.Description,
		&product.Completed,
		&product.UserID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}
