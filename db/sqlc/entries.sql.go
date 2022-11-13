// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: entries.sql

package db

import (
	"context"
)

const createEntry = `-- name: CreateEntry :one
INSERT INTO entries (account_id,amount) VALUES ($1,$2) RETURNING id, account_id, amount, created_at
`

type CreateEntryParams struct {
	AccountID int64 `json:"accountID"`
	Amount    int64 `json:"amount"`
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, createEntry, arg.AccountID, arg.Amount)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getEntriesByAccountId = `-- name: GetEntriesByAccountId :many
SELECT id, account_id, amount, created_at FROM entries
WHERE account_id = $1
`

func (q *Queries) GetEntriesByAccountId(ctx context.Context, accountID int64) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, getEntriesByAccountId, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Entry{}
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEntriesByAmount = `-- name: GetEntriesByAmount :many
SELECT id, account_id, amount, created_at FROM entries
WHERE amount = $1
`

func (q *Queries) GetEntriesByAmount(ctx context.Context, amount int64) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, getEntriesByAmount, amount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Entry{}
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEntryById = `-- name: GetEntryById :one
SELECT id, account_id, amount, created_at FROM entries
WHERE id = $1
`

func (q *Queries) GetEntryById(ctx context.Context, id int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntryById, id)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
