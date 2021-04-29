package main

import (
	"fmt"
	"github.com/zmb3/spotify"
	"log"
)

func getPlaylistTracksIDs(client *spotify.Client, playlist *spotify.SimplePlaylist) []spotify.ID {
	tracks, err := client.GetPlaylistTracks(playlist.ID)
	if err != nil {
		log.Fatal(err)
	}

	var IDs []spotify.ID

	for page := 1; ; page++ {
		for _, track := range tracks.Tracks {
			if track.Track.Type != "show" && track.Track.Type != "episode" {
				IDs = append(IDs, track.Track.ID)
				fmt.Println(track.Track.Name, "-", track.Track.Artists[0].Name)
			}
		}
		err = client.NextPage(tracks)
		if err == spotify.ErrNoMorePages {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	return IDs
}
