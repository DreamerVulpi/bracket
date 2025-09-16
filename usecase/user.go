package usecase

import (
	"context"
	"fmt"

	"github.com/DreamerVulpi/bracket/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AddUser(conn *pgxpool.Pool, request entity.RequestUserAdd) (entity.ResponseUserAdd, error) {
	if request.Nickname == "" {
		return entity.ResponseUserAdd{}, fmt.Errorf("failed ADD user")
	}

	sql := "INSERT INTO \"users\" (Nickname) VALUES ($1) RETURNING id;"
	var userId int
	err := conn.QueryRow(context.Background(), sql, request.Nickname).Scan(&userId)
	if err != nil {
		fmt.Printf("Unable to add to database, %s", err)
		return entity.ResponseUserAdd{}, fmt.Errorf("failed ADD user (add db)")
	}

	return entity.ResponseUserAdd{Id: userId}, nil
}

func EditUser(conn *pgxpool.Pool, request entity.RequestUserEdit) error {
	if request.Player.Id == 0 {
		return fmt.Errorf("failed EDIT user: no ID")
	}

	sql := "UPDATE users SET Nickname = $1 WHERE Id = $2;"
	_, err := conn.Exec(context.Background(), sql, request.Player.Nickname, request.Player.Id)
	if err != nil {
		fmt.Printf("Unable to add to database, %s", err)
		return fmt.Errorf("failed ADD user (add db)")
	}

	return nil
}

func DeleteUser(conn *pgxpool.Pool, request entity.RequestUserDelete) error {
	if request.Id == 0 {
		return fmt.Errorf("failed DELETE user: no ID")
	}

	// TODO: CASCADE?
	sql := "DELETE FROM users WHERE Id = $1;"
	_, err := conn.Exec(context.Background(), sql, request.Id)
	if err != nil {
		fmt.Printf("Unable to delete to database, %s", err)
		return fmt.Errorf("failed DELETE user")
	}

	fmt.Println("correct delete")
	return nil
}

func GetUser(conn *pgxpool.Pool, request entity.RequestUserGet) (entity.ResponseUserGet, error) {
	if request.Id == 0 {
		return entity.ResponseUserGet{}, fmt.Errorf("failed GET user: no ID")
	}

	sql := "SELECT * FROM users WHERE Id = $1;"
	var user entity.User
	err := conn.QueryRow(context.Background(), sql, request.Id).Scan(&user.Id, &user.Nickname)
	if err != nil {
		fmt.Printf("Unable to get to database, %s", err)
		return entity.ResponseUserGet{}, fmt.Errorf("failed GET user")
	}

	return entity.ResponseUserGet{Player: user}, nil
}
