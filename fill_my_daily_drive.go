package main

import (
	"fmt"
	"github.com/zmb3/spotify"
	"log"
)

func fillMyDailyDrive(client *spotify.Client, dailyDrive *spotify.SimplePlaylist, myDailyDrive *spotify.SimplePlaylist) {
	IDs := getPlaylistTracksIDs(client, dailyDrive)
	IDs = IDs[1:]

	fmt.Println(IDs)
	if _, err := client.AddTracksToPlaylist(myDailyDrive.ID, IDs...); err != nil {
		log.Fatal(err)
	}
}
