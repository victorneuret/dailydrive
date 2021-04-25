package main

import (
	"github.com/zmb3/spotify"
	"log"
)

func clearMyDailyDrive(client *spotify.Client, playlist *spotify.SimplePlaylist) {
	IDs := getPlaylistTracksIDs(client, playlist)

	if len(IDs) == 0 {
		return
	}
	if _, err := client.RemoveTracksFromPlaylist(playlist.ID, IDs...); err != nil {
		log.Fatal(err)
	}
}
