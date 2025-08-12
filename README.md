# Sistema de Gestão de Autorizações de Fornecimento

## ER model
![SGAF ER model](sgaf.svg)

## Setup
```
for f in schema/*; do
	sqlite3 test.db < f
done
sqlite-history test.db AF Item Local Valor
sqlc-1.25 generate
pandoc SPECIFICATION.md --pdf-engine=weasyprint -o SPECIFICATION.pdf
go build -ldflags=-s
```
