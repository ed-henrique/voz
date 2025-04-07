package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ed-henrique/voz/internal/logger"
	"github.com/ed-henrique/voz/internal/models"
)

func (s *Server) viewBoard(w http.ResponseWriter, r *http.Request) {
	cards, err := s.q.GetCards(r.Context())
	if err != nil {
		msg := "could not fetch cards"
		logger.Error(msg, slog.String("err", err.Error()))
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	if err := s.templates["board"].Execute(w, cards); err != nil {
		msg := "could not execute board template"
		logger.Error(msg, slog.String("err", err.Error()))
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}

func (s *Server) viewBoardCardsNew(w http.ResponseWriter, r *http.Request) {
	if err := s.templates["board_cards_new"].Execute(w, nil); err != nil {
		msg := "could not execute board_cards_new template"
		logger.Error(msg, slog.String("err", err.Error()))
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}

func (s *Server) apiInsertCard(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		msg := "could not parse form"
		logger.Error(msg, slog.String("err", err.Error()))
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	name := r.PostFormValue("name")
	description := r.PostFormValue("description")

	id, err := s.q.InsertCard(r.Context(), models.InsertCardParams{
		Name:        name,
		Description: description,
	})
	if err != nil {
		msg := "could not insert card"
		logger.Error(msg, slog.String("err", err.Error()))
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/board/cards/%d", id), http.StatusSeeOther)
}
