package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

//go:embed templates/*
var templates embed.FS

type Reaction = uint8

const (
	NONE Reaction = iota
	UPVOTTED
	DOWNVOTTED
)

type Card struct {
	Title       string
	Description string
	Upvotes     int
	Downvotes   int
	Comments    []string

	Reaction
}

type Board struct {
	Title string
	Cards []Card
}

func FinalErr(msg string) {
	fmt.Fprintln(os.Stderr, "Error: "+msg)
	os.Exit(1)
}

func getTemplates(names ...string) *template.Template {
	templatePaths := make([]string, 0)
	for _, tmplName := range names {
		templatePaths = append(templatePaths, "templates/"+tmplName)
	}

	tmpl, err := template.ParseFS(templates, templatePaths...)
	if err != nil {
		FinalErr(err.Error())
		return nil
	}

	return tmpl
}

func main() {
	mux := http.DefaultServeMux

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("POST /auth/google/login", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("POST /auth/google/callback", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("POST /auth/logout", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("GET /board", func(w http.ResponseWriter, r *http.Request) {
		tmpl := getTemplates("board.html", "board_card.html")

		data := Board{
			Title: "Board",
			Cards: []Card{
				{
					Title: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris vitae diam sit amet mi scelerisque porttitor eget id felis. Suspendisse erat orci, bibendum quis purus ut, venenatis volutpat nibh. Sed pretium ac risus et maximus. Nulla cursus cursus neque ut fermentum. Phasellus neque ipsum, placerat sit amet tempus ut, suscipit a dui. In hac habitasse platea dictumst. Interdum et malesuada fames ac ante ipsum primis in faucibus. Morbi nec ipsum viverra, malesuada nisl sit amet, fringilla turpis.",
					Description: `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris vitae diam sit amet mi scelerisque porttitor eget id felis. Suspendisse erat orci, bibendum quis purus ut, venenatis volutpat nibh. Sed pretium ac risus et maximus. Nulla cursus cursus neque ut fermentum. Phasellus neque ipsum, placerat sit amet tempus ut, suscipit a dui. In hac habitasse platea dictumst. Interdum et malesuada fames ac ante ipsum primis in faucibus. Morbi nec ipsum viverra, malesuada nisl sit amet, fringilla turpis.

Cras venenatis nibh quis libero dignissim ultrices. Nunc rutrum nisl eros, quis blandit ex ultricies id. Curabitur blandit risus et neque malesuada, eget euismod elit faucibus. Mauris eu bibendum est. Proin vel sodales diam. Integer tortor ante, molestie sit amet egestas ac, ultrices id arcu. Nam viverra augue vel placerat pretium. Duis et erat vel eros ultricies fermentum ut eleifend massa. Praesent mattis ac nisl ut semper. Vestibulum pretium justo nibh, a dapibus magna mollis eu. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Interdum et malesuada fames ac ante ipsum primis in faucibus. Donec sed sem nisi.

Quisque non lorem libero. Donec congue massa tempus nibh finibus maximus. Aenean condimentum consectetur scelerisque. Pellentesque eget imperdiet tortor, eu fermentum enim. Vivamus et laoreet nisi. Praesent vel diam tempor, vestibulum velit in, blandit lorem. Duis convallis orci at nibh dictum lobortis. Nulla nec elit mi. Aenean egestas lectus a ligula aliquet euismod. Quisque ultrices nisl a ex maximus pretium. Duis orci mauris, venenatis vitae sodales id, pretium nec mauris. Vivamus nec euismod metus. Pellentesque sem odio, pulvinar vel iaculis eu, sodales vitae ante. Nulla facilisi. Integer varius lorem purus, quis sagittis dolor commodo vel.

Nulla aliquet nunc mauris, et lacinia lorem eleifend in. Pellentesque semper, neque sit amet pulvinar efficitur, purus lacus placerat orci, id vestibulum tellus mauris sed sapien. Aliquam nec diam eu neque suscipit congue. Etiam nec elit orci. Sed nec dolor sed mauris vehicula vestibulum vitae a nisi. Sed vitae eros massa. Curabitur vitae augue auctor, posuere mi vel, finibus lacus. Nullam commodo lorem enim, in posuere sapien placerat quis. Fusce finibus, erat et rutrum ultricies, est enim tincidunt justo, nec semper massa massa eget ante. In sem quam, mollis ac accumsan quis, imperdiet vitae tellus. Fusce pulvinar et magna vitae ultrices. Suspendisse vitae purus sit amet mauris fermentum egestas ut nec ligula.

Etiam gravida sed massa ut ullamcorper. Cras facilisis urna id euismod efficitur. Nunc luctus sem id posuere blandit. Maecenas nec eros ac magna molestie vehicula. Maecenas viverra consectetur augue, blandit pellentesque libero aliquet sed. Sed eleifend lacinia enim fermentum cursus. Aenean malesuada, metus nec efficitur congue, augue eros volutpat nunc, vel porta erat est interdum magna. Nulla ornare dolor ac risus ultricies, eget efficitur felis sodales. Vivamus facilisis vestibulum lacinia. Pellentesque dapibus ipsum orci, vel placerat nisi sagittis non. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Nunc in mauris feugiat, imperdiet ante vel, ornare quam. Aliquam varius, purus eget blandit congue, augue nisl sagittis diam, at tincidunt velit ex nec neque. Etiam molestie nisi sed lectus porta blandit.`,
					Reaction: DOWNVOTTED,
				},
				{
					Title:       "a",
					Description: "b",
					Reaction:    DOWNVOTTED,
				},
				{
					Title:       "a",
					Description: "b",
					Reaction:    DOWNVOTTED,
				},
				{
					Title:       "a",
					Description: "b",
					Reaction:    DOWNVOTTED,
				},
			},
		}

		tmpl.Execute(w, data)
	})
	mux.HandleFunc("GET /board/cards/new", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("POST /board/cards/new", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("GET /board/cards/:card_id", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("PATCH /board/cards/:card_id", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("DELETE /board/cards/:card_id", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("PUT /board/cards/:card_id/react", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("POST /board/cards/:card_id/comments", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("PATCH /board/cards/:card_id/comments/:comment_id", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("DELETE /board/cards/:card_id/:comment_id", func(w http.ResponseWriter, r *http.Request) {})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		FinalErr(err.Error())
	}
}
