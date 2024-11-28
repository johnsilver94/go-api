package user

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/johnsilver94/go-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = $1", email)

	if err != nil {
		return nil, err
	}

	user := new(types.User)
	for rows.Next() {
		user, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == uuid.Nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *Store) GetUserByID(id string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = $1;", id)

	if err != nil {
		return nil, err
	}

	user := new(types.User)
	for rows.Next() {
		user, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == uuid.Nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (id, email, name, password ,created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6);", user.ID, user.Email, user.Name, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}
