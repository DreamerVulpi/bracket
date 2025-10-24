package repo

import (
	"fmt"

	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Auth struct {
	Conn *pgxpool.Pool
}

func (a *Auth) GetHash(nickname, password string) (string, error) {
	const sql = "SELECT u.password FROM users u WHERE nickname = $1"

	var responseHash string

	err := a.Conn.QueryRow(context.Background(), sql, nickname).Scan(&responseHash)
	if err != nil {
		return "", fmt.Errorf("unable to get from database using nickname, %w", err)
	}

	return responseHash, nil
}
