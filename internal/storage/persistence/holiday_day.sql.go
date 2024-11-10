// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: holiday_day.sql

package persistence

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createHolidayDay = `-- name: CreateHolidayDay :one
INSERT INTO
    "holiday_day" ("date", "holiday_id")
VALUES
    ($1, $2) RETURNING id, date, holiday_id
`

type CreateHolidayDayParams struct {
	Date      pgtype.Date `db:"date"`
	HolidayID int32       `db:"holiday_id"`
}

func (q *Queries) CreateHolidayDay(ctx context.Context, arg CreateHolidayDayParams) (HolidayDay, error) {
	row := q.db.QueryRow(ctx, createHolidayDay, arg.Date, arg.HolidayID)
	var i HolidayDay
	err := row.Scan(&i.ID, &i.Date, &i.HolidayID)
	return i, err
}

const deleteHolidayDay = `-- name: DeleteHolidayDay :one
DELETE FROM
    "holiday_day"
WHERE
    "id" = $1 RETURNING id, date, holiday_id
`

func (q *Queries) DeleteHolidayDay(ctx context.Context, id int32) (HolidayDay, error) {
	row := q.db.QueryRow(ctx, deleteHolidayDay, id)
	var i HolidayDay
	err := row.Scan(&i.ID, &i.Date, &i.HolidayID)
	return i, err
}

const listHolidayDays = `-- name: ListHolidayDays :many
SELECT
    holiday_day.id, date, holiday_id, holiday.id, name, color, description
FROM
    "holiday_day"
    INNER JOIN "holiday" ON "holiday_day"."holiday_id" = "holiday"."id"
WHERE
    "date" BETWEEN $1 AND $2
    AND (
        $3::integer IS NULL
        OR "holiday_id" = $3
    )
    AND (
        $4::text IS NULL
        OR "holiday"."name" ILIKE '%' || $4 || '%'
    )
    LIMIT $6 OFFSET $5
`

type ListHolidayDaysParams struct {
	FromDate     pgtype.Date `db:"from_date"`
	ToDate       pgtype.Date `db:"to_date"`
	HolidayID    pgtype.Int4 `db:"holiday_id"`
	HolidayName  pgtype.Text `db:"holiday_name"`
	OffsetNumber int32       `db:"offset_number"`
	LimitNumber  int32       `db:"limit_number"`
}

type ListHolidayDaysRow struct {
	ID          int32       `db:"id"`
	Date        pgtype.Date `db:"date"`
	HolidayID   int32       `db:"holiday_id"`
	ID_2        int32       `db:"id_2"`
	Name        string      `db:"name"`
	Color       pgtype.Text `db:"color"`
	Description pgtype.Text `db:"description"`
}

func (q *Queries) ListHolidayDays(ctx context.Context, arg ListHolidayDaysParams) ([]ListHolidayDaysRow, error) {
	rows, err := q.db.Query(ctx, listHolidayDays,
		arg.FromDate,
		arg.ToDate,
		arg.HolidayID,
		arg.HolidayName,
		arg.OffsetNumber,
		arg.LimitNumber,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListHolidayDaysRow{}
	for rows.Next() {
		var i ListHolidayDaysRow
		if err := rows.Scan(
			&i.ID,
			&i.Date,
			&i.HolidayID,
			&i.ID_2,
			&i.Name,
			&i.Color,
			&i.Description,
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
