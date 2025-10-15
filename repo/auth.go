package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Auth struct {
	Conn *pgxpool.Pool
}

func (a *Auth) CheckToken(token string) (bool, error) {
	const sql = "SELECT COUNT(*) FROM users u WHERE u.token = $1"

	var count int
	err := a.Conn.QueryRow(context.Background(), sql, token).Scan(&count)
	if err != nil {
		return false, err
	}

	if count != 1 {
		return false, fmt.Errorf("not finded user with tokenInput: %v", token)
	}

	return true, nil
}
