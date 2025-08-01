-- name: CreateItem :exec
INSERT INTO Item (descricao, banda_maxima, banda_instalada, data_instalacao) VALUES (?, ?, ?, ?);

-- name: ListItems :many
SELECT * FROM Item;

-- name: UpdateItem :exec
UPDATE Item SET descricao = ?, banda_maxima = ?, banda_instalada = ?, data_instalacao = ? WHERE id = ?;

-- name: DeleteItem :exec
DELETE FROM Item WHERE id = ?;
