package main

import (
	"embed"

	"github.com/ed-henrique/voz/internal/server"
)

//go:embed views/*
var views embed.FS

const dbDsnURI = "db.sqlite3"

func main() {
	s := server.New(views, dbDsnURI)
	s.Routes()
	s.Run()
}
