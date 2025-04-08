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
	"github.com/ed-henrique/voz/internal/views"
)

type Server struct {
	mux   *http.ServeMux
	conn  *sql.DB
	views fs.FS

	q         *models.Queries
	templates map[views.View]*template.Template
}

func New(viewsFS fs.FS, dbDsnURI string) *Server {
	conn := db.New(dbDsnURI)

	return &Server{
		mux:       http.NewServeMux(),
		conn:      conn,
		views:     viewsFS,
		q:         models.New(conn),
		templates: make(map[views.View]*template.Template),
	}
}

func (s *Server) LoadTemplates() {
	templates := map[views.View][]string{
		views.BOARD:           {"components/base.html", "board.html", "components/card.html"},
		views.BOARD_CARDS_NEW: {"components/base.html", "board_cards_new.html"},
		views.SIGNUP:          {"components/base_auth.html", "signup.html"},
	}

	for id, tmpls := range templates {
		s.templates[id] = render.Load(s.views, tmpls...)
	}
}

func (s *Server) Routes() {
	s.mux.HandleFunc("GET /board", s.viewBoard)
	s.mux.HandleFunc("GET /board/cards/new", s.viewBoardCardsNew)
	s.mux.HandleFunc("POST /board/cards/new", s.apiInsertCard)
	s.mux.HandleFunc("GET /auth/simple/signup", s.viewSignUp)
	s.mux.HandleFunc("POST /auth/simple/signup", s.apiInsertUser)

	// TODO
	// Auth
	s.mux.HandleFunc("POST /auth/simple/login", func(w http.ResponseWriter, r *http.Request) {})

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
