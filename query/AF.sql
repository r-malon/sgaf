-- name: CreateAF :one
INSERT INTO AF (numero, fornecedor, descricao, data_inicial, data_final, status) VALUES (?, ?, ?, ?, ?, ?) RETURNING *;

-- name: ListAFs :many
SELECT * FROM AF ORDER BY numero;

-- name: UpdateAF :exec
UPDATE AF SET numero = ? WHERE numero = ?;

-- name: DeleteAF :exec
DELETE FROM AF WHERE numero = ?;

