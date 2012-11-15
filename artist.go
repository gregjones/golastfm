package lastfm

type ArtistMethods struct {
	Client *Client
}

func (a ArtistMethods) method(suffix string) string {
	return "artist." + suffix
}

func (a ArtistMethods) getInfo(query map[string]string) (*ArtistInfo, error) {
	wrapper := struct {
		Track *ArtistInfo `xml:"artist"`
	}{}
	query["method"] = a.method("getinfo")
	err := a.Client.Execute(query, &wrapper)
	if err != nil {
		return nil, err
	}
	return wrapper.Track, nil
}

func (a ArtistMethods) GetInfo(artist string, autocorrect bool, username, lang string) (*ArtistInfo, error) {
	query := map[string]string{
		"artist":      artist,
		"autocorrect": boolToString(autocorrect),
		"username":    username,
		"lang":        lang,
	}
	return a.getInfo(query)
}

func (a ArtistMethods) GetInfoByMBID(mbid, username, lang string) (*ArtistInfo, error) {
	query := map[string]string{
		"mbid":     mbid,
		"username": username,
		"lang":     lang,
	}
	return a.getInfo(query)
}
