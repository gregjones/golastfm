package lastfm

import (
	"encoding/xml"
)

// AlbumInfo is a type representing the metadata and tracklist for an album.
// When returned as part of search results, the AlbumInfo objects will be incomplete, 
// importantly they don't contain the track-listing.
// 
// See http://www.last.fm/api/show/album.search
type AlbumInfo struct {
	XMLName     xml.Name     `xml:"album"`
	Name        string       `xml:"name"`
	Artist      string       `xml:"artist"`
	ID          string       `xml:"id"`
	MBID        string       `xml:"mbid"`
	Url         string       `xml:"url"`
	ReleaseDate string       `xml:"releaseDate"`
	Images      []Image      `xml:"image"`
	Listeners   string       `xml:"listeners"`
	Playcount   string       `xml:"playcount"`
	Tracks      []AlbumTrack `xml:"tracks>track"`
	TopTags     []Tag        `xml:"toptags>tag"`
	Wiki        Wiki         `xml:"wiki"`
}

type AlbumTrack struct {
	// TrackNumber is the "rank" attribute in the API response
	TrackNumber int    `xml:"rank,attr"`
	Name        string `xml:"name"`
	// Duration is the length of the track in seconds
	Duration int         `xml:"duration"`
	MBID     string      `xml:"mbid"`
	Url      string      `xml:"url"`
	Artist   TrackArtist `xml:"artist"`
}

// GetFullInfo will use either the musicbrainz ID or the name/artist to fetch the full info for the track
func (at AlbumTrack) GetFullInfo(c *Client, username, lang string) (*TrackInfo, error) {
	if at.MBID != "" {
		return c.Track().GetInfoByMBID(at.MBID, username)
	}
	return c.Track().GetInfo(at.Artist.Name, at.Name, false, username)
}

// Tag
type Tag struct {
	Name string `xml:"name"`
	// Count is a weighted count of how often the tag has been applied, with a maximum value of 100
	// Only present in GetTopTags calls
	Count int    `xml:"count"`
	Url   string `xml:"url"`
}

// Wiki is a type to represent the wiki/bio properties found in albums, artists and tracks
type Wiki struct {
	Published string `xml:"published"`
	Summary   string `xml:"summary"`
	Content   string `xml:"content"`
	// Only Artists have a YearFormed
	YearFormed string `xml:"yearformed"`
}

type Image struct {
	Url  string `xml:",chardata"`
	Size string `xml:"size,attr"`
}

// Affiliation is a type that contains the details of a "buy link", for either a physical or e-commerce retailer.
type Affiliation struct {
	SupplierName string `xml:"supplierName"`
	Price        struct {
		Currency  string `xml:"currency:`
		Amount    string `xml:"amount"`
		Formatted string `xml:"formatted"`
	} `xml:"price"`
	BuyLink      string `xml:"buyLink"`
	SupplierIcon string `xml:"supplierIcon"`
	IsSearch     string `xml:"isSearch"`
}

// BuyLinks contains two lists of Affiliations, differentiating between Physical and Download retailers
type BuyLinks struct {
	Physicals []Affiliation `xml:"physicals>affiliation"`
	Downloads []Affiliation `xml:"downloads>affiliation"`
}

// A Shout is a comment on an album, artist or track
type Shout struct {
	Body   string `xml:"body"`
	Author string `xml:"author"`
	Date   string `xml:"date"`
}

// AlbumSearchResults contains a list of AlbumInfos, along with the totals and paging information for the call
// 
// The AlbumInfo values contained won't be complete. Specifically, they don't include track listings.
// To get the full version, call GetInfo with the MBID
type AlbumSearchResults struct {
	TotalResults int         `xml:"totalResults"`
	StartIndex   int         `xml:"startIndex"`
	ItemsPerPage int         `xml:"itemsPerPage"`
	Albums       []AlbumInfo `xml:"albummatches>album"`
}

type ArtistInfo struct {
	Name       string  `xml:"name"`
	MBID       string  `xml:"name"`
	Url        string  `xml:"url"`
	Images     []Image `xml:"image"`
	Streamable string  `xml:"streamable"`
	Stats      struct {
		Listeners int `xml:"listeners"`
		Plays     int `xml:"plays"`
	} `xml:"stats"`
	// SimilarArtists is a list of incomplete ArtistInfos
	SimilarArtists []ArtistInfo `xml:"similar>artist"`
	Tags           []Tag        `xml:"tags>tag"`
	Bio            Wiki         `xml:"bio:"`
}

type SimilarArtist struct {
	Name   string  `xml:"name"`
	MBID   string  `xml:"mbid"`
	Url    string  `xml:"url"`
	Images []Image `xml:"image"`
}

// GetFullInfo returns the full ArtistInfo for the track's artist
func (a *SimilarArtist) GetFullInfo(c *Client, username, lang string) (*ArtistInfo, error) {
	if a.MBID != "" {
		return c.Artist().GetInfoByMBID(a.MBID, username, lang)
	}
	return c.Artist().GetInfo(a.Name, false, username, lang)
}

type TrackInfo struct {
	// No idea what this is used for
	ID   string `xml:"id"`
	Name string `xml:"name"`
	// Duration is the length of the track in seconds
	Duration   int         `xml:"duration"`
	Streamable string      `xml:"streamable"`
	Listeners  int         `xml:"listeners"`
	Playcount  int         `xml:"playcount"`
	MBID       string      `xml:"mbid"`
	Url        string      `xml:"url"`
	Artist     TrackArtist `xml:"artist"`
	Album      TrackAlbum  `xml:"album"`
	TopTags    []Tag       `xml:"toptags>tag"`
	Wiki       Wiki        `xml:"wiki:"`
}

// TrackAlbum contains a subset of the properties of the Album that a track belongs to
type TrackAlbum struct {
	// TrackNumber is the "position" attribute in the response
	TrackNumber int     `xml:"position,attr"`
	Artist      string  `xml:"artist"`
	Title       string  `xml:"title"`
	MBID        string  `xml:"mbid"`
	Url         string  `xml:"string"`
	Images      []Image `xml:"image"`
}

// GetFullInfo returns the full AlbumInfo for the track's album
func (ta TrackAlbum) GetFullInfo(c *Client, username, lang string) (*AlbumInfo, error) {
	if ta.MBID != "" {
		return c.Album().GetInfoByMBID(ta.MBID, username, lang)
	}
	return c.Album().GetInfo(ta.Artist, ta.Title, false, username, lang)
}

// TrackArtist contains the properties of an Artist when found inside a TrackInfo (or AlbumTrackInfo etc.)
type TrackArtist struct {
	Name string `xml:"name"`
	MBID string `xml:"mbid"`
	Url  string `xml:"url"`
}

// GetFullInfo returns the full ArtistInfo for the track's artist
func (a *TrackArtist) GetFullInfo(c *Client, username, lang string) (*ArtistInfo, error) {
	if a.MBID != "" {
		return c.Artist().GetInfoByMBID(a.MBID, username, lang)
	}
	return c.Artist().GetInfo(a.Name, false, username, lang)
}
