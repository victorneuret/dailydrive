package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/robfig/cron"
	"github.com/zmb3/spotify"
)

// redirectURI is the OAuth redirect URI for the application.
// You must register an application at Spotify's developer portal
// and enter this value.
const redirectURI = "http://localhost:2021/callback"

var (
	auth = spotify.NewAuthenticator(
		redirectURI,
		spotify.ScopePlaylistModifyPrivate,
		spotify.ScopePlaylistReadPrivate,
		spotify.ScopePlaylistModifyPublic,
		spotify.ScopePlaylistReadCollaborative)
	ch     = make(chan *spotify.Client)
	state  = "abc12342"
	client *spotify.Client
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// first start an HTTP server
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	go func() {
		// wait for auth to complete
		client = <-ch

		// use the client to make calls that require authorization
		user, err := client.CurrentUser()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("You are logged in as:", user.ID)
		updateDailyDrive()

		c := cron.New()
		// Running every day at 10AM
		_ = c.AddFunc("0 10 * * *", updateDailyDrive)
		c.Start()
	}()

	http.ListenAndServe(":2021", nil)
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}
	// use the token to get an authenticated client
	client := auth.NewClient(tok)
	_, _ = fmt.Fprintf(w, "Login Completed!")
	ch <- &client
}
