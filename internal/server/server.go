package server

import (
	"database/sql"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/ed-henrique/voz/internal/db"
	"github.com/ed-henrique/voz/internal/errkit"
	"github.com/ed-henrique/voz/internal/models"
	"github.com/ed-henrique/voz/internal/render"
)

type Server struct {
	mux   *http.ServeMux
	conn  *sql.DB
	views fs.FS

	q         *models.Queries
	templates map[string]*template.Template
}

func New(views fs.FS, dbDsnURI string) *Server {
	conn := db.New(dbDsnURI)

	return &Server{
		mux:       http.NewServeMux(),
		conn:      conn,
		views:     views,
		q:         models.New(conn),
		templates: make(map[string]*template.Template),
	}
}

func (s *Server) LoadTemplates() {
	templates := map[string][]string{
		"board":           {"components/base.html", "board.html", "components/card.html"},
		"board_cards_new": {"components/base.html", "board_cards_new.html"},
	}

	for name, tmpls := range templates {
		s.templates[name] = render.Load(s.views, tmpls...)
	}
}

func (s *Server) Routes() {
	s.mux.HandleFunc("GET /board", s.viewBoard)
	s.mux.HandleFunc("GET /board/cards/new", s.viewBoardCardsNew)
	s.mux.HandleFunc("POST /board/cards/new", s.apiInsertCard)

	// TODO
	// Auth
	s.mux.HandleFunc("POST /auth/simple/login", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("POST /auth/simple/signup", func(w http.ResponseWriter, r *http.Request) {})

	// OAuth2
	s.mux.HandleFunc("POST /auth/google/login", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("POST /auth/google/callback", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("POST /auth/logout", func(w http.ResponseWriter, r *http.Request) {})

	// Landpage
	s.mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {})

	// Read, Update and Delete Card
	s.mux.HandleFunc("GET /board/cards/:card_id", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("PATCH /board/cards/:card_id", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("DELETE /board/cards/:card_id", func(w http.ResponseWriter, r *http.Request) {})

	// Upvote/Downvote Card
	s.mux.HandleFunc("POST /board/cards/:card_id/upvotes", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("DELETE /board/cards/:card_id/upvotes", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("POST /board/cards/:card_id/downvotes", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("DELETE /board/cards/:card_id/downvotes", func(w http.ResponseWriter, r *http.Request) {})

	// Create, Update and Delete Comment
	s.mux.HandleFunc("POST /board/cards/:card_id/comments", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("PATCH /board/cards/:card_id/comments/:comment_id", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("DELETE /board/cards/:card_id/:comment_id", func(w http.ResponseWriter, r *http.Request) {})
}

func (s *Server) Run() {
	if err := http.ListenAndServe(":8080", s.mux); err != nil {
		errkit.FinalErr(err)
		return
	}
}
