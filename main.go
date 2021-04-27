package main

import (
	"fmt"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"

	"github.com/robfig/cron"
	"github.com/zmb3/spotify"
)

var redirectURI = os.Getenv("REDIRECT_URL")

func makeAuth() spotify.Authenticator {
	auth := spotify.NewAuthenticator(redirectURI,
		spotify.ScopePlaylistModifyPrivate,
		spotify.ScopePlaylistReadPrivate,
		spotify.ScopePlaylistModifyPublic,
		spotify.ScopePlaylistReadCollaborative)
	auth.SetAuthInfo(os.Getenv("SPOTIFY_ID"), os.Getenv("SPOTIFY_SECRET"))
	return auth
}

var (
	auth   = makeAuth()
	state  = "abc12342"
	ch     = make(chan spotify.Client)
	client spotify.Client
	token  *oauth2.Token
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
		client = <-ch
		user, err := client.CurrentUser()
		if err != nil {
			log.Fatal(err)
		}

		// Extract the token.
		token, err = client.Token()
		if err != nil {
			log.Fatal(err)
		}
		// Load token again and create client from it.
		client = clientFromToken(token)

		user, err = client.CurrentUser()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("You are logged in as:", user.ID)

		updateDailyDrive()

		c := cron.New()
		// Running every day at 10AM
		_ = c.AddFunc("0 0 10 * * *", updateDailyDrive)
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
	ch <- client
}

func clientFromToken(token *oauth2.Token) spotify.Client {
	auth := makeAuth()
	return auth.NewClient(token)
}
