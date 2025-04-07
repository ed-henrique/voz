package server

import (
	"database/sql"
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

func New(views fs.FS, dbDnsURI string) *Server {
	conn := db.New(dbDnsURI)

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
		tmpl := render.Render(s.views, "board.html", "components/card.html")

		data := models.Board{
			Title: "Board",
			Cards: []models.Card{
				{
					Name: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris vitae diam sit amet mi scelerisque porttitor eget id felis. Suspendisse erat orci, bibendum quis purus ut, venenatis volutpat nibh. Sed pretium ac risus et maximus. Nulla cursus cursus neque ut fermentum. Phasellus neque ipsum, placerat sit amet tempus ut, suscipit a dui. In hac habitasse platea dictumst. Interdum et malesuada fames ac ante ipsum primis in faucibus. Morbi nec ipsum viverra, malesuada nisl sit amet, fringilla turpis.",
					Description: `
Lorem s.muxpsum dolor sit amet, consectetur adipiscing elit. Mauris vitae diam sit amet mi scelerisque porttitor eget id felis. Suspendisse erat orci, bibendum quis purus ut, venenatis volutpat nibh. Sed pretium ac risus et maximus. Nulla cursus cursus neque ut fermentum. Phasellus neque ipsum, placerat sit amet tempus ut, suscipit a dui. In hac habitasse platea dictumst. Interdum et malesuada fames ac ante ipsum primis in faucibus. Morbi nec ipsum viverra, malesuada nisl sit amet, fringilla turpis.

Cras vs.muxnenatis nibh quis libero dignissim ultrices. Nunc rutrum nisl eros, quis blandit ex ultricies id. Curabitur blandit risus et neque malesuada, eget euismod elit faucibus. Mauris eu bibendum est. Proin vel sodales diam. Integer tortor ante, molestie sit amet egestas ac, ultrices id arcu. Nam viverra augue vel placerat pretium. Duis et erat vel eros ultricies fermentum ut eleifend massa. Praesent mattis ac nisl ut semper. Vestibulum pretium justo nibh, a dapibus magna mollis eu. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Interdum et malesuada fames ac ante ipsum primis in faucibus. Donec sed sem nisi.

Quisqus.mux non lorem libero. Donec congue massa tempus nibh finibus maximus. Aenean condimentum consectetur scelerisque. Pellentesque eget imperdiet tortor, eu fermentum enim. Vivamus et laoreet nisi. Praesent vel diam tempor, vestibulum velit in, blandit lorem. Duis convallis orci at nibh dictum lobortis. Nulla nec elit mi. Aenean egestas lectus a ligula aliquet euismod. Quisque ultrices nisl a ex maximus pretium. Duis orci mauris, venenatis vitae sodales id, pretium nec mauris. Vivamus nec euismod metus. Pellentesque sem odio, pulvinar vel iaculis eu, sodales vitae ante. Nulla facilisi. Integer varius lorem purus, quis sagittis dolor commodo vel.

Nulla s.muxliquet nunc mauris, et lacinia lorem eleifend in. Pellentesque semper, neque sit amet pulvinar efficitur, purus lacus placerat orci, id vestibulum tellus mauris sed sapien. Aliquam nec diam eu neque suscipit congue. Etiam nec elit orci. Sed nec dolor sed mauris vehicula vestibulum vitae a nisi. Sed vitae eros massa. Curabitur vitae augue auctor, posuere mi vel, finibus lacus. Nullam commodo lorem enim, in posuere sapien placerat quis. Fusce finibus, erat et rutrum ultricies, est enim tincidunt justo, nec semper massa massa eget ante. In sem quam, mollis ac accumsan quis, imperdiet vitae tellus. Fusce pulvinar et magna vitae ultrices. Suspendisse vitae purus sit amet mauris fermentum egestas ut nec ligula.

Etiam s.muxravida sed massa ut ullamcorper. Cras facilisis urna id euismod efficitur. Nunc luctus sem id posuere blandit. Maecenas nec eros ac magna molestie vehicula. Maecenas viverra consectetur augue, blandit pellentesque libero aliquet sed. Sed eleifend lacinia enim fermentum cursus. Aenean malesuada, metus nec efficitur congue, augue eros volutpat nunc, vel porta erat est interdum magna. Nulla ornare dolor ac risus ultricies, eget efficitur felis sodales. Vivamus facilisis vestibulum lacinia. Pellentesque dapibus ipsum orci, vel placerat nisi sagittis non. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Nunc in mauris feugiat, imperdiet ante vel, ornare quam. Aliquam varius, purus eget blandit congue, augue nisl sagittis diam, at tincidunt velit ex nec neque. Etiam molestie nisi sed lectus porta blandit.`,
					Reaction: models.DOWNVOTTED,
				},
				{
					Title:       "a",
					Description: "b",
					Reaction:    models.DOWNVOTTED,
				},
				{
					Title:       "a",
					Description: "b",
					Reaction:    models.DOWNVOTTED,
					Upvotes:     1423480890,
				},
				{
					Title:       "a",
					Description: "b",
					Reaction:    models.DOWNVOTTED,
				},
			},
		}

		if err := tmpl.Execute(w, data); err != nil {
			errkit.FinalErr(err.Error())
			return
		}
	})

	s.mux.HandleFunc("GET /board/cards/new", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("POST /board/cards/new", func(w http.ResponseWriter, r *http.Request) {})
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
	}
}
