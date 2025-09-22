package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/DreamerVulpi/bracket/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Set struct {
	Conn *pgxpool.Pool
}

func (s *Set) Add(player1Id, player2Id, poolId int) (int, error) {
	const sql = "INSERT INTO sets (player1_id, player2_id, pool_id) VALUES ($1, $2, $3) RETURNING id"

	var id int

	err := s.Conn.QueryRow(context.Background(), sql, player1Id, player2Id, poolId).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("unable to create set in database, %w", err)
	}

	return id, nil
}

func (s *Set) Get(id int) (entity.Set, error) {
	const sql = "SELECT s.id, s.player1_id, s.player2_id, s.pool_id FROM sets s WHERE id = $1"

	var set entity.Set

	err := s.Conn.QueryRow(context.Background(), sql, id).Scan(&set.Id, &set.Player1Id, &set.Player2Id, &set.PoolId)
	if err != nil {
		return entity.Set{}, fmt.Errorf("unable to get from database, %w", err)
	}

	return set, nil
}

// TODO: Обновления 1 поля или 2 полей
func (s *Set) Edit(set entity.Set) error {
	const sql = "UPDATE sets SET player1_id = $1, player2_id = $2, pool_id = $3 WHERE id = $4"

	tag, err := s.Conn.Exec(context.Background(), sql, set.Player1Id, set.Player2Id, set.PoolId, set.Id)
	log.Println(set)
	if err != nil {
		return fmt.Errorf("unable to edit set from database, %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("set doesn't exist")
	}

	return nil
}

func (s *Set) Delete(id int) error {
	const sql = "DELETE FROM sets WHERE id = $1"

	tag, err := s.Conn.Exec(context.Background(), sql, id)
	if err != nil {
		return fmt.Errorf("don't deleted set from database, %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("set doesn't exist")
	}

	return nil
}
