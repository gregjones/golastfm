package lastfm

import ()

type albumMethods struct {
	Client *Client
}

func (q *albumMethods) Method(suffix string) string {
	return "album." + suffix
}

func (a *albumMethods) GetBuyLinks(artist, album string, autocorrect bool, country string) (*BuyLinks, error) {
	links := BuyLinks{}
	wrapper := struct {
		Affiliations *BuyLinks `xml:"affiliations"`
	}{Affiliations: &links}

	query := map[string]string{
		"method":      a.Method("getbuylinks"),
		"artist":      artist,
		"album":       album,
		"autocorrect": boolToString(autocorrect),
		"country":     country,
	}
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
	query := map[string]string{
		"method":      a.Method("getinfo"),
		"artist":      artist,
		"album":       album,
		"autocorrect": boolToString(autocorrect),
		"username":    username,
		"lang":        lang,
	}
	err := a.Client.Execute(query, &wrapper)
	if err != nil {
		return nil, err
	}
	return &info, nil
}
