// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countUser = `-- name: CountUser :one
SELECT
    COUNT(*) AS "count"
FROM
    "user"
WHERE
    (
        $1::text IS NULL
        OR "username" ILIKE '%' || $1 || '%'
    )
    AND (
        $2::text IS NULL
        OR "role" = $2
    )
    AND (
        $3::smallint IS NULL
        OR (
            $3 = 1
            AND "parent_id" IS NOT NULL OR "employee_id" IS NOT NULL
        )
        OR (
            $3 = 0
            AND "parent_id" IS NULL AND "employee_id" IS NULL
        )
        OR ($3 = -1)
    )
`

type CountUserParams struct {
	Q           pgtype.Text
	Role        pgtype.Text
	HasRelation int16
}

func (q *Queries) CountUser(ctx context.Context, arg CountUserParams) (int64, error) {
	row := q.db.QueryRow(ctx, countUser, arg.Q, arg.Role, arg.HasRelation)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO
    "user" ("role", "username", "password")
VALUES
    ($1::user_role, $2::text, $3::text) RETURNING id, role, username, password
`

type CreateUserParams struct {
	Role     UserRole
	Username string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Role, arg.Username, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Role,
		&i.Username,
		&i.Password,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :one
DELETE FROM
    "user"
WHERE
    "id" = $1
RETURNING id, role, username, password
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, deleteUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Role,
		&i.Username,
		&i.Password,
	)
	return i, err
}

const queryUserAscUsername = `-- name: QueryUserAscUsername :many
SELECT
    "user"."id",
    "user"."username",
    "user"."role",
    COALESCE("parent"."id", 0) AS "parentID",
    "parent"."name" AS "parentName",
    COALESCE("employee"."id", 0) AS "employeeID",
    "employee"."name" AS "employeeName"
FROM
    "user"
    LEFT JOIN "parent" ON "user"."id" = "parent"."user_id"
    LEFT JOIN "employee" ON "user"."id" = "employee"."user_id"
WHERE
    (
        $1::text IS NULL
        OR "user"."username" ILIKE '%' || $1 || '%'
    )
    AND (
        $2::user_role IS NULL
        OR "user"."role" = $2
    )
    AND (
        $3::smallint IS NULL
        OR (
            $3 = 1
            AND "parent"."id" IS NOT NULL OR "employee"."id" IS NOT NULL
        )
        OR (
            $3 = 0
            AND "parent"."id" IS NULL AND "employee"."id" IS NULL
        )
        OR ($3 = -1)
    )
ORDER BY
    "user"."username" ASC
LIMIT
    $5 OFFSET $4
`

type QueryUserAscUsernameParams struct {
	Q            pgtype.Text
	Role         NullUserRole
	HasRelation  int16
	OffsetNumber int32
	LimitNumber  int32
}

type QueryUserAscUsernameRow struct {
	ID           int32
	Username     pgtype.Text
	Role         NullUserRole
	ParentID     int32
	ParentName   pgtype.Text
	EmployeeID   int32
	EmployeeName pgtype.Text
}

func (q *Queries) QueryUserAscUsername(ctx context.Context, arg QueryUserAscUsernameParams) ([]QueryUserAscUsernameRow, error) {
	rows, err := q.db.Query(ctx, queryUserAscUsername,
		arg.Q,
		arg.Role,
		arg.HasRelation,
		arg.OffsetNumber,
		arg.LimitNumber,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []QueryUserAscUsernameRow{}
	for rows.Next() {
		var i QueryUserAscUsernameRow
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Role,
			&i.ParentID,
			&i.ParentName,
			&i.EmployeeID,
			&i.EmployeeName,
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

const queryUserDescUsername = `-- name: QueryUserDescUsername :many
SELECT
    "user"."id",
    "user"."username",
    "user"."role",
    "parent"."id" AS "parentID",
    "parent"."name" AS "parentName",
    "employee"."id" AS "employeeID",
    "employee"."name" AS "employeeName"
FROM
    "user"
    LEFT JOIN "parent" ON "user"."id" = "parent"."user_id"
    LEFT JOIN "employee" ON "user"."id" = "employee"."user_id"
WHERE
    (
        $1::text IS NULL
        OR "user"."username" ILIKE '%' || $1 || '%'
    )
    AND (
        $2::text IS NULL
        OR "user"."role" = $2
    )
    AND (
        $3::smallint IS NULL
        OR (
            $3 = 1
            AND "parent"."id" IS NOT NULL OR "employee"."id" IS NOT NULL
        )
        OR (
            $3 = 0
            AND "parent"."id" IS NULL AND "employee"."id" IS NULL
        )
        OR ($3 = -1)
    )
ORDER BY
    "user"."username" DESC
LIMIT
    $5 OFFSET $4
`

type QueryUserDescUsernameParams struct {
	Q            pgtype.Text
	Role         pgtype.Text
	HasRelation  int16
	OffsetNumber int32
	LimitNumber  int32
}

type QueryUserDescUsernameRow struct {
	ID           int32
	Username     pgtype.Text
	Role         NullUserRole
	ParentID     pgtype.Int4
	ParentName   pgtype.Text
	EmployeeID   pgtype.Int4
	EmployeeName pgtype.Text
}

func (q *Queries) QueryUserDescUsername(ctx context.Context, arg QueryUserDescUsernameParams) ([]QueryUserDescUsernameRow, error) {
	rows, err := q.db.Query(ctx, queryUserDescUsername,
		arg.Q,
		arg.Role,
		arg.HasRelation,
		arg.OffsetNumber,
		arg.LimitNumber,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []QueryUserDescUsernameRow{}
	for rows.Next() {
		var i QueryUserDescUsernameRow
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Role,
			&i.ParentID,
			&i.ParentName,
			&i.EmployeeID,
			&i.EmployeeName,
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

const updateUser = `-- name: UpdateUser :one
UPDATE
    "user"
SET
    "role" = $1,
    "username" = $2,
    "password" = $3
WHERE
    "id" = $4
RETURNING id, role, username, password
`

type UpdateUserParams struct {
	Role     NullUserRole
	Username pgtype.Text
	Password pgtype.Text
	ID       int32
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Role,
		arg.Username,
		arg.Password,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Role,
		&i.Username,
		&i.Password,
	)
	return i, err
}
