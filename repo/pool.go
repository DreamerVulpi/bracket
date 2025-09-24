package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/DreamerVulpi/bracket/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool struct {
	Conn *pgxpool.Pool
}

func (p *Pool) Add(bracketId int) (int, error) {
	const sql = "INSERT INTO pools (bracket_id) VALUES ($1) RETURNING id"

	var id int

	err := p.Conn.QueryRow(context.Background(), sql, bracketId).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("unable to create pool in database, %w", err)
	}

	return id, nil
}

func (p *Pool) Get(id int) (entity.Pool, error) {
	const sql = "SELECT p.id, p.bracket_id FROM pools p WHERE id = $1"

	var pool entity.Pool

	err := p.Conn.QueryRow(context.Background(), sql, id).Scan(&pool.Id, &pool.BracketId)
	if err != nil {
		return entity.Pool{}, fmt.Errorf("unable to get from database, %w", err)
	}

	return pool, nil
}

func (p *Pool) Edit(pool entity.Pool) error {
	const sql = "UPDATE pools SET bracket_id = $1 WHERE id = $2"

	tag, err := p.Conn.Exec(context.Background(), sql, pool.BracketId, pool.Id)
	log.Println(pool)
	if err != nil {
		return fmt.Errorf("unable to edit pool from database, %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("pool doesn't exist")
	}

	return nil
}

func (p *Pool) Delete(id int) error {
	const sql = "DELETE FROM pools WHERE id = $1"

	tag, err := p.Conn.Exec(context.Background(), sql, id)
	if err != nil {
		return fmt.Errorf("don't deleted pool from database, %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("pool doesn't exist")
	}

	return nil
}
