package main

import (
	"flag"
	"fmt"
	"github.com/gregjones/golastfm/lastfm"
)

var (
	artist    = flag.String("artist", "Muse", "The album artist to search for")
	albumName = flag.String("album", "The 2nd Law", "The album name to search for")
	username  = flag.String("username", "", "A username (results will include user's playcounts)")
)

func main() {
	flag.Parse()
	fmt.Println("Example album methods")
	apiKey := "d0a8c6b594b43669503d9f51aaabea22"
	c := lastfm.NewClient(apiKey, "")
	album, err := c.Album().GetInfo(*artist, *albumName, true, *username, "en")
	if err != nil {
		fmt.Println("There was an error")
		panic(err)
	}
	fmt.Printf("%s by %s has %d tracks:\n", album.Name, album.Artist, len(album.Tracks))
	for _, track := range album.Tracks {
		fmt.Printf("%d) %s\n", track.TrackNumber, track.Name)
	}

	tags, err := c.Album().GetTopTagsByMBID(album.MBID)
	if err != nil {
		fmt.Println("There was an error")
		panic(err)
	}
	fmt.Printf("\nSome album tags:\n")
	n := 5
	if len(tags) < n {
		n = len(tags)
	}
	for i, tag := range tags[:n] {
		fmt.Printf("%d) %s (%d%%)\n", i+1, tag.Name, tag.Count)
	}

	shouts, err := c.Album().GetShoutsByMBID(album.MBID, 1)
	if err != nil {
		fmt.Println("There was an error")
		panic(err)
	}
	fmt.Printf("\nWhat people are saying:\n")
	for i, shout := range shouts {
		if i > 10 {
			break
		}
		fmt.Printf("%d) %s (by %s)\n", i+1, shout.Body, shout.Author)
	}
}
