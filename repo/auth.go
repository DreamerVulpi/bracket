package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Auth struct {
	Conn *pgxpool.Pool
}

func (a *Auth) GetToken(id int) (string, error) {
	const sql = "SELECT u.token FROM users u WHERE u.id = $1"

	var token string

	err := a.Conn.QueryRow(context.Background(), sql, id).Scan(&token)
	if err != nil {
		return "", err
	}

	return token, nil
}
