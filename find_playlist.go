package main

import (
	"github.com/zmb3/spotify"
	"log"
)

func findPlaylist(client *spotify.Client, query string) *spotify.SearchResult {
	search, err := client.Search(query, spotify.SearchTypePlaylist)
	if err != nil {
		log.Fatal(err)
	}
	return search
}

func findDailyDrive(client *spotify.Client) *spotify.SimplePlaylist {
	search := findPlaylist(client, "daily drive")
	for _, playlist := range search.Playlists.Playlists {
		if playlist.Owner.DisplayName == "Spotify" {
			return &playlist
		}
	}
	return nil
}

func findMyDailyDrive(client *spotify.Client) *spotify.SimplePlaylist {
	playlists, err := client.CurrentUsersPlaylists()
	if err != nil {
		log.Fatal(err)
	}

	for page := 1; ; page++ {
		for _, playlist := range playlists.Playlists {
			if playlist.Name == "My Music Daily Drive" {
				return &playlist
			}
		}
		err = client.NextPage(playlists)
		if err == spotify.ErrNoMorePages {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
