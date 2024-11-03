-- name: CreateSantri :one
INSERT INTO
    "santri" (
        "nis",
        "name",
        "gender",
        "is_active",
        "generation",
        "photo",
        "occupation_id",
        "parent_id"
    )
VALUES
    (
        @nis,
        @name,
        @gender,
        @is_active,
        @generation,
        sqlc.narg(photo) :: text,
        @occupation_id,
        @parent_id
    ) RETURNING *;

-- name: ListSantriAscName :many
SELECT
    "santri".*,
    "parent"."id" AS "parent_id",
    "parent"."name" AS "parent_name",
    "parent"."wa_phone" AS "parent_wa_phone",
    "santri_occupation"."id" AS "occupation_id",
    "santri_occupation"."name" AS "occupation_name"
FROM
    "santri"
    LEFT JOIN "parent" ON "santri"."parent_id" = "parent"."id"
    LEFT JOIN "santri_occupation" ON "santri"."occupation_id" = "santri_occupation"."id"
WHERE
    (
        sqlc.narg(q) :: text IS NULL
        OR "santri"."name" ILIKE '%' || sqlc.narg(q) || '%'
        OR "santri"."nis" ILIKE '%' || sqlc.narg(q) || '%'
    )
    AND (
        sqlc.narg(parent_id) :: integer IS NULL
        OR "parent_id" = sqlc.narg(parent_id) :: integer
    )
    AND (
        sqlc.narg(occupation_id) :: integer IS NULL
        OR "occupation_id" = sqlc.narg(occupation_id) :: integer
    )
    AND (
        sqlc.narg(generation) :: integer IS NULL
        OR "generation" = sqlc.narg(generation) :: integer
    )
ORDER BY
    "santri"."name" ASC
LIMIT
    @limit_number OFFSET @offset_number;

-- name: ListSantriAscNis :many
SELECT
    "santri".*,
    "parent"."id" AS "parent_id",
    "parent"."name" AS "parent_name",
    "parent"."wa_phone" AS "parent_wa_phone",
    "santri_occupation"."id" AS "occupation_id",
    "santri_occupation"."name" AS "occupation_name"
FROM
    "santri"
    LEFT JOIN "parent" ON "santri"."parent_id" = "parent"."id"
    LEFT JOIN "santri_occupation" ON "santri"."occupation_id" = "santri_occupation"."id"
WHERE
    (
        sqlc.narg(q) :: text IS NULL
        OR "santri"."name" ILIKE '%' || sqlc.narg(q) || '%'
        OR "santri"."nis" ILIKE '%' || sqlc.narg(q) || '%'
    )
    AND (
        sqlc.narg(parent_id) :: integer IS NULL
        OR "parent_id" = sqlc.narg(parent_id) :: integer
    )
    AND (
        sqlc.narg(occupation_id) :: integer IS NULL
        OR "occupation_id" = sqlc.narg(occupation_id) :: integer
    )
    AND (
        sqlc.narg(generation) :: integer IS NULL
        OR "generation" = sqlc.narg(generation) :: integer
    )
ORDER BY
    "nis" ASC
LIMIT
    @limit_number OFFSET @offset_number;

-- name: ListSantriAscGeneration :many
SELECT
    "santri".*,
    "parent"."id" AS "parent_id",
    "parent"."name" AS "parent_name",
    "parent"."wa_phone" AS "parent_wa_phone",
    "santri_occupation"."id" AS "occupation_id",
    "santri_occupation"."name" AS "occupation_name"
FROM
    "santri"
    LEFT JOIN "parent" ON "santri"."parent_id" = "parent"."id"
    LEFT JOIN "santri_occupation" ON "santri"."occupation_id" = "santri_occupation"."id"
WHERE
    (
        sqlc.narg(q) :: text IS NULL
        OR "santri"."name" ILIKE '%' || sqlc.narg(q) || '%'
        OR "santri"."nis" ILIKE '%' || sqlc.narg(q) || '%'
    )
    AND (
        sqlc.narg(parent_id) :: integer IS NULL
        OR "parent_id" = sqlc.narg(parent_id) :: integer
    )
    AND (
        sqlc.narg(occupation_id) :: integer IS NULL
        OR "occupation_id" = sqlc.narg(occupation_id) :: integer
    )
    AND (
        sqlc.narg(generation) :: integer IS NULL
        OR "generation" = sqlc.narg(generation) :: integer
    )
ORDER BY
    "generation" ASC
LIMIT
    @limit_number OFFSET @offset_number;

-- name: ListSantriAscOccupation :many
SELECT
    "santri".*,
    "parent"."id" AS "parent_id",
    "parent"."name" AS "parent_name",
    "parent"."wa_phone" AS "parent_wa_phone",
    "santri_occupation"."id" AS "occupation_id",
    "santri_occupation"."name" AS "occupation_name"
FROM
    "santri"
    LEFT JOIN "parent" ON "santri"."parent_id" = "parent"."id"
    LEFT JOIN "santri_occupation" ON "santri"."occupation_id" = "santri_occupation"."id"
WHERE
    (
        sqlc.narg(q) :: text IS NULL
        OR "santri"."name" ILIKE '%' || sqlc.narg(q) || '%'
        OR "santri"."nis" ILIKE '%' || sqlc.narg(q) || '%'
    )
    AND (
        sqlc.narg(parent_id) :: integer IS NULL
        OR "parent_id" = sqlc.narg(parent_id) :: integer
    )
    AND (
        sqlc.narg(occupation_id) :: integer IS NULL
        OR "occupation_id" = sqlc.narg(occupation_id) :: integer
    )
    AND (
        sqlc.narg(generation) :: integer IS NULL
        OR "generation" = sqlc.narg(generation) :: integer
    )
ORDER BY
    "occupation_id" ASC
LIMIT
    @limit_number OFFSET @offset_number;

-- name: GetSantri :one
SELECT
    "santri".*,
    "parent"."id" AS "parent_id",
    "parent"."name" AS "parent_name",
    "parent"."wa_phone" AS "parent_wa_phone",
    "parent"."address" AS "parentAddress",
    "parent"."photo" AS "parentPhoto"
FROM
    "santri"
    LEFT JOIN "parent" ON "santri"."parent_id" = "parent"."id"
    LEFT JOIN "santri_occupation" ON "santri"."occupation_id" = "santri_occupation"."id"
WHERE
    "santri"."id" = @id;

-- name: UpdateSantri :one
UPDATE
    "santri"
SET
    "nis" = @nis,
    "name" = @name,
    "generation" = @generation,
    "is_active" = @is_active :: boolean,
    "photo" = sqlc.narg(photo) :: text,
    "occupation_id" = @occupation_id,
    "parent_id" = sqlc.narg(parent_id) :: integer
WHERE
    "id" = @id RETURNING *;

-- name: DeleteSantri :one
DELETE FROM
    "santri"
WHERE
    "id" = @id RETURNING *;