package models

import (
	"context"
	"database/sql"
	"github.com/ngohoang211020/blog/internal/common"
	"time"
)

type UserModel struct {
	DB *sql.DB
}

func (m UserModel) Insert(user *User) error {
	query := `INSERT INTO users (name, email, password_hash, activated)
	VALUES ($1,$2,$3,$4)
	RETURNING user_id,created_at,version`

	args := []interface{}{user.Name, user.Email, user.Password.Hash, user.Activated}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt, &user.Version)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return common.ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}
