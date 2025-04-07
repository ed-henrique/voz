package gen

//go:generate sqlc -f ./internal/models/sqlc.yml generate
//go:generate sqlite3 db.sqlite3 ".read  ./internal/models/schema.sql"
