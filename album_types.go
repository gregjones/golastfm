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
	XMLName     xml.Name `xml:"album"`
	Name        string   `xml:"name"`
	Artist      string   `xml:"artist"`
	ID          string   `xml:"id"`
	MBID        string   `xml:"mbid"`
	Url         string   `xml:"url"`
	ReleaseDate string   `xml:"releaseDate"`
	Images      []struct {
		Url  string `xml:",chardata"`
		Size string `xml:"size,attr"`
	} `xml:"image"`
	Listeners string `xml:"listeners"`
	Playcount string `xml:"playcount"`
	Tracks    []struct {
		Rank int    `xml:"rank,attr"`
		Name string `xml:"name"`
		// Duration is the length of the track in seconds
		Duration int    `xml:"duration"`
		MBID     string `xml:"mbid"`
		Url      string `xml:"url"`
		Artist   struct {
			Name string `xml:"name"`
			MBID string `xml:"mbid"`
			Url  string `xml:"url"`
		} `xml:"artist"`
	} `xml:"tracks>track"`
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

// Tag
type Tag struct {
	Name string `xml:"name"`
	// Count is a weighted count of how often the tag has been applied, with a maximum value of 100
	Count int    `xml:"count"`
	Url   string `xml:"url"`
}

// AlbumSearchResults contains a list of AlbumInfos, along with the totals and paging information for the call
// 
// The AlbumInfo values contained won't be complete. Specifically, they don't include track listings.
type AlbumSearchResults struct {
	TotalResults int         `xml:"totalResults"`
	StartIndex   int         `xml:"startIndex"`
	ItemsPerPage int         `xml:"itemsPerPage"`
	Albums       []AlbumInfo `xml:"albummatches>album"`
}
