package main

import (
	"github.com/zmb3/spotify"
	"log"
)

func createMyDailyDrivePlaylist(client *spotify.Client, user *spotify.PrivateUser) {
	if _, err := client.CreatePlaylistForUser(
		user.ID,
		"My Music Daily Drive",
		"The Spotify's Daily Drive playlist without the podcasts",
		false,
	); err != nil {
		log.Fatal(err)
	}
}
