package usecase

import (
	"fmt"

	"github.com/DreamerVulpi/bracket/entity"
)

func AddUser(request entity.RequestUserAdd) (string, error) {
	if request.Nickname == "" {
		return "", fmt.Errorf("failed ADD user")
	}

	id := "11111"
	// TODO: SQL ADD USER
	return id, nil
}

func EditUser(request entity.RequestUserEdit) error {
	if request.Player.ID == "" {
		return fmt.Errorf("failed EDIT user: no ID")
	}

	// TODO: SQL EDIT USER
	fmt.Println(request.Player)
	return nil
}

func DeleteUser(request entity.RequestUserDelete) error {
	if request.ID == "" {
		return fmt.Errorf("failed DELETE user: no ID")
	}

	// TODO: SQL DELETE USER
	fmt.Println("correct delete")
	return nil
}

func GetUser(request entity.RequestUserGet) (entity.User, error) {
	if request.ID == "" {
		return entity.User{}, fmt.Errorf("failed GET user: no ID")
	}

	// TODO: SQL GET USER
	result := entity.User{
		ID:       "12345",
		Nickname: "Player1",
	}
	return result, nil
}
