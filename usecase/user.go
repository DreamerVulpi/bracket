package usecase

import (
	"context"
	"fmt"

	"github.com/DreamerVulpi/bracket/entity"
	"github.com/jackc/pgx/v5"
)

func AddUser(request entity.RequestUserAdd) (entity.ResponseUserAdd, error) {
	if request.Nickname == "" {
		return entity.ResponseUserAdd{}, fmt.Errorf("failed ADD user")
	}

	// TODO: Нормально реализовать соединение
	conn, err := pgx.Connect(context.Background(), "postgresql://superuser:1234@localhost:5432/bracketProject")
	if err != nil {
		fmt.Printf("Unable to connect to database, %s", err)
		return entity.ResponseUserAdd{}, fmt.Errorf("failed ADD user (connect db)")
	}
	defer conn.Close(context.Background())

	sql := "INSERT INTO \"users\" (Nickname) VALUES ($1) RETURNING id;"
	var userId int
	err = conn.QueryRow(context.Background(), sql, request.Nickname).Scan(&userId)
	if err != nil {
		fmt.Printf("Unable to add to database, %s", err)
		return entity.ResponseUserAdd{}, fmt.Errorf("failed ADD user (add db)")
	}

	return entity.ResponseUserAdd{Id: userId}, nil
}

func EditUser(request entity.RequestUserEdit) error {
	if request.Player.Id == "" {
		return fmt.Errorf("failed EDIT user: no ID")
	}

	// TODO: SQL EDIT USER
	fmt.Println(request.Player)
	return nil
}

func DeleteUser(request entity.RequestUserDelete) error {
	if request.Id == "" {
		return fmt.Errorf("failed DELETE user: no ID")
	}

	// TODO: SQL DELETE USER
	fmt.Println("correct delete")
	return nil
}

func GetUser(request entity.RequestUserGet) (entity.User, error) {
	if request.Id == "" {
		return entity.User{}, fmt.Errorf("failed GET user: no ID")
	}

	// TODO: SQL GET USER
	result := entity.User{
		Id:       "12345",
		Nickname: "Player1",
	}
	return result, nil
}
