package server

import (
	"database/sql"
	"fmt"
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

	q *models.Queries
}

func New(views fs.FS, dbDsnURI string) *Server {
	conn := db.New(dbDsnURI)

	return &Server{
		mux:   http.NewServeMux(),
		conn:  conn,
		views: views,
		q:     models.New(conn),
	}
}

func (s *Server) Routes() {
	s.mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("POST /auth/google/login", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("POST /auth/google/callback", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("POST /auth/logout", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("GET /board", func(w http.ResponseWriter, r *http.Request) {
		tmpl := render.Render(s.views, "components/base.html", "board.html", "components/card.html")

		cards, err := s.q.GetCards(r.Context())
		if err != nil {
			errkit.FinalErr(err.Error())
			return
		}

		data := map[string]any{
			"Cards": cards,
		}

		if err := tmpl.Execute(w, data); err != nil {
			errkit.FinalErr(err.Error())
			return
		}
	})

	s.mux.HandleFunc("GET /board/cards/new", func(w http.ResponseWriter, r *http.Request) {
		tmpl := render.Render(s.views, "components/base.html", "board_cards_new.html")

		if err := tmpl.Execute(w, nil); err != nil {
			errkit.FinalErr(err.Error())
			return
		}
	})

	s.mux.HandleFunc("POST /board/cards/new", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			errkit.FinalErr(err.Error())
			return
		}

		name := r.PostFormValue("name")
		description := r.PostFormValue("description")

		id, err := s.q.InsertCard(r.Context(), models.InsertCardParams{
			Name:        name,
			Description: description,
			UserID:      1,
		})
		if err != nil {
			errkit.FinalErr(err.Error())
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/board/cards/%d", id), http.StatusSeeOther)
	})

	s.mux.HandleFunc("GET /board/cards/:card_id", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("PATCH /board/cards/:card_id", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("DELETE /board/cards/:card_id", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("PUT /board/cards/:card_id/react", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("POST /board/cards/:card_id/comments", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("PATCH /board/cards/:card_id/comments/:comment_id", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("DELETE /board/cards/:card_id/:comment_id", func(w http.ResponseWriter, r *http.Request) {})
}

func (s *Server) Run() {
	if err := http.ListenAndServe(":8080", s.mux); err != nil {
		errkit.FinalErr(err.Error())
		return
	}
}
