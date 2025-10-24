package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/DreamerVulpi/bracket/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Conn *pgxpool.Pool
}

func (u *User) Add(nickname string, password_hash string) (int, error) {
	const sql = "INSERT INTO users (nickname, password_hash) VALUES ($1, $2) RETURNING id"

	var id int

	err := u.Conn.QueryRow(context.Background(), sql, nickname, password_hash).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("unable to create player in database, %w", err)
	}

	return id, nil
}

func (u *User) Get(id int) (entity.User, error) {
	const sql = "SELECT u.id, u.nickname, u.password_hash FROM users u WHERE id = $1"

	var user entity.User

	var ignoredField any
	err := u.Conn.QueryRow(context.Background(), sql, id).Scan(&user.Id, &user.Nickname, &ignoredField)
	if err != nil {
		return entity.User{}, fmt.Errorf("unable to get from database, %w", err)
	}

	return user, nil
}

func (u *User) Edit(user entity.User) error {
	const sql = "UPDATE users SET nickname = $1 WHERE id = $2"

	tag, err := u.Conn.Exec(context.Background(), sql, user.Nickname, user.Id)
	log.Println(user)
	if err != nil {
		return fmt.Errorf("unable to edit user from database, %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("user doesn't exist")
	}

	return nil
}

func (u *User) Delete(id int) error {
	const sql = "DELETE FROM users WHERE id = $1"

	tag, err := u.Conn.Exec(context.Background(), sql, id)
	if err != nil {
		return fmt.Errorf("don't deleted user from database, %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("user doesn't exist")
	}

	return nil
}
