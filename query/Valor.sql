-- name: CreateValor :exec
INSERT INTO Valor (valor, data_inicio, data_fim) VALUES (?, ?, ?);

-- name: ListValors :many
SELECT * FROM Valor;

-- name: UpdateValor :exec
UPDATE Valor SET valor = ?, data_inicio = ?, data_fim = ? WHERE id = ?;

-- name: DeleteValor :exec
DELETE FROM Valor WHERE id = ?;

