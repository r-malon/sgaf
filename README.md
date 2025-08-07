# Sistema de Gestão de Autorizações de Fornecimento
---
Use BEGIN IMMEDIATE transactions and STRICT tables.
```
writeDB.SetMaxOpenConns(1)
readDB.SetMaxOpenConns(max(4, runtime.NumCPU()))
```
## Setup
```
export SGAF_DB
export SGAF_PORT
for f in schema/*; do
	sqlite3 test.db < f
done
python3 -m sqlite_history test.db AF Item Local Valor
sqlc-1.25 generate
go build -ldflags=-s
```
