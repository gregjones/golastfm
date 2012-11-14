package lastfm

import ()

type AlbumMethods struct {
	Client *Client
}

func (a AlbumMethods) method(suffix string) string {
	return "album." + suffix
}

// standardQuery returns a map of the parameters that most of the album API methods expect to receive.
func (a AlbumMethods) standardQuery(artist, album string, autocorrect bool) map[string]string {
	return map[string]string{
		"artist":      artist,
		"album":       album,
		"autocorrect": boolToString(autocorrect),
	}
}

func (a *AlbumMethods) getBuyLinks(query map[string]string) (*BuyLinks, error) {
	wrapper := struct {
		Affiliations *BuyLinks `xml:"affiliations"`
	}{}
	query["method"] = a.method("getbuylinks")
	err := a.Client.Execute(query, &wrapper)
	if err != nil {
		return nil, err
	}
	return wrapper.Affiliations, nil
}

// GetBuyLinks fetches a list of Buy Links for a particular Album referenced by its name and artist-name.
//
// See http://www.last.fm/api/show/album.getBuylinks
func (a *AlbumMethods) GetBuyLinks(artist, album string, autocorrect bool, country string) (*BuyLinks, error) {
	query := a.standardQuery(artist, album, autocorrect)
	query["country"] = country
	return a.getBuyLinks(query)
}

// GetBuyLinksByMBID fetches a list of Buy Links for a particular Album referenced by its musicbrainz ID.
//
// See http://www.last.fm/api/show/album.getBuylinks
func (a *AlbumMethods) GeyBuyLinksByMBID(mbid string, country string) {

}

func (a *AlbumMethods) getInfo(query map[string]string) (*AlbumInfo, error) {
	wrapper := struct {
		Album *AlbumInfo `xml:"album"`
	}{}
	query["method"] = a.method("getinfo")
	err := a.Client.Execute(query, &wrapper)
	if err != nil {
		return nil, err
	}
	return wrapper.Album, nil
}

// GetInfo fetches the metadata and tracklist for an album referenced by its name and artist-name.
//
// If username is provided, the user's playcount for this album is included in the response.
//
// lang determiens the language of the biography content, expressed as an ISO 639 alpha-2 code (English is en, French is fr etc.).
// For a full list see http://en.wikipedia.org/wiki/List_of_ISO_639-1_codes
// 
// See http://www.last.fm/api/show/album.getInfo
func (a *AlbumMethods) GetInfo(artist, album string, autocorrect bool, username string, lang string) (*AlbumInfo, error) {
	query := a.standardQuery(artist, album, autocorrect)
	query["username"] = username
	query["lang"] = lang
	return a.getInfo(query)
}

// GetInfoByMBID fetches the metadata and tracklist for an album from its Musicbrainz ID.
//
// If username is provided, the user's playcount for this album is included in the response.
//
// lang determiens the language of the biography content, expressed as an ISO 639 alpha-2 code (English is en, French is fr etc.).
// For a full list see http://en.wikipedia.org/wiki/List_of_ISO_639-1_codes
func (a *AlbumMethods) GetInfoByMBID(mbid, username, lang string) (*AlbumInfo, error) {
	query := map[string]string{
		"mbid":     mbid,
		"username": username,
		"lang":     lang,
	}
	return a.getInfo(query)
}

func (a *AlbumMethods) getShouts(query map[string]string) ([]Shout, error) {
	wrapper := struct {
		Shouts []Shout `xml:"shouts>shout"`
	}{}
	query["method"] = a.method("getshouts")
	err := a.Client.Execute(query, &wrapper)
	if err != nil {
		return nil, err
	}
	return wrapper.Shouts, nil
}

// GetShouts fetches a page of 50 "shouts" (comments) for an album.
//
// See http://www.last.fm/api/show/album.getShouts
func (a *AlbumMethods) GetShouts(artist, album string, autocorrect bool, page int) ([]Shout, error) {

	query := map[string]string{
		"artist":      artist,
		"album":       album,
		"autocorrect": boolToString(autocorrect),
		"page":        string(page),
	}
	return a.getShouts(query)
}

// GetShoutsByMBID returns a page of 50 "shouts" (comments) for an album referenced by its Musicbrainz ID.
//
// See http://www.last.fm/api/show/album.getShouts
func (a *AlbumMethods) GetShoutsByMBID(mbid string, page int) ([]Shout, error) {
	query := map[string]string{
		"mbid": mbid,
		"page": string(page),
	}
	return a.getShouts(query)
}

func (a *AlbumMethods) getTopTags(query map[string]string) ([]Tag, error) {
	wrapper := struct {
		Tags []Tag `xml:"toptags>tag"`
	}{}
	query["method"] = a.method("gettoptags")
	err := a.Client.Execute(query, &wrapper)
	if err != nil {
		return nil, err
	}
	return wrapper.Tags, nil
}

// GetTopTags returns the top (most used) tags for an album.
//
// See http://www.last.fm/api/show/album.getTopTags
func (a *AlbumMethods) GetTopTags(artist, album string, autocorrect bool) ([]Tag, error) {
	query := a.standardQuery(artist, album, autocorrect)
	return a.getTopTags(query)
}

// GetTopTagsByMBID returns the top (most used) tags for an album referenced by its musicbrainz ID.
//
// See http://www.last.fm/api/show/album.getTopTags
func (a *AlbumMethods) GetTopTagsByMBID(mbid string) ([]Tag, error) {
	query := map[string]string{
		"mbid": mbid,
	}
	return a.getTopTags(query)
}

// Search finds albums whose name matches (or contains) the value of album
//
// See http://www.last.fm/api/show/album.search
func (a *AlbumMethods) Search(album string, page, limit int) (*AlbumSearchResults, error) {
	wrapper := struct {
		Results *AlbumSearchResults `xml:"results"`
	}{}
	query := map[string]string{
		"method": a.method("search"),
		"album":  album,
		"page":   string(page),
		"limit":  string(limit),
	}
	err := a.Client.Execute(query, &wrapper)
	if err != nil {
		return nil, err
	}
	return wrapper.Results, nil
}
