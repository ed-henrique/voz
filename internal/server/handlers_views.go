package server

import (
	"log/slog"
	"net/http"

	"github.com/ed-henrique/voz/internal/logger"
	"github.com/ed-henrique/voz/internal/views"
)

func (s *Server) viewBoard(w http.ResponseWriter, r *http.Request) {
	cards, err := s.q.GetCards(r.Context())
	if err != nil {
		msg := "could not fetch cards"
		logger.Error(msg, slog.String("err", err.Error()))
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	if err := s.templates[views.BOARD].Execute(w, cards); err != nil {
		msg := "could not execute board template"
		logger.Error(msg, slog.String("err", err.Error()))
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}

func (s *Server) viewBoardCardsNew(w http.ResponseWriter, _ *http.Request) {
	if err := s.templates[views.BOARD_CARDS_NEW].Execute(w, nil); err != nil {
		msg := "could not execute board_cards_new template"
		logger.Error(msg, slog.String("err", err.Error()))
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}

func (s *Server) viewSignUp(w http.ResponseWriter, r *http.Request) {
	userTypes, err := s.q.GetUserTypes(r.Context())
	if err != nil {
		msg := "could not fetch user types"
		logger.Error(msg, slog.String("err", err.Error()))
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	if err := s.templates[views.SIGNUP].Execute(w, userTypes); err != nil {
		msg := "could not execute signup template"
		logger.Error(msg, slog.String("err", err.Error()))
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}
