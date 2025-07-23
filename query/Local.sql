-- name: CreateLocal :one
INSERT INTO Local (nome) VALUES (?) RETURNING nome;

-- name: ListLocals :many
SELECT * FROM Local;

-- name: UpdateLocal :exec
UPDATE Local SET nome = ? WHERE id = ?;

-- name: DeleteLocal :exec
DELETE FROM Local WHERE nome = ?;

