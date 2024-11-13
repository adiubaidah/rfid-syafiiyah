// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: santri_permission.sql

package persistence

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSantriPermission = `-- name: CreateSantriPermission :one
INSERT INTO
    "santri_permission" (
        santri_id,
        start_permission,
        end_permission,
        "type",
        excuse
    )
VALUES
    (
        $1,
        $2,
        $3,
        $4 :: santri_permission_type,
        $5
    ) RETURNING id, santri_id, type, start_permission, end_permission, excuse
`

type CreateSantriPermissionParams struct {
	SantriID        int32                `db:"santri_id"`
	StartPermission pgtype.Timestamptz   `db:"start_permission"`
	EndPermission   pgtype.Timestamptz   `db:"end_permission"`
	Type            SantriPermissionType `db:"type"`
	Excuse          string               `db:"excuse"`
}

func (q *Queries) CreateSantriPermission(ctx context.Context, arg CreateSantriPermissionParams) (SantriPermission, error) {
	row := q.db.QueryRow(ctx, createSantriPermission,
		arg.SantriID,
		arg.StartPermission,
		arg.EndPermission,
		arg.Type,
		arg.Excuse,
	)
	var i SantriPermission
	err := row.Scan(
		&i.ID,
		&i.SantriID,
		&i.Type,
		&i.StartPermission,
		&i.EndPermission,
		&i.Excuse,
	)
	return i, err
}

const deleteSantriPermission = `-- name: DeleteSantriPermission :one
DELETE FROM
    "santri_permission"
WHERE
    "id" = $1 RETURNING id, santri_id, type, start_permission, end_permission, excuse
`

func (q *Queries) DeleteSantriPermission(ctx context.Context, id int32) (SantriPermission, error) {
	row := q.db.QueryRow(ctx, deleteSantriPermission, id)
	var i SantriPermission
	err := row.Scan(
		&i.ID,
		&i.SantriID,
		&i.Type,
		&i.StartPermission,
		&i.EndPermission,
		&i.Excuse,
	)
	return i, err
}

const getSantriPermission = `-- name: GetSantriPermission :one
SELECT
    santri_permission.id, santri_permission.santri_id, santri_permission.type, santri_permission.start_permission, santri_permission.end_permission, santri_permission.excuse,
    "santri"."name" AS "santri_name"
FROM
    "santri_permission"
    INNER JOIN "santri" ON "santri_permission"."santri_id" = "santri"."id"
WHERE
    "santri_permission"."id" = $1
`

type GetSantriPermissionRow struct {
	ID              int32                `db:"id"`
	SantriID        int32                `db:"santri_id"`
	Type            SantriPermissionType `db:"type"`
	StartPermission pgtype.Timestamptz   `db:"start_permission"`
	EndPermission   pgtype.Timestamptz   `db:"end_permission"`
	Excuse          string               `db:"excuse"`
	SantriName      string               `db:"santri_name"`
}

func (q *Queries) GetSantriPermission(ctx context.Context, id int32) (GetSantriPermissionRow, error) {
	row := q.db.QueryRow(ctx, getSantriPermission, id)
	var i GetSantriPermissionRow
	err := row.Scan(
		&i.ID,
		&i.SantriID,
		&i.Type,
		&i.StartPermission,
		&i.EndPermission,
		&i.Excuse,
		&i.SantriName,
	)
	return i, err
}

const listSantriPermissions = `-- name: ListSantriPermissions :many
SELECT
    santri_permission.id, santri_permission.santri_id, santri_permission.type, santri_permission.start_permission, santri_permission.end_permission, santri_permission.excuse,
    "santri"."name" AS "santri_name"
FROM
    "santri_permission"
    INNER JOIN "santri" ON "santri_permission"."santri_id" = "santri"."id"
WHERE
    ($1 :: text IS NULL
    OR "santri"."name" ILIKE '%' || $1 || '%')
    AND (
        $2 :: integer IS NULL
        OR "santri_id" = $2 :: integer
    )
    AND (
        $3 :: santri_permission_type IS NULL
        OR "type" = $3 :: santri_permission_type
    )
    AND (
        $4 :: timestamptz IS NULL
        OR "start_permission" >= $4 :: timestamptz
    )
    AND (
        $5 :: timestamptz IS NULL
        OR "end_permission" <= $5 :: timestamptz
    )
LIMIT
    $7 OFFSET $6
`

type ListSantriPermissionsParams struct {
	Q            pgtype.Text              `db:"q"`
	SantriID     pgtype.Int4              `db:"santri_id"`
	Type         NullSantriPermissionType `db:"type"`
	FromDate     pgtype.Timestamptz       `db:"from_date"`
	EndDate      pgtype.Timestamptz       `db:"end_date"`
	OffsetNumber int32                    `db:"offset_number"`
	LimitNumber  int32                    `db:"limit_number"`
}

type ListSantriPermissionsRow struct {
	ID              int32                `db:"id"`
	SantriID        int32                `db:"santri_id"`
	Type            SantriPermissionType `db:"type"`
	StartPermission pgtype.Timestamptz   `db:"start_permission"`
	EndPermission   pgtype.Timestamptz   `db:"end_permission"`
	Excuse          string               `db:"excuse"`
	SantriName      string               `db:"santri_name"`
}

func (q *Queries) ListSantriPermissions(ctx context.Context, arg ListSantriPermissionsParams) ([]ListSantriPermissionsRow, error) {
	rows, err := q.db.Query(ctx, listSantriPermissions,
		arg.Q,
		arg.SantriID,
		arg.Type,
		arg.FromDate,
		arg.EndDate,
		arg.OffsetNumber,
		arg.LimitNumber,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListSantriPermissionsRow{}
	for rows.Next() {
		var i ListSantriPermissionsRow
		if err := rows.Scan(
			&i.ID,
			&i.SantriID,
			&i.Type,
			&i.StartPermission,
			&i.EndPermission,
			&i.Excuse,
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

const updateSantriPermission = `-- name: UpdateSantriPermission :one
UPDATE
    "santri_permission"
SET
    "santri_id" = COALESCE($1, santri_id),
    "start_permission" = COALESCE($2, start_permission),
    "end_permission" = $3,
    "excuse" = COALESCE($4, excuse)
WHERE
    "id" = $5 RETURNING id, santri_id, type, start_permission, end_permission, excuse
`

type UpdateSantriPermissionParams struct {
	SantriID        pgtype.Int4        `db:"santri_id"`
	StartPermission pgtype.Timestamptz `db:"start_permission"`
	EndPermission   pgtype.Timestamptz `db:"end_permission"`
	Excuse          pgtype.Text        `db:"excuse"`
	ID              int32              `db:"id"`
}

func (q *Queries) UpdateSantriPermission(ctx context.Context, arg UpdateSantriPermissionParams) (SantriPermission, error) {
	row := q.db.QueryRow(ctx, updateSantriPermission,
		arg.SantriID,
		arg.StartPermission,
		arg.EndPermission,
		arg.Excuse,
		arg.ID,
	)
	var i SantriPermission
	err := row.Scan(
		&i.ID,
		&i.SantriID,
		&i.Type,
		&i.StartPermission,
		&i.EndPermission,
		&i.Excuse,
	)
	return i, err
}
