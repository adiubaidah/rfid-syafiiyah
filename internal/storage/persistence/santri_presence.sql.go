// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: santri_presence.sql

package persistence

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countSantriPresences = `-- name: CountSantriPresences :one
SELECT
    COUNT(*)
FROM
    "santri_presence"
    INNER JOIN "santri" ON "santri_presence"."santri_id" = "santri"."id"
WHERE
    (
        $1 :: integer IS NULL
        OR "santri_id" = $1 :: integer
    )
    AND (
        $2 :: text IS NULL
        OR "santri"."name" ILIKE '%' || $2 || '%'
    )
    AND (
        $3 :: presence_type IS NULL
        OR "type" = $3 :: presence_type
    )
    AND (
        $4 :: integer IS NULL
        OR "schedule_id" = $4 :: integer
    )
    AND (
        $5 :: date IS NULL
        OR DATE("created_at") >= $5 :: date
    )
    AND (
        $6 :: date IS NULL
        OR DATE("created_at") <= $6 :: date
    )
`

type CountSantriPresencesParams struct {
	SantriID   pgtype.Int4      `db:"santri_id"`
	Q          pgtype.Text      `db:"q"`
	Type       NullPresenceType `db:"type"`
	ScheduleID pgtype.Int4      `db:"schedule_id"`
	FromDate   pgtype.Date      `db:"from_date"`
	ToDate     pgtype.Date      `db:"to_date"`
}

func (q *Queries) CountSantriPresences(ctx context.Context, arg CountSantriPresencesParams) (int64, error) {
	row := q.db.QueryRow(ctx, countSantriPresences,
		arg.SantriID,
		arg.Q,
		arg.Type,
		arg.ScheduleID,
		arg.FromDate,
		arg.ToDate,
	)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createSantriPresence = `-- name: CreateSantriPresence :one
INSERT INTO
    "santri_presence" (
        "schedule_id",
        "schedule_name",
        "type",
        "santri_id",
        "notes",
        "created_by",
        "santri_permission_id"
    )
VALUES
    (
        $1,
        $2,
        $3 :: presence_type,
        $4,
        $5,
        $6 :: presence_created_by_type,
        $7
    ) RETURNING id, schedule_id, schedule_name, type, santri_id, created_at, created_by, notes, santri_permission_id, created_date
`

type CreateSantriPresenceParams struct {
	ScheduleID         int32                 `db:"schedule_id"`
	ScheduleName       string                `db:"schedule_name"`
	Type               PresenceType          `db:"type"`
	SantriID           int32                 `db:"santri_id"`
	Notes              pgtype.Text           `db:"notes"`
	CreatedBy          PresenceCreatedByType `db:"created_by"`
	SantriPermissionID pgtype.Int4           `db:"santri_permission_id"`
}

func (q *Queries) CreateSantriPresence(ctx context.Context, arg CreateSantriPresenceParams) (SantriPresence, error) {
	row := q.db.QueryRow(ctx, createSantriPresence,
		arg.ScheduleID,
		arg.ScheduleName,
		arg.Type,
		arg.SantriID,
		arg.Notes,
		arg.CreatedBy,
		arg.SantriPermissionID,
	)
	var i SantriPresence
	err := row.Scan(
		&i.ID,
		&i.ScheduleID,
		&i.ScheduleName,
		&i.Type,
		&i.SantriID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.Notes,
		&i.SantriPermissionID,
		&i.CreatedDate,
	)
	return i, err
}

type CreateSantriPresencesParams struct {
	ScheduleID         int32                 `db:"schedule_id"`
	ScheduleName       string                `db:"schedule_name"`
	Type               PresenceType          `db:"type"`
	SantriID           int32                 `db:"santri_id"`
	Notes              pgtype.Text           `db:"notes"`
	CreatedAt          pgtype.Timestamptz    `db:"created_at"`
	CreatedBy          PresenceCreatedByType `db:"created_by"`
	SantriPermissionID pgtype.Int4           `db:"santri_permission_id"`
}

const deleteSantriPresence = `-- name: DeleteSantriPresence :one
DELETE FROM
    "santri_presence"
WHERE
    "id" = $1
RETURNING id, schedule_id, schedule_name, type, santri_id, created_at, created_by, notes, santri_permission_id, created_date
`

func (q *Queries) DeleteSantriPresence(ctx context.Context, id int32) (SantriPresence, error) {
	row := q.db.QueryRow(ctx, deleteSantriPresence, id)
	var i SantriPresence
	err := row.Scan(
		&i.ID,
		&i.ScheduleID,
		&i.ScheduleName,
		&i.Type,
		&i.SantriID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.Notes,
		&i.SantriPermissionID,
		&i.CreatedDate,
	)
	return i, err
}

const listAbsentSantri = `-- name: ListAbsentSantri :many
SELECT 
    "santri"."id", "santri"."name"
FROM
    "santri"
WHERE
    NOT EXISTS (
        SELECT
            1
        FROM
            "santri_presence"
        WHERE
            "santri_presence"."santri_id" = "santri"."id"
            AND DATE("santri_presence"."created_at") = $1::date
            AND "santri_presence"."schedule_id" = $2::integer
    )
`

type ListAbsentSantriParams struct {
	Date       pgtype.Date `db:"date"`
	ScheduleID pgtype.Int4 `db:"schedule_id"`
}

type ListAbsentSantriRow struct {
	ID   int32  `db:"id"`
	Name string `db:"name"`
}

func (q *Queries) ListAbsentSantri(ctx context.Context, arg ListAbsentSantriParams) ([]ListAbsentSantriRow, error) {
	rows, err := q.db.Query(ctx, listAbsentSantri, arg.Date, arg.ScheduleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListAbsentSantriRow{}
	for rows.Next() {
		var i ListAbsentSantriRow
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

const listSantriPresences = `-- name: ListSantriPresences :many
SELECT
    santri_presence.id, santri_presence.schedule_id, santri_presence.schedule_name, santri_presence.type, santri_presence.santri_id, santri_presence.created_at, santri_presence.created_by, santri_presence.notes, santri_presence.santri_permission_id, santri_presence.created_date,
    "santri"."name" AS "santri_name"
FROM
    "santri_presence"
    INNER JOIN "santri" ON "santri_presence"."santri_id" = "santri"."id"
WHERE
    (
        $1 :: integer IS NULL
        OR "santri_id" = $1 :: integer
    )
    AND (
        $2 :: text IS NULL
        OR "santri"."name" ILIKE '%' || $2 || '%'
    )
    AND (
        $3 :: presence_type IS NULL
        OR "type" = $3 :: presence_type
    )
    AND (
        $4 :: integer IS NULL
        OR "schedule_id" = $4 :: integer
    )
    AND (
        $5 :: date IS NULL
        OR DATE("created_at") >= $5 :: date
    )
    AND (
        $6 :: date IS NULL
        OR DATE("created_at") <= $6 :: date
    )
ORDER BY
    "santri_presence"."id" DESC
LIMIT
    $8 OFFSET $7
`

type ListSantriPresencesParams struct {
	SantriID     pgtype.Int4      `db:"santri_id"`
	Q            pgtype.Text      `db:"q"`
	Type         NullPresenceType `db:"type"`
	ScheduleID   pgtype.Int4      `db:"schedule_id"`
	FromDate     pgtype.Date      `db:"from_date"`
	ToDate       pgtype.Date      `db:"to_date"`
	OffsetNumber int32            `db:"offset_number"`
	LimitNumber  int32            `db:"limit_number"`
}

type ListSantriPresencesRow struct {
	ID                 int32                 `db:"id"`
	ScheduleID         int32                 `db:"schedule_id"`
	ScheduleName       string                `db:"schedule_name"`
	Type               PresenceType          `db:"type"`
	SantriID           int32                 `db:"santri_id"`
	CreatedAt          pgtype.Timestamptz    `db:"created_at"`
	CreatedBy          PresenceCreatedByType `db:"created_by"`
	Notes              pgtype.Text           `db:"notes"`
	SantriPermissionID pgtype.Int4           `db:"santri_permission_id"`
	CreatedDate        pgtype.Date           `db:"created_date"`
	SantriName         string                `db:"santri_name"`
}

func (q *Queries) ListSantriPresences(ctx context.Context, arg ListSantriPresencesParams) ([]ListSantriPresencesRow, error) {
	rows, err := q.db.Query(ctx, listSantriPresences,
		arg.SantriID,
		arg.Q,
		arg.Type,
		arg.ScheduleID,
		arg.FromDate,
		arg.ToDate,
		arg.OffsetNumber,
		arg.LimitNumber,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListSantriPresencesRow{}
	for rows.Next() {
		var i ListSantriPresencesRow
		if err := rows.Scan(
			&i.ID,
			&i.ScheduleID,
			&i.ScheduleName,
			&i.Type,
			&i.SantriID,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.Notes,
			&i.SantriPermissionID,
			&i.CreatedDate,
			&i.SantriName,
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

const updateSantriPresence = `-- name: UpdateSantriPresence :one
UPDATE
    "santri_presence"
SET
    "schedule_id" = COALESCE($1, schedule_id),
    "schedule_name" = COALESCE($2, schedule_name),
    "type" = COALESCE($3::presence_type, type),
    "santri_id" = COALESCE($4, santri_id),
    "notes" = $5,
    "santri_permission_id" = $6
WHERE
    "id" = $7
RETURNING id, schedule_id, schedule_name, type, santri_id, created_at, created_by, notes, santri_permission_id, created_date
`

type UpdateSantriPresenceParams struct {
	ScheduleID         pgtype.Int4      `db:"schedule_id"`
	ScheduleName       pgtype.Text      `db:"schedule_name"`
	Type               NullPresenceType `db:"type"`
	SantriID           pgtype.Int4      `db:"santri_id"`
	Notes              pgtype.Text      `db:"notes"`
	SantriPermissionID pgtype.Int4      `db:"santri_permission_id"`
	ID                 int32            `db:"id"`
}

func (q *Queries) UpdateSantriPresence(ctx context.Context, arg UpdateSantriPresenceParams) (SantriPresence, error) {
	row := q.db.QueryRow(ctx, updateSantriPresence,
		arg.ScheduleID,
		arg.ScheduleName,
		arg.Type,
		arg.SantriID,
		arg.Notes,
		arg.SantriPermissionID,
		arg.ID,
	)
	var i SantriPresence
	err := row.Scan(
		&i.ID,
		&i.ScheduleID,
		&i.ScheduleName,
		&i.Type,
		&i.SantriID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.Notes,
		&i.SantriPermissionID,
		&i.CreatedDate,
	)
	return i, err
}
