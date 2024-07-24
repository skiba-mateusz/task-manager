package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/skiba-mateusz/task-manager/models"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateUser(user models.User) error {
	query := `
		INSERT INTO users (username, email, password)
		VALUES($1, $2, $3);
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	_, err := s.db.ExecContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.Password,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetUserByEmail(email string) (models.User, error) {
	query := `
		SELECT id, username, email, password, created_at FROM users WHERE email = $1;
	`
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	var user models.User
	err := s.db.QueryRowContext(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *Store) GetUserByID(id int) (models.User, error) {
	query := `
		SELECT id, username, email, password, created_at FROM users WHERE id = $1;
	`
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	var user models.User
	err := s.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}