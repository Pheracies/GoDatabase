-- name: GetItem :one
SELECT *
FROM items
WHERE id = ? LIMIT 1;

-- name: ListItems :many
SELECT *
FROM items;

-- name: CreateItem :one
INSERT INTO items
    (name, amount)
VALUES
    (?, ?)
RETURNING id;

-- name: DeleteItem :exec
DELETE FROM items
WHERE id = ?;
