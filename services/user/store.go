package user

import (
	"context"
	"database/sql"
	"time"
)

type UserStore interface {
	CreateUser(User) error
	GetUserByEmail(string) (User, error)
}

type Store struct {
	db *sql.DB
}

type User struct {
	ID 				int64		`json:"id"`
	Username 	string	`json:"username"`
	Email			string	`json:"email"`
	Password	string	`json:"-"`
	CreatedAt string 	`json:"created_at"` 
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateUser(user User) error {
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

func (s *Store) GetUserByEmail(email string) (User, error) {
	query := `
		SELECT id, username, email, password, created_at FROM users WHERE email = $1;
	`
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	var user User
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