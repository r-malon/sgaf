# Sistema de Gestão de Autorizações de Fornecimento
---
```
PRAGMA journal_mode = WAL;
PRAGMA busy_timeout = 5000;
PRAGMA synchronous = NORMAL;
PRAGMA cache_size = 1000000000;
PRAGMA foreign_keys = true;
PRAGMA temp_store = memory;
```
Use BEGIN IMMEDIATE transactions and STRICT tables.
```
writeDB.SetMaxOpenConns(1)
readDB.SetMaxOpenConns(max(4, runtime.NumCPU()))
```
python3 -m sqlite_history test.db AF Item Local Valor
sqlc-1.25 generate
