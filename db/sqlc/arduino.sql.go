// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: arduino.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createArduino = `-- name: CreateArduino :one
INSERT INTO "arduino" ("name") VALUES ($1) RETURNING id, name
`

func (q *Queries) CreateArduino(ctx context.Context, name string) (Arduino, error) {
	row := q.db.QueryRow(ctx, createArduino, name)
	var i Arduino
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const deleteArduino = `-- name: DeleteArduino :one
DELETE FROM "arduino" WHERE "id" = $1 RETURNING id, name
`

func (q *Queries) DeleteArduino(ctx context.Context, id int32) (Arduino, error) {
	row := q.db.QueryRow(ctx, deleteArduino, id)
	var i Arduino
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const listArduinos = `-- name: ListArduinos :many
SELECT id, name FROM "arduino" WHERE "name" ILIKE '%' || $1 || '%' LIMIT $3 OFFSET $2
`

type ListArduinosParams struct {
	Name         pgtype.Text
	OffsetNumber int32
	LimitNumber  int32
}

func (q *Queries) ListArduinos(ctx context.Context, arg ListArduinosParams) ([]Arduino, error) {
	rows, err := q.db.Query(ctx, listArduinos, arg.Name, arg.OffsetNumber, arg.LimitNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Arduino{}
	for rows.Next() {
		var i Arduino
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateArduino = `-- name: UpdateArduino :one
UPDATE "arduino" SET "name" = $1 WHERE "id" = $2 RETURNING id, name
`

type UpdateArduinoParams struct {
	Name string
	ID   int32
}

func (q *Queries) UpdateArduino(ctx context.Context, arg UpdateArduinoParams) (Arduino, error) {
	row := q.db.QueryRow(ctx, updateArduino, arg.Name, arg.ID)
	var i Arduino
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
