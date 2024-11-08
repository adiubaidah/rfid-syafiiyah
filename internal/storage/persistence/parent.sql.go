// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: parent.sql

package persistence

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countParents = `-- name: CountParents :one
SELECT
    COUNT(*) AS "count"
FROM
    "parent"
WHERE
    (
        $1 :: text IS NULL
        OR "name" ILIKE '%' || $1 || '%'
    )
    AND (
        $2 :: smallint IS NULL
        OR (
            $2 = 1
            AND "user_id" IS NOT NULL
        )
        OR (
            $2 = 0
            AND "user_id" IS NULL
        )
        OR ($2 = -1)
    )
`

type CountParentsParams struct {
	Q       pgtype.Text `db:"q" json:"q"`
	HasUser int16       `db:"has_user" json:"has_user"`
}

func (q *Queries) CountParents(ctx context.Context, arg CountParentsParams) (int64, error) {
	row := q.db.QueryRow(ctx, countParents, arg.Q, arg.HasUser)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createParent = `-- name: CreateParent :one
INSERT INTO
    "parent" (
        "name",
        "address",
        "gender",
        "whatsapp_number",
        "photo",
        "user_id"
    )
VALUES
    (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6
    ) RETURNING id, name, address, gender, whatsapp_number, photo, user_id
`

type CreateParentParams struct {
	Name           string      `db:"name" json:"name"`
	Address        string      `db:"address" json:"address"`
	Gender         Gender      `db:"gender" json:"gender"`
	WhatsappNumber pgtype.Text `db:"whatsapp_number" json:"whatsapp_number"`
	Photo          pgtype.Text `db:"photo" json:"photo"`
	UserID         pgtype.Int4 `db:"user_id" json:"user_id"`
}

func (q *Queries) CreateParent(ctx context.Context, arg CreateParentParams) (Parent, error) {
	row := q.db.QueryRow(ctx, createParent,
		arg.Name,
		arg.Address,
		arg.Gender,
		arg.WhatsappNumber,
		arg.Photo,
		arg.UserID,
	)
	var i Parent
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.Gender,
		&i.WhatsappNumber,
		&i.Photo,
		&i.UserID,
	)
	return i, err
}

const deleteParent = `-- name: DeleteParent :one
DELETE FROM
    "parent"
WHERE
    "id" = $1 RETURNING id, name, address, gender, whatsapp_number, photo, user_id
`

func (q *Queries) DeleteParent(ctx context.Context, id int32) (Parent, error) {
	row := q.db.QueryRow(ctx, deleteParent, id)
	var i Parent
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.Gender,
		&i.WhatsappNumber,
		&i.Photo,
		&i.UserID,
	)
	return i, err
}

const getParent = `-- name: GetParent :one
SELECT
    parent.id, parent.name, parent.address, parent.gender, parent.whatsapp_number, parent.photo, parent.user_id,
    "user"."id" AS "user_id",
    "user"."username" AS "user_username"
FROM
    "parent"
    LEFT JOIN "user" ON "parent"."user_id" = "user"."id"
WHERE
    "parent"."id" = $1
`

type GetParentRow struct {
	ID             int32       `db:"id" json:"id"`
	Name           string      `db:"name" json:"name"`
	Address        string      `db:"address" json:"address"`
	Gender         Gender      `db:"gender" json:"gender"`
	WhatsappNumber pgtype.Text `db:"whatsapp_number" json:"whatsapp_number"`
	Photo          pgtype.Text `db:"photo" json:"photo"`
	UserID         pgtype.Int4 `db:"user_id" json:"user_id"`
	UserID_2       pgtype.Int4 `db:"user_id_2" json:"user_id_2"`
	UserUsername   pgtype.Text `db:"user_username" json:"user_username"`
}

func (q *Queries) GetParent(ctx context.Context, id int32) (GetParentRow, error) {
	row := q.db.QueryRow(ctx, getParent, id)
	var i GetParentRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.Gender,
		&i.WhatsappNumber,
		&i.Photo,
		&i.UserID,
		&i.UserID_2,
		&i.UserUsername,
	)
	return i, err
}

const listParentsAsc = `-- name: ListParentsAsc :many
SELECT
    parent.id, parent.name, parent.address, parent.gender, parent.whatsapp_number, parent.photo, parent.user_id,
    "user"."id" AS "user_id",
    "user"."username" AS "user_username"
FROM
    "parent"
    LEFT JOIN "user" ON "parent"."user_id" = "user"."id"
WHERE
    (
        $1 :: text IS NULL
        OR "name" ILIKE '%' || $1 || '%'
    )
    AND (
        $2 :: smallint IS NULL
        OR (
            $2 = 1
            AND "user_id" IS NOT NULL
        )
        OR (
            $2 = 0
            AND "user_id" IS NULL
        )
        OR ($2 = -1)
    )
ORDER BY
    "name" ASC
LIMIT
    $4 OFFSET $3
`

type ListParentsAscParams struct {
	Q            pgtype.Text `db:"q" json:"q"`
	HasUser      int16       `db:"has_user" json:"has_user"`
	OffsetNumber int32       `db:"offset_number" json:"offset_number"`
	LimitNumber  int32       `db:"limit_number" json:"limit_number"`
}

type ListParentsAscRow struct {
	ID             int32       `db:"id" json:"id"`
	Name           string      `db:"name" json:"name"`
	Address        string      `db:"address" json:"address"`
	Gender         Gender      `db:"gender" json:"gender"`
	WhatsappNumber pgtype.Text `db:"whatsapp_number" json:"whatsapp_number"`
	Photo          pgtype.Text `db:"photo" json:"photo"`
	UserID         pgtype.Int4 `db:"user_id" json:"user_id"`
	UserID_2       pgtype.Int4 `db:"user_id_2" json:"user_id_2"`
	UserUsername   pgtype.Text `db:"user_username" json:"user_username"`
}

func (q *Queries) ListParentsAsc(ctx context.Context, arg ListParentsAscParams) ([]ListParentsAscRow, error) {
	rows, err := q.db.Query(ctx, listParentsAsc,
		arg.Q,
		arg.HasUser,
		arg.OffsetNumber,
		arg.LimitNumber,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListParentsAscRow{}
	for rows.Next() {
		var i ListParentsAscRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Address,
			&i.Gender,
			&i.WhatsappNumber,
			&i.Photo,
			&i.UserID,
			&i.UserID_2,
			&i.UserUsername,
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

const listParentsDesc = `-- name: ListParentsDesc :many
SELECT
    parent.id, parent.name, parent.address, parent.gender, parent.whatsapp_number, parent.photo, parent.user_id,
    "user"."id" AS "user_id",
    "user"."username" AS "user_username"
FROM
    "parent"
    LEFT JOIN "user" ON "parent"."user_id" = "user"."id"
WHERE
    (
        $1 :: text IS NULL
        OR "name" ILIKE '%' || $1 || '%'
    )
    AND (
        $2 :: smallint IS NULL
        OR (
            $2 = 1
            AND "user_id" IS NOT NULL
        )
        OR (
            $2 = 0
            AND "user_id" IS NULL
        )
        OR ($2 = -1)
    )
ORDER BY
    "name" ASC
LIMIT
    $4 OFFSET $3
`

type ListParentsDescParams struct {
	Q            pgtype.Text `db:"q" json:"q"`
	HasUser      int16       `db:"has_user" json:"has_user"`
	OffsetNumber int32       `db:"offset_number" json:"offset_number"`
	LimitNumber  int32       `db:"limit_number" json:"limit_number"`
}

type ListParentsDescRow struct {
	ID             int32       `db:"id" json:"id"`
	Name           string      `db:"name" json:"name"`
	Address        string      `db:"address" json:"address"`
	Gender         Gender      `db:"gender" json:"gender"`
	WhatsappNumber pgtype.Text `db:"whatsapp_number" json:"whatsapp_number"`
	Photo          pgtype.Text `db:"photo" json:"photo"`
	UserID         pgtype.Int4 `db:"user_id" json:"user_id"`
	UserID_2       pgtype.Int4 `db:"user_id_2" json:"user_id_2"`
	UserUsername   pgtype.Text `db:"user_username" json:"user_username"`
}

func (q *Queries) ListParentsDesc(ctx context.Context, arg ListParentsDescParams) ([]ListParentsDescRow, error) {
	rows, err := q.db.Query(ctx, listParentsDesc,
		arg.Q,
		arg.HasUser,
		arg.OffsetNumber,
		arg.LimitNumber,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListParentsDescRow{}
	for rows.Next() {
		var i ListParentsDescRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Address,
			&i.Gender,
			&i.WhatsappNumber,
			&i.Photo,
			&i.UserID,
			&i.UserID_2,
			&i.UserUsername,
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

const updateParent = `-- name: UpdateParent :one
UPDATE
    "parent"
SET
    "name" = $1,
    "address" = $2,
    "gender" = $3,
    "whatsapp_number" = $4,
    "photo" = $5,
    "user_id" = $6
WHERE
    "id" = $7 RETURNING id, name, address, gender, whatsapp_number, photo, user_id
`

type UpdateParentParams struct {
	Name           string      `db:"name" json:"name"`
	Address        string      `db:"address" json:"address"`
	Gender         Gender      `db:"gender" json:"gender"`
	WhatsappNumber pgtype.Text `db:"whatsapp_number" json:"whatsapp_number"`
	Photo          pgtype.Text `db:"photo" json:"photo"`
	UserID         pgtype.Int4 `db:"user_id" json:"user_id"`
	ID             int32       `db:"id" json:"id"`
}

func (q *Queries) UpdateParent(ctx context.Context, arg UpdateParentParams) (Parent, error) {
	row := q.db.QueryRow(ctx, updateParent,
		arg.Name,
		arg.Address,
		arg.Gender,
		arg.WhatsappNumber,
		arg.Photo,
		arg.UserID,
		arg.ID,
	)
	var i Parent
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.Gender,
		&i.WhatsappNumber,
		&i.Photo,
		&i.UserID,
	)
	return i, err
}
