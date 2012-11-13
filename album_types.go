package lastfm

import (
	"encoding/xml"
)

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
		Rank     int    `xml:"rank,attr"`
		Name     string `xml:"name"`
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

type affiliation struct {
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

type BuyLinks struct {
	Physicals []affiliation `xml:"physicals>affiliation"`
	Downloads []affiliation `xml:"downloads>affiliation"`
}

type Shout struct {
	Body   string `xml:"body"`
	Author string `xml:"author"`
	Date   string `xml:"date"`
}
