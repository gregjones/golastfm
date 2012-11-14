package lastfm

import ()

type albumMethods struct {
	Client *Client
}

func (a albumMethods) method(suffix string) string {
	return "album." + suffix
}

func (a albumMethods) standardQuery(method, artist, album string, autocorrect bool) map[string]string {
	return map[string]string{
		"method":      a.method(method),
		"artist":      artist,
		"album":       album,
		"autocorrect": boolToString(autocorrect),
	}
}

func (a *albumMethods) GetBuyLinks(artist, album string, autocorrect bool, country string) (*BuyLinks, error) {
	links := BuyLinks{}
	wrapper := struct {
		Affiliations *BuyLinks `xml:"affiliations"`
	}{Affiliations: &links}

	query := a.standardQuery("getbuylinks", artist, album, autocorrect)
	query["country"] = country

	err := a.Client.Execute(query, &wrapper)
	if err != nil {
		return nil, err
	}
	return &links, nil
}

func (a *albumMethods) GeyBuyLinksByMBID(mbid string, country string) {

}

func (a *albumMethods) GetInfo(artist, album string, autocorrect bool, username string, lang string) (*AlbumInfo, error) {
	info := AlbumInfo{}
	wrapper := struct {
		Album *AlbumInfo `xml:"album"`
	}{Album: &info}
	query := a.standardQuery("getinfo", artist, album, autocorrect)
	query["username"] = username
	query["lang"] = lang

	err := a.Client.Execute(query, &wrapper)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (a *albumMethods) GetShouts(artist, album string, autocorrect bool, page int) ([]Shout, error) {
	wrapper := struct {
		Shouts []Shout `xml:"shouts>shout"`
	}{}
	query := map[string]string{
		"method":      a.method("getshouts"),
		"artist":      artist,
		"album":       album,
		"autocorrect": boolToString(autocorrect),
		"page":        string(page),
	}
	err := a.Client.Execute(query, &wrapper)
	if err != nil {
		return nil, err
	}
	return wrapper.Shouts, nil
}

func (a *albumMethods) GetTopTags(artist, album string, autocorrect bool) ([]Tag, error) {
	wrapper := struct {
		Tags []Tag `xml:"toptags>tag"`
	}{}
	query := a.standardQuery("gettoptags", artist, album, autocorrect)
	err := a.Client.Execute(query, &wrapper)
	if err != nil {
		return nil, err
	}
	return wrapper.Tags, nil
}
