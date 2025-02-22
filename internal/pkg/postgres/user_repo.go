package postgres

import (
	"database/sql"
	"errors"

	"github.com/VitaliyGopher/messanger/internal/pkg/model"
)

type UserRepo struct {
	store *Storage
}

func NewUserRepo(s *Storage) *UserRepo {
	return &UserRepo{
		store: s,
	}
}

func (r *UserRepo) Create(u *model.User) error {
	return r.store.DB.QueryRow(
		"INSERT INTO users (email, username) VALUES ($1, $2) RETURNING user_id;",
		u.Email,
		u.Username,
	).Scan(&u.ID)
}

func (r *UserRepo) FindByID(id uint) (*model.User, error) {
	u := &model.User{}

	if err := r.store.DB.QueryRow(
		"SELECT id, username, email FROM users WHERE id = $1",
		id,
	).Scan(&u.ID,
		&u.Username,
		&u.Email,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("rows not found")
		}

		return nil, err
	}

	return u, nil
}

func (r *UserRepo) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	if err := r.store.DB.QueryRow(
		"SELECT user_id, username, email FROM users WHERE email = $1",
		email,
	).Scan(&u.ID,
		&u.Username,
		&u.Email,
	); err != nil {
		return nil, err
	}

	return u, nil
}
