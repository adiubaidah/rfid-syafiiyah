// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: holiday.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createHoliday = `-- name: CreateHoliday :one
INSERT INTO "holiday" ("name", "color", "description") VALUES ($1, $2, $3) RETURNING id, name, description, color
`

type CreateHolidayParams struct {
	Name        string
	Color       pgtype.Text
	Description pgtype.Text
}

func (q *Queries) CreateHoliday(ctx context.Context, arg CreateHolidayParams) (Holiday, error) {
	row := q.db.QueryRow(ctx, createHoliday, arg.Name, arg.Color, arg.Description)
	var i Holiday
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Color,
	)
	return i, err
}

const deleteHoliday = `-- name: DeleteHoliday :one
DELETE FROM "holiday" WHERE "id" = $1 RETURNING id, name, description, color
`

func (q *Queries) DeleteHoliday(ctx context.Context, id int32) (Holiday, error) {
	row := q.db.QueryRow(ctx, deleteHoliday, id)
	var i Holiday
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Color,
	)
	return i, err
}

const updateHoliday = `-- name: UpdateHoliday :one
UPDATE "holiday" SET "name" = $1, "color" = $2, "description" = $3 WHERE "id" = $4 RETURNING id, name, description, color
`

type UpdateHolidayParams struct {
	Name        string
	Color       pgtype.Text
	Description pgtype.Text
	ID          int32
}

func (q *Queries) UpdateHoliday(ctx context.Context, arg UpdateHolidayParams) (Holiday, error) {
	row := q.db.QueryRow(ctx, updateHoliday,
		arg.Name,
		arg.Color,
		arg.Description,
		arg.ID,
	)
	var i Holiday
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Color,
	)
	return i, err
}
