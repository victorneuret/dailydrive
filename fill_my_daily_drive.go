package main

import (
	"github.com/zmb3/spotify"
	"log"
)

func fillMyDailyDrive(client *spotify.Client, dailyDrive *spotify.SimplePlaylist, myDailyDrive *spotify.SimplePlaylist) {
	IDs := getPlaylistTracksIDs(client, dailyDrive)
	IDs = IDs[1:]

	if _, err := client.AddTracksToPlaylist(myDailyDrive.ID, IDs...); err != nil {
		log.Fatal(err)
	}
}
