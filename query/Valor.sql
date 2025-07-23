-- name: CreateValor :one
INSERT INTO Valor (valor, data_inicio, data_fim) VALUES (?, ?, ?) RETURNING *;

-- name: ListValors :many
SELECT * FROM Valor;

-- name: UpdateValor :exec
UPDATE Valor SET valor = ?, data_inicio = ?, data_fim = ? WHERE id = ?;

-- name: DeleteValor :exec
DELETE FROM Valor WHERE id = ?;

