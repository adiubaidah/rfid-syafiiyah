// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: santri_schedule.sql

package persistence

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSantriSchedule = `-- name: CreateSantriSchedule :one
INSERT INTO
    "santri_schedule" (
        "name",
        "description",
        "start_presence",
        "start_time",
        "finish_time"
    )
VALUES
    (
        $1,
        $2,
        $3 :: time,
        $4 :: time,
        $5 :: time
    ) RETURNING id, name, description, start_presence, start_time, finish_time
`

type CreateSantriScheduleParams struct {
	Name          string      `db:"name"`
	Description   pgtype.Text `db:"description"`
	StartPresence pgtype.Time `db:"start_presence"`
	StartTime     pgtype.Time `db:"start_time"`
	FinishTime    pgtype.Time `db:"finish_time"`
}

func (q *Queries) CreateSantriSchedule(ctx context.Context, arg CreateSantriScheduleParams) (SantriSchedule, error) {
	row := q.db.QueryRow(ctx, createSantriSchedule,
		arg.Name,
		arg.Description,
		arg.StartPresence,
		arg.StartTime,
		arg.FinishTime,
	)
	var i SantriSchedule
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.StartPresence,
		&i.StartTime,
		&i.FinishTime,
	)
	return i, err
}

const deleteSantriSchedule = `-- name: DeleteSantriSchedule :one
DELETE FROM
    "santri_schedule"
WHERE
    "id" = $1 RETURNING id, name, description, start_presence, start_time, finish_time
`

func (q *Queries) DeleteSantriSchedule(ctx context.Context, id int32) (SantriSchedule, error) {
	row := q.db.QueryRow(ctx, deleteSantriSchedule, id)
	var i SantriSchedule
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.StartPresence,
		&i.StartTime,
		&i.FinishTime,
	)
	return i, err
}

const getLastSantriSchedule = `-- name: GetLastSantriSchedule :one
SELECT
    id, name, description, start_presence, start_time, finish_time
FROM
    "santri_schedule"
WHERE
    start_time = (
        SELECT
            MAX(start_time)
        FROM
            "santri_schedule"
    )
`

func (q *Queries) GetLastSantriSchedule(ctx context.Context) (SantriSchedule, error) {
	row := q.db.QueryRow(ctx, getLastSantriSchedule)
	var i SantriSchedule
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.StartPresence,
		&i.StartTime,
		&i.FinishTime,
	)
	return i, err
}

const getSantriSchedule = `-- name: GetSantriSchedule :one

SELECT
    id, name, description, start_presence, start_time, finish_time
FROM
    "santri_schedule"
WHERE
    $1::time BETWEEN start_presence AND finish_time
LIMIT
    1
`

func (q *Queries) GetSantriSchedule(ctx context.Context, time pgtype.Time) (SantriSchedule, error) {
	row := q.db.QueryRow(ctx, getSantriSchedule, time)
	var i SantriSchedule
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.StartPresence,
		&i.StartTime,
		&i.FinishTime,
	)
	return i, err
}

const listSantriSchedules = `-- name: ListSantriSchedules :many
SELECT
    id, name, description, start_presence, start_time, finish_time
FROM
    "santri_schedule"
WHERE
(   
    $1::time IS NULL OR 
    $1::time BETWEEN start_presence AND finish_time
)   
ORDER BY
    "start_time"
`

func (q *Queries) ListSantriSchedules(ctx context.Context, time pgtype.Time) ([]SantriSchedule, error) {
	rows, err := q.db.Query(ctx, listSantriSchedules, time)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SantriSchedule{}
	for rows.Next() {
		var i SantriSchedule
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.StartPresence,
			&i.StartTime,
			&i.FinishTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSantriSchedule = `-- name: UpdateSantriSchedule :one
UPDATE
    "santri_schedule"
SET
    "name" = COALESCE($1, name),
    "description" = $2,
    "start_presence" = COALESCE(
        $3 :: time,
        start_presence
    ),
    "start_time" = COALESCE($4 :: time, start_time),
    "finish_time" = COALESCE($5 :: time, finish_time)
WHERE
    "id" = $6 RETURNING id, name, description, start_presence, start_time, finish_time
`

type UpdateSantriScheduleParams struct {
	Name          pgtype.Text `db:"name"`
	Description   pgtype.Text `db:"description"`
	StartPresence pgtype.Time `db:"start_presence"`
	StartTime     pgtype.Time `db:"start_time"`
	FinishTime    pgtype.Time `db:"finish_time"`
	ID            int32       `db:"id"`
}

func (q *Queries) UpdateSantriSchedule(ctx context.Context, arg UpdateSantriScheduleParams) (SantriSchedule, error) {
	row := q.db.QueryRow(ctx, updateSantriSchedule,
		arg.Name,
		arg.Description,
		arg.StartPresence,
		arg.StartTime,
		arg.FinishTime,
		arg.ID,
	)
	var i SantriSchedule
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.StartPresence,
		&i.StartTime,
		&i.FinishTime,
	)
	return i, err
}
