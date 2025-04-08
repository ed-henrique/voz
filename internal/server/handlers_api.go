package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/ed-henrique/voz/internal/logger"
	"github.com/ed-henrique/voz/internal/models"
)

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

func (s *Server) apiInsertUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		msg := "could not parse form"
		logger.Error(msg, slog.String("err", err.Error()))
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	userTypeIDRaw := r.PostFormValue("user_type_id")
	userTypeID, err := strconv.ParseInt(userTypeIDRaw, 10, 64)
	if err != nil {
		msg := "could not parse user type id from string to int"
		logger.Debug(msg, slog.String("err", err.Error()))
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	_, err = s.q.InsertUser(r.Context(), models.InsertUserParams{
		Name:       name,
		Email:      email,
		Username:   username,
		Password:   password,
		UserTypeID: userTypeID,
	})
	if err != nil {
		msg := "could not insert user"
		logger.Error(msg, slog.String("err", err.Error()))
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/auth/simple/login", http.StatusSeeOther)
}
