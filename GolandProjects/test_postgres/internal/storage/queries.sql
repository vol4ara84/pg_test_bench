-- noinspection SqlNoDataSourceInspectionForFile

-- name: InsertFiles :one
INSERT INTO files (mask)
VALUES ($1)
RETURNING id
;

-- name: UpdateFileMask :exec
UPDATE files set mask = mask << 1 where id=$1;
