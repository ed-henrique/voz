package db

import (
	"database/sql"

	"github.com/ed-henrique/voz/internal/errkit"
)

var params = "?_pragma=foreign_keys(1)"

func New(dsnURI string) *sql.DB {
	conn, err := sql.Open("sqlite", dsnURI+params)
	if err != nil {
		errkit.FinalErr(err.Error())
	}

	if err := conn.Ping(); err != nil {
		errkit.FinalErr(err.Error())
	}

	return conn
}
